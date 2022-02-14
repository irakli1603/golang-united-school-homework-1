package main

import (
	"testing"

	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func TestMessage(t *testing.T) {
	expectedOutput := emoji.Sprint("Hello :world_map:!")

	actual := getMessage()

	if actual != expectedOutput {
		errorMsg := fmt.Sprint("expected output to be - ", expectedOutput, " - but instead got - ", actual)
		t.Errorf(errorMsg)
	}
}
