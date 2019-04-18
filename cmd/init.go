package cmd

import (
  "fmt"
  "os/exec"
)

func Init()  {
  fmt.Println("Init...")
  _, err := exec.Command("sh", "./scripts/init.sh").Output()
  if err != nil {
      fmt.Println(err)
  }
}
