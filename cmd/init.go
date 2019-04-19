package cmd

import (
  "fmt"
  "os/exec"
)

var initScript string = `mkdir -p .goboil`

func Init()  {
  fmt.Println("Init...")
  
  _, err := exec.Command("/bin/sh", "-c", initScript).Output()
  if err != nil {
    fmt.Println(err)
  }
}
