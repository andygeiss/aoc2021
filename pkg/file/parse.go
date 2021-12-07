package file

import (
	"io/ioutil"
	"strings"
)

// Parse ...
func Parse(path string) (lines []string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	return strings.Split(string(data), "\r\n")
}
