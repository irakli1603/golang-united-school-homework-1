package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func getMessage() string {
	message := emoji.Sprint("Hello :world_map:!")

	return message
}
