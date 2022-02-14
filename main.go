package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	word := "Hello "
	mapEmoji := ":world_map:"
	exclemationChar := "!"
	message := emoji.Sprint(word, mapEmoji, exclemationChar)
	fmt.Print(message)
}
