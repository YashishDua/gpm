package cmd

import (
  "fmt"
  "bufio"
  "os"

  "github.com/yashishdua/gpm/internal"
)

func SetupMod() {
  internal.PrintDescribe("Creating modules file...")

  if isFileExist, _ := internal.CheckFileExist("go.mod"); isFileExist {
    internal.PrintStep("Modules file already exist")
    return
  }

  reader := bufio.NewReader(os.Stdin)
  internal.PrintStep("Enter module name: ")
  text, _ := reader.ReadString('\n')
  modScript := fmt.Sprintf(`go mod init %s`, text)

  // Check if inside GOPATH
  dir, dirErr := internal.GetCurrentDir()
	if dirErr != nil {
		internal.PrintError(dirErr)
    return
	}

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    modScript = fmt.Sprintf(`GO111MODULE=on %s`, modScript)
  }

  if scriptErr := internal.ConfigureScript(modScript).Run(); scriptErr != nil {
    internal.PrintError(scriptErr)
  }
}
