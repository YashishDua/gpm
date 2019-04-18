package internal

import (
  "os"
  "errors"
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
