package mlnotes

import (
	"testing"
)

func TestHi(t *testing.T) {
	ret, err := Hi("a", "cn")
	if err != nil || ret != "Hi, a and cn" {
		t.Error("expect hia, got", ret)
	}
}
