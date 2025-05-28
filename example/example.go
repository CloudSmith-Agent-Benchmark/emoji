package main

import (
	"flag"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	emojiKeyword := flag.String("e", ":beer: Beer!!!", "emoji name")
	message := "This won't be used"
	flag.Parse()

	emoji.Print(*emojiKeyword)
}
