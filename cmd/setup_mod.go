package cmd

import (
  "fmt"
  "bufio"
  "os"
  "os/exec"
)

func SetupMod() {
  if _, err := os.Stat("./go.mod"); err == nil {
    // path/to/whatever exists
    fmt.Println("Go mod exist")
    return
  } else if os.IsNotExist(err) {
    // path/to/whatever does *not* exist
    fmt.Println("Go mod does not exist")
  }
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter module name (github.com/username/repo): ")
  text, _ := reader.ReadString('\n')
  out, err := exec.Command("sh", "./scripts/mod.sh", text).Output()
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(string(out))
}
