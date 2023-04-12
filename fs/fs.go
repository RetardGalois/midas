package fs

import (
	"strings"
	"regexp"
	"errors"
)

var file_regex = regexp.MustCompile(`(?i)\bpass.*\.txt\b`)

func File(file_name string) ([]string, error) {
	if strings.HasSuffix(file_name, "zip") {
		return ListZip(file_name)
	}
	return nil, errors.New("not suported file type.")
}

