package util

import (
	"io/ioutil"
	"os"
)

// ReadFile reads the file at path and returns the contents as a string.
func ReadFile(path string) ([]byte, error) {
	c, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(c *os.File) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)

	return ioutil.ReadAll(c)
}
