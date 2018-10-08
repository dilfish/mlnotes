package mlnotes

import (
	"errors"
	"fmt"
)

func Hi(name, lang string) (string, error) {
	if lang != "cn" {
		return "", errors.New("bad lang")
	}
	return fmt.Sprintf("Hi, %s and %s", name, lang), nil
}
