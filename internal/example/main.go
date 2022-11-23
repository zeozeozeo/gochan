package main

import (
	"fmt"
	"strings"

	"github.com/zeozeozeo/gochan"
)

func main() {
	// create a new session
	sess := gochan.NewSession()

	// list the first page of /po/
	page, err := sess.GetPage("po", 1)
	panicIfError(err)

	for _, t := range page {
		// OP
		fmt.Printf("(OP) %s, No.%d: (text hidden)\n", t.OP.Name, t.OP.No)

		// replies
		for _, r := range t.Replies {
			fmt.Printf("    %s No.%d says:\n", r.Name, r.No)

			// the comments are HTML escaped
			for _, line := range strings.Split(r.Comment, "<br>") {
				fmt.Printf("        %s\n", line)
			}
		}
		fmt.Print("\n\n\n")
	}

	// get all boards
	boards, err := sess.Boards()
	panicIfError(err)
	// 1 second delay here, because of the API ratelimit

	// print all boards and their titles
	fmt.Println("All boards:")
	for _, board := range boards {
		fmt.Printf("%s: %s\n", board.Name, board.Title)
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
