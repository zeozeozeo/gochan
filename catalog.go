package gochan

import "fmt"

type Catalog []struct {
	Page    int    `json:"page"`    // The page number (starts with 1)
	Threads []Post `json:"threads"` // All threads on this page
}

// BoardCatalog returns the catalog of a board. A catalog is a
// comprehensive list of all threads+attributes on a board.
func (sess *Session) BoardCatalog(boardName string) (Catalog, error) {
	var catalog Catalog
	err := sess.getUnmarshal(sess.APIURL, fmt.Sprintf("/%s/catalog.json", boardName), &catalog, nil)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}
