package main
import (
	"testing"
	// "fmt"
)

func TestMain(t *testing.T) {
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`
	expected := "BUY NOW, SAVE 10%"
	s := CleanupMessage(message)

	if s != expected {
		t.Errorf("Strings are not equals lol")
	}
}
