package cmd

import (
  "fmt"

  "github.com/yashishdua/gpm/internal"
)

func UpdateVersion(internalFlags internal.Flags) {
  internal.PrintDescribe("Updating Go version...")

  goBinaryFile := fmt.Sprintf(`go%s.darwin-amd64.tar.gz`, internalFlags.Version)
  downloadURL := fmt.Sprintf(`https://dl.google.com/go/%s`, goBinaryFile)
  uninstallScript := `sudo rm -rf /usr/local/go`
  extractScript := fmt.Sprintf(`sudo tar -C /usr/local -xzf %s`, goBinaryFile)
  removeBinaryScript := fmt.Sprintf(`sudo rm %s`, goBinaryFile)

  internal.PrintStep("Uninstalling previous version")
  if scriptErr := internal.ConfigureScript(uninstallScript).Run(); scriptErr != nil {
    internal.PrintError(scriptErr)
    return
  }

  internal.PrintStep("Downloading latest Go binary")
  if fileExist, _ := internal.CheckFileExist(goBinaryFile); fileExist {
    internal.PrintStep("Go binary file already exist")
  } else {
    downloadErr := internal.DownloadFile(goBinaryFile, downloadURL)
    if downloadErr != nil {
      internal.PrintError(downloadErr)
      internal.PrintStep("Go Server error or Check version entered once")
      return
    }
  }

  internal.PrintStep("Extracting Go archive")
  if scriptErr := internal.ConfigureScript(extractScript).Run(); scriptErr != nil {
    internal.PrintError(scriptErr)
    return
  }

  if scriptErr := internal.ConfigureScript(removeBinaryScript).Run(); scriptErr != nil {
    internal.PrintError(scriptErr)
    return
  }

  internal.PrintStep("Go updated successfuly")
}
