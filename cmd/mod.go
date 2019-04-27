package cmd

import (
  "fmt"
  "bufio"
  "os"

  "goboil/internal"
)

func SetupMod() {
  if isFileExist, _ := internal.CheckFileExist("go.mod"); isFileExist {
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

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    modScript = fmt.Sprintf(`GO111MODULE=on %s`, modScript)
  }

  if scriptErr := internal.ConfigureScript(modScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
