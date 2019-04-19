package internal

import (
  "os"
  "errors"
  "strings"
)

func CheckFileExist(fileName string) (bool, error) {
    if _, err := os.Stat(fileName); err == nil {
      // path/to/whatever exists
      return true, nil
    } else if os.IsNotExist(err) {
      // path/to/whatever does *not* exist
      return false, nil
    }
    return false, errors.New("Something went wrong")
}

func GetCurrentDir() (string, error) {
  dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
  return dir, nil
}

func CheckInsideGoPath(dir string) bool {
  return strings.Contains(dir, os.Getenv("GOPATH"))
}
