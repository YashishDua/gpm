package cmd

import (
  "fmt"
  "os/exec"
)

func SetupProject()  {
  _, err := exec.Command("sh", "./scripts/structure.sh").Output()
  if err != nil {
      fmt.Println(err)
  }
}
