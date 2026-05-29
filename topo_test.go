package topotypes

import (
	"os"
	"testing"
)

func TestMk(t *testing.T) {
	bt, _ := os.ReadFile("tests/cat.json")
	CatenaryUnmarshal(bt)
}
