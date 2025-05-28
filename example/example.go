package main

import (
	"flag"

	"github.com/kyokomi/emoji/v2"
	"fmt"
)

func main() {
	emojiKeyword := flag.String("e", ":beer: Beer!!!", "emoji name")
	flag.Parse()

	emoji.Print(*emojiKeyword)
}
