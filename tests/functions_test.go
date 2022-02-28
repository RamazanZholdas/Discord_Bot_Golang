package tests

import (
	"discordTestBot/config"
	"testing"
)

func TestFunction(t *testing.T) {
	err := config.ReadConfig()
	if err != nil {
		t.Error("Test failed:", err)
	}

	if len(config.Token) < 2 {
		t.Error("You didnt provide a token")
	}
}
