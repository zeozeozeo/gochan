package gochan

import "fmt"

// BoardArchive returns the OP numbers of archived threads of a board.
// Archived threads are read-only threads that are are closed to new
// replies and images.
func (sess *Session) BoardArchive(boardName string) ([]int, error) {
	var archive []int
	err := sess.getUnmarshal(sess.APIURL, fmt.Sprintf("/%s/archive.json", boardName), &archive, nil)
	if err != nil {
		return nil, err
	}
	return archive, nil
}
