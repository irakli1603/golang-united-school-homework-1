package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	fmt.Print(getMessage())
}

func getMessage() string {
	word := "Hello "
	mapEmoji := ":world_map:"
	exclemationChar := "!"
	message := emoji.Sprint(word, mapEmoji, exclemationChar)

	return message
}
