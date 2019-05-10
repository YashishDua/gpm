package cmd

import (
  "fmt"
  
  "gpm/internal"
  "gpm/pkg/logger"
)

func SetupVendor() {
  var vendorScript = `go mod vendor`

  dir, dirErr := internal.GetCurrentDir()
  if dirErr != nil {
    logger.PrintError(dirErr)
    return
  }

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    vendorScript = fmt.Sprintf(`GO111MODULE=on %s`, vendorScript)
  }

  if scriptErr := internal.ConfigureScript(vendorScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }

  logger.PrintStep("Vendor created")
}
