package cmd

import (
  "fmt"
  "bufio"
  "os"
  "gpm/internal"
  "gpm/pkg/logger"
)

func SetupMod() {
  if isFileExist, _ := internal.CheckFileExist("go.mod"); isFileExist {
    logger.PrintStep("Modules file already exist")
    return
  }

  reader := bufio.NewReader(os.Stdin)
  logger.PrintStep("Enter module name: ")
  text, _ := reader.ReadString('\n')
  modScript := fmt.Sprintf(`go mod init %s`, text)

  // Check if inside GOPATH
  dir, dirErr := internal.GetCurrentDir()
	if dirErr != nil {
		logger.PrintError(dirErr)
    return
	}

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    modScript = fmt.Sprintf(`GO111MODULE=on %s`, modScript)
  }

  if scriptErr := internal.ConfigureScript(modScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
  }
}
