package mlnotes

import (
	"testing"
)

func TestHi(t *testing.T) {
	ret := Hi("a")
	if ret != "Hi, a" {
		t.Error("expect hia, got", ret)
	}
}
