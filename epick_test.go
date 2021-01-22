package epick_test

import (
	"fmt"
	"testing"

	"gitlab.com/ethanbakerdev/epick"
)

func Test_Start(t *testing.T) {
	emojiValue, err := epick.Start(false)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(emojiValue)
}
