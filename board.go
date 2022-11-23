package gochan

import "errors"

var (
	ErrBoardNotFound error = errors.New("board not found")
)

// BoardData represents a 4chan board.
type BoardData struct {
	Name            string         `json:"board"`             // The directory the board is located in.
	Title           string         `json:"title"`             // The readable title at the top of the board
	Worksafe        jsonBool       `json:"ws_board"`          // Is the board worksafe
	PerPage         int            `json:"per_page"`          // How many threads are on a single index page
	Pages           int            `json:"pages"`             // How many index pages does the board have
	MaxFilesize     int            `json:"max_filesize"`      // Maximum file size allowed for non .webm attachments
	MaxWebmFilesize int            `json:"max_webm_filesize"` // Maximum file size allowed for .webm attachments
	MaxCommentChars int            `json:"max_comment_chars"` // Maximum number of characters allowed in a post comment
	MaxWebmDuration int            `json:"max_webm_duration"` // Maximum duration of a .webm attachment (in seconds)
	BumpLimit       int            `json:"bump_limit"`        // Maximum number of replies allowed to a thread before the thread stops bumping
	ImageLimit      int            `json:"image_limit"`       // Maximum number of image replies per thread before image replies are discarded
	Cooldowns       BoardCooldowns `json:"cooldowns"`         // Board cooldowns for threads, replies and images (in seconds)
	MetaDescription string         `json:"meta_description"`  // SEO meta description content for a board
	Spoilers        jsonBool       `json:"spoilers"`          // Are spoilers enabled
	CustromSpoilers jsonBool       `json:"custom_spoilers"`   // How many custom spoilers does the board have
	IsArchived      jsonBool       `json:"is_archived"`       // Are archives enabled for the board
	CountryFlags    jsonBool       `json:"country_flags"`     // Are flags showing the poster's country enabled on the board
	UserIDs         jsonBool       `json:"user_ids"`          // Are poster ID tags enabled on the board
	Oekaki          jsonBool       `json:"oekaki"`            // Can users submit drawings via browser the Oekaki app
	SjisTags        jsonBool       `json:"sjis_tags"`         // Can users submit sjis drawings using the [sjis] tags
	CodeTags        jsonBool       `json:"code_tags"`         // Board supports code syntax highlighting using the [code] tags
	MathTags        jsonBool       `json:"math_tags"`         // Board supports [math] TeX and [eqn] tags
	TextOnly        jsonBool       `json:"text_only"`         // Is image posting disabled for the board
	ForcedAnon      jsonBool       `json:"forced_anon"`       // Is the name field disabled on the board
	WebmAudio       jsonBool       `json:"webm_audio"`        // Are webms with audio allowed?
	RequireSubject  jsonBool       `json:"require_subject"`   // Do OPs require a subject
	MinImageWidth   int            `json:"min_image_width"`   // What is the minimum image width (in pixels)
	MinImageHeight  int            `json:"min_image_Height"`  // What is the minimum image height (in pixels)
	// TODO: board_flags. no board seems to have them

	session *Session
}

// BoardCooldowns specifies the cooldowns of the board, in seconds.
type BoardCooldowns struct {
	Threads int `json:"threads"` // Threads cooldown (seconds)
	Replies int `json:"replies"` // Replies cooldown (seconds)
	Images  int `json:"images"`  // Images cooldown (seconds)
}

// Boards returns a list of all boards.
func (sess *Session) Boards() ([]BoardData, error) {
	var b struct {
		Boards []BoardData `json:"boards"`
	}
	err := sess.getUnmarshal(sess.APIURL, "/boards.json", &b, nil)
	if err != nil {
		return nil, err
	}

	for _, board := range b.Boards {
		board.session = sess
	}

	return b.Boards, nil
}

// Board returns a board by name.
func (sess *Session) Board(name string) (*BoardData, error) {
	boards, err := sess.Boards()
	if err != nil {
		return nil, err
	}

	// find the board
	for _, board := range boards {
		if board.Name == name {
			return &board, nil
		}
	}
	return nil, ErrBoardNotFound
}

// BoardNames returns all existing board names.
func (sess *Session) BoardNames() ([]string, error) {
	boards, err := sess.Boards()
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, board := range boards {
		names = append(names, board.Name)
	}
	return names, nil
}

// ArchivedThreads returns the OP numbers of archived threads
// in the board.
func (board *BoardData) ArchivedThreads() ([]int, error) {
	return board.session.BoardArchive(board.Name)
}

// Catalog returns the catalog of the board. A catalog is a
// comprehensive list of all threads+attributes on a board.
func (board *BoardData) Catalog() (Catalog, error) {
	return board.session.BoardCatalog(board.Name)
}

// Page returns all threads on a page index. Pages are
// indexed with 1.
func (board *BoardData) Page(page int) ([]Thread, error) {
	return board.session.GetPage(board.Name, page)
}

// Thread returns a thread by OP's id.
func (board *BoardData) Thread(id int) (*Thread, error) {
	return board.session.GetThread(board.Name, id)
}
