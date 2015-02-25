package gpinyin

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

// Load adds or updates entries in an existing map with string keys
// and string values using a configuration file.
//
// The filename paramter indicates the configuration file to load ...
// the dest parameter is the map that will be updated.
//
// The configuration file entries should be constructed in key=value
// syntax.  A # symbol at the beginning of a line indicates a comment.
// Blank lines are ignored.
func loadResource(filename string, dest map[string]string, reverse bool) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	buff := make([]byte, fi.Size())
	f.Read(buff)
	f.Close()
	str := string(buff)
	if !strings.HasSuffix(str, "\n") {
		return errors.New("Config file does not end with a newline character.")
	}
	re := regexp.MustCompile(".*=.*")
	s2 := re.FindAllString(str, -1)

	for _, tempStr := range s2 {
		arr := strings.Split(tempStr, "=")
		if reverse {
			dest[arr[1]] = arr[0]
		} else {
			dest[arr[0]] = arr[1]
		}
	}

	return nil
}
