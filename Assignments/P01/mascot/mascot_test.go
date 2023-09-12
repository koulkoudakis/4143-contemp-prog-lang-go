package mascot_test

import (
	"testing"

	"github.com/koulkoudakis/4143-contemp-prog-lang-go/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot")
	}
}
