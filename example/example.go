package main

import (
	"flag"
	"fmt"
)

func main() {
	emojiKeyword := flag.String("e", ":beer: Beer!!!", "emoji name")
	flag.Parse()
	fmt.Print(*emojiKeyword)
}
