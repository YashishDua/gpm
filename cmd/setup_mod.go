package cmd

import (
  "fmt"
  "bufio"
  "os"
  "os/exec"

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
  dir, err := internal.GetCurrentDir()
	if err != nil {
		fmt.Println(err)
    return
	}

  insideGoPath := internal.CheckInsideGoPath(dir)
  if (insideGoPath) {
    modScript = fmt.Sprintf(`GO111MODULE=on %s`, modScript)
  }

  _, err = exec.Command("/bin/sh", "-c", modScript).Output()
  if err != nil {
      fmt.Println(err)
  }
}
