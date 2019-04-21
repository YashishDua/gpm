package cmd

import (
  "os/exec"
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

  insideGoPath := internal.CheckInsideGoPath(dir)
  if (insideGoPath) {
    vendorScript = fmt.Sprintf(`GO111MODULE=on %s`, vendorScript)
  }

  _, scriptErr := exec.Command("/bin/sh", "-c", vendorScript).Output()
  if scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
