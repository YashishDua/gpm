package cmd

import (
  "fmt"
  "bufio"
  "os"
  "os/exec"
)

func SetupMod() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter module name (github.com/username/repo): ")
  text, _ := reader.ReadString('\n')
  _, err := exec.Command("sh", "./scripts/mod.sh", text).Output()
  if err != nil {
      fmt.Println(err)
  }
}
