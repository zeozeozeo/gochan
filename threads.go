package gochan

import (
	"fmt"
	"net/http"
	"time"
)

type Thread struct {
	Replies      []*Post `json:"posts"` // Preview replies to the thread
	OP           *Post   // OP's post
	Board        string  // The board this thread is on
	session      *Session
	cooldown     <-chan time.Time
	timeRecieved time.Time
}

// GetThread returns a board thread by OP's id.
func (sess *Session) GetThread(boardName string, id int) (*Thread, error) {
	return sess.getThreadStale(boardName, id, time.Unix(0, 0))
}

func (sess *Session) getThreadStale(boardName string, id int, staleTime time.Time) (*Thread, error) {
	var thread Thread

	err := sess.getUnmarshal(
		sess.APIURL,
		fmt.Sprintf("/%s/thread/%d.json", boardName, id),
		&thread,
		// https://github.com/4chan/4chan-API#api-rules
		func(r *http.Request) error {
			if staleTime.Unix() != 0 {
				r.Header.Add("If-Modified-Since", staleTime.UTC().Format(http.TimeFormat))
			}
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	if len(thread.Replies) > 0 {
		thread.OP = thread.Replies[0]
	}
	thread.Board = boardName
	thread.session = sess
	thread.timeRecieved = time.Now()
	return &thread, nil
}

// Update updates the thread with a cooldown (the default is 15 seconds,
// and the minimum is 10 seconds).
func (thread *Thread) Update() (newPosts, deletedPosts int, err error) {
	sess := thread.session

	// update the thread
	sess.cooldownMutex.Lock()
	if thread.cooldown != nil {
		<-thread.cooldown
	}

	newThread, err := sess.getThreadStale(
		thread.Board,
		thread.No(),
		thread.timeRecieved,
	)

	// thread updating should be set to a minimum of 10 seconds, preferably higher
	if sess.UpdateCooldown < 10*time.Second {
		sess.UpdateCooldown = 10 * time.Second
	}
	thread.cooldown = time.After(sess.UpdateCooldown)
	sess.cooldownMutex.Unlock()

	if err != nil {
		return 0, 0, err
	}

	// check for new and deleted posts
	var a, b int
	for a, b = 0, 0; a < len(thread.Replies); a, b = a+1, b+1 {
		if thread.Replies[a].No == thread.Replies[b].No {
			continue
		}

		// post has been deleted, go back one post
		deletedPosts++
		b--
	}

	newPosts = len(thread.Replies) - b
	thread.Replies = newThread.Replies
	return
}

// ID returns the OP's numeric post ID.
func (thread *Thread) No() int {
	return thread.OP.No
}
