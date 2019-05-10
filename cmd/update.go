package cmd

import (
  "fmt"
  "strings"

  "gpm/internal"
  "gpm/pkg/logger"
)

func UpdateVersion(version string) {
  if len(version) <= 0 { // Default Version
    version = "1.12.5"
  }

  if strings.Contains(version, "go") {
    logger.PrintStep("Version cannot contain 'go' keyword")
    return
  }

  goBinaryFile := fmt.Sprintf(`go%s.darwin-amd64.tar.gz`, version)
  downloadURL := fmt.Sprintf(`https://dl.google.com/go/%s`, goBinaryFile)
  uninstallScript := `sudo rm -rf /usr/local/go`
  extractScript := fmt.Sprintf(`sudo tar -C /usr/local -xzf %s`, goBinaryFile)
  removeBinaryScript := fmt.Sprintf(`sudo rm %s`, goBinaryFile)

  logger.PrintStep("Uninstalling previous version")
  if scriptErr := internal.ConfigureScript(uninstallScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  logger.PrintStep("Downloading latest Go binary")
  if fileExist, _ := internal.CheckFileExist(goBinaryFile); fileExist {
    logger.PrintStep("Go binary file already exist")
  } else {
    downloadErr := internal.DownloadFile(goBinaryFile, downloadURL)
    if downloadErr != nil {
      logger.PrintError(downloadErr)
      logger.PrintStep("Go Server error or Check version entered once")
      return
    }
  }

  logger.PrintStep("Extracting Go archive")
  if scriptErr := internal.ConfigureScript(extractScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  if scriptErr := internal.ConfigureScript(removeBinaryScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  logger.PrintStep("Go updated successfuly")
}
