package cmd

import (
  "fmt"
  "os"
  "gpm/internal"
  "gpm/pkg/logger"
)

func UpdateVersion(pathFlag string) {
  if (len(pathFlag) <= 0) {
    logger.PrintStep("Command: gpm update -path=<ABSOLUTE_FILE_PATH>")
    logger.PrintStep("No file path given")
    logger.PrintStep("Download the Go latest version from here: https://golang.org/dl/")
    return
  }

  file, err := os.Open(pathFlag)
  if err != nil {
    logger.PrintError(err)
    return
  }
  defer file.Close()

  if contentType, _ := internal.GetFileContentType(file); contentType != "application/x-gzip" {
    logger.PrintStep(pathFlag + " is not a valid zip file")
    return
  }

  uninstallScript := `sudo rm -rf /usr/local/go`
  extractScript := fmt.Sprintf(`sudo tar -C /usr/local -xzf %s`, pathFlag)
  setScript := `echo $PATH | grep "/usr/local/go/bin"`

  logger.PrintStep("Uninstalling previous version")
  if scriptErr := internal.ConfigureScript(uninstallScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  logger.PrintStep("Extracting archive")
  if scriptErr := internal.ConfigureScript(extractScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  logger.PrintStep("Setting path")
  if scriptErr := internal.ConfigureScript(setScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  logger.PrintStep("Go updated")
}
