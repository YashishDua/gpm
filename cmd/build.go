package cmd

import (
  "os/exec"
  "fmt"

  "goboil/internal"
)

func Build() {
  dir, err := internal.GetCurrentDir()
  if (err != nil) {
    fmt.Println(err)
    return
  }

  buildScript := `go build`
  insideGoPath := internal.CheckInsideGoPath(dir)
  if (insideGoPath) {
    buildScript = fmt.Sprintf(`GO111MODULE=on %s`, buildScript)
  }

  _, err = exec.Command("/bin/sh", "-c", buildScript).Output()
  if err != nil {
      fmt.Println(err)
  }
}
