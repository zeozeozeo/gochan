package gochan

import (
	"fmt"
	"time"
)

// GetPage returns all threads on a page index. Pages are
// indexed with 1.
func (sess *Session) GetPage(boardName string, page int) ([]Thread, error) {
	if page < 1 {
		return nil, fmt.Errorf("pages are indexed with 1, got index %d", page)
	}

	var pageStruct struct {
		Threads []Thread `json:"threads"`
	}

	err := sess.getUnmarshal(
		sess.APIURL,
		fmt.Sprintf("/%s/%d.json", boardName, page),
		&pageStruct,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// add OP's and board names to each thread
	timeRecieved := time.Now()
	for i := 0; i < len(pageStruct.Threads); i++ {
		if len(pageStruct.Threads[i].Replies) > 0 {
			pageStruct.Threads[i].OP = pageStruct.Threads[i].Replies[0]
		}
		pageStruct.Threads[i].Board = boardName
		pageStruct.Threads[i].session = sess
		pageStruct.Threads[i].timeRecieved = timeRecieved
	}

	return pageStruct.Threads, nil
}
