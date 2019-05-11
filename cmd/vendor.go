package cmd

import (
  "fmt"

  "github.com/yashishdua/gpm/internal"
)

func SetupVendor() {
  internal.PrintDescribe("Creating vendor...")

  var vendorScript = `go mod vendor`

  dir, dirErr := internal.GetCurrentDir()
  if dirErr != nil {
    internal.PrintError(dirErr)
    return
  }

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    vendorScript = fmt.Sprintf(`GO111MODULE=on %s`, vendorScript)
  }

  internal.PrintStep("using modules to build vendor")
  if scriptErr := internal.ConfigureScript(vendorScript).Run(); scriptErr != nil {
    internal.PrintError(scriptErr)
    return
  }

  internal.PrintStep("Vendor created")
}
