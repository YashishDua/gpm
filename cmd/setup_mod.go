package cmd

import (
  "fmt"
  "bufio"
  "os"
  "os/exec"
  "strings"

  "goboil/internal"
)

func SetupMod() {
  isFileExist, _ := internal.CheckFileExist("go.mod")
  if (isFileExist) {
    fmt.Println("mod file already exist")
    return
  }

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter module name (github.com/username/repo): ")
  text, _ := reader.ReadString('\n')
  modScript := fmt.Sprintf(`go mod init %s`, text)

  // Check if inside GOPATH
  dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

  insideGoPath := strings.Contains(dir, os.Getenv("GOPATH"))
  if (insideGoPath) {
    modScript = fmt.Sprintf(`GO111MODULE=on %s`, modScript)
  }

  _, err = exec.Command("/bin/sh", "-c", modScript).Output()
  if err != nil {
      fmt.Println(err)
  }
}
