package cmd

import (
  "fmt"

  "goboil/internal"
)

func SetupVendor() {
  var vendorScript = `go mod vendor`

  // Check if inside GOPATH
  dir, dirErr := internal.GetCurrentDir()
  if dirErr != nil {
    fmt.Println(dirErr)
    return
  }

  if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
    vendorScript = fmt.Sprintf(`GO111MODULE=on %s`, vendorScript)
  }

  if scriptErr := internal.ConfigureScript(vendorScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
