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
  dir, dirErr := internal.GetCurrentDir()
	if dirErr != nil {
		fmt.Println(dirErr)
    return
	}

  insideGoPath := internal.CheckInsideGoPath(dir)
  if (insideGoPath) {
    modScript = fmt.Sprintf(`GO111MODULE=on %s`, modScript)
  }

  _, scriptErr := exec.Command("/bin/sh", "-c", modScript).Output()
  if scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
