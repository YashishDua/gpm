package cmd

import (
  "os/exec"
  "fmt"

  "goboil/internal"
)

func Build(vendorFlag bool, modFlag bool) {
  buildScript := `go build`

  if isFileExist, _ := internal.CheckFileExist("go.mod"); isFileExist {
    // Check if inside GOPATH
    dir, dirErr := internal.GetCurrentDir()
    if dirErr != nil {
      fmt.Println(dirErr)
      return
    }

    if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath {
      if (modFlag) {
        fmt.Println("Using mod file to build..")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s`, buildScript)
      } else
      if (vendorFlag) {
        fmt.Println("Using vendor to build..")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s -mod=vendor`, buildScript)
      }
    } // Takes mod file if outside GOPATH
  }

  out, scriptErr := exec.Command("/bin/sh", "-c", buildScript).Output()
  if scriptErr != nil {
    fmt.Println(scriptErr)
  }
  fmt.Println(string(out))
}
