package internal

import (
  "fmt"
  "io"
  "os"
  "os/exec"
  "net/http"
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
    return false, errors.New("Error checking file")
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

func CheckGoVersion() (string, error) {
  out, err := exec.Command("/bin/sh", "-c", "go version").Output()
  if err != nil {
      return "", err
  }
  words := strings.Fields(string(out))
  return words[2], nil
}

func GetFileContentType(out *os.File) (string, error) {
    // Only the first 512 bytes are used to sniff the content type.
    buffer := make([]byte, 512)

    _, err := out.Read(buffer)
    if err != nil {
      return "", err
    }

    // Use the net/http package's handy DectectContentType function. Always returns a valid
    // content-type by returning "application/octet-stream" if no others seemed to match.
    contentType := http.DetectContentType(buffer)
    return contentType, nil
}

func ConfigureScript(script string) *exec.Cmd {
  cmd := exec.Command("/bin/sh", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd
}

func DownloadFile(filepath string, url string) error {
  // Create the file
  out, createErr := os.Create(filepath)
  if createErr != nil  {
    return createErr
  }
  fmt.Println("CREATED")
  defer out.Close()

  // Get the data
  resp, httpErr := http.Get(url)
  if httpErr != nil {
    return httpErr
  }
  fmt.Println("HTTP")
  defer resp.Body.Close()

  // Check server response
  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("Bad status from golang.org: %s", resp.Status)
  }

  fmt.Println("COPYING")
  // Writer the body to file
  _, err := io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}
