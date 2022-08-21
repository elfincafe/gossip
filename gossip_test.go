package gossip

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	_, err := Open("./jipcode.2205.zip")
	if err != nil {
		fmt.Println(err)
	}
}

func TestEndWith(t *testing.T) {
}

func TestExtract(t *testing.T) {
	err := Extract("/tmp/test.zip", "/tmp/extract")
	if err != nil {
		t.Error(err)
	}
}
