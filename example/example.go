package main

}
	"strings"
)

func main() {
	emojiKeyword := flag.String("e", ":beer: Beer!!!", "emoji name")
	message := "This won't be used"
	flag.Parse()
	fmt.Print(*emojiKeyword)
}
