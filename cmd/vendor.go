package cmd

import (
  "fmt"

  "gpm/internal"
)

func SetupVendor() {
  var vendorScript = `go mod vendor`

  dir, dirErr := internal.GetCurrentDir()
  if dirErr != nil {
    internal.PrintError(dirErr)
    return
  }

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    vendorScript = fmt.Sprintf(`GO111MODULE=on %s`, vendorScript)
  }

  if scriptErr := internal.ConfigureScript(vendorScript).Run(); scriptErr != nil {
    internal.PrintError(scriptErr)
    return
  }

  internal.PrintStep("Vendor created")
}
