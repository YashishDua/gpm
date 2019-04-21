package cmd

import (
  "os/exec"
  "fmt"

  "goboil/internal"
)

func Build(vendorPtr *bool, modPtr *bool) {
  var buildScript = `go build`
  
  isFileExist, _ := internal.CheckFileExist("go.mod")
  if (isFileExist) {
    // Check if inside GOPATH
    dir, dirErr := internal.GetCurrentDir()
    if dirErr != nil {
      fmt.Println(dirErr)
      return
    }

    insideGoPath := internal.CheckInsideGoPath(dir)
    if (insideGoPath) {
      if (*modPtr) {
        fmt.Println("Using mod file to build..")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s`, buildScript)
      } else
      if (*vendorPtr) {
        fmt.Println("Using vendor to build..")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s -mod=vendor`, buildScript)
      }
    } // Takes mod file if outside GOPATH
  }

  _, scriptErr := exec.Command("/bin/sh", "-c", buildScript).Output()
  if scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
