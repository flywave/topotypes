package topotypes

import (
	"io/ioutil"
	"testing"
)

func TestMk(t *testing.T) {
	bt, _ := ioutil.ReadFile("tests/cat.json")
	CatenaryUnMarshal(bt)
}
