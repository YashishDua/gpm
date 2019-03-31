package main

import (
  "flag"
  "fmt"
  "bufio"
  "os"
  "os/exec"
)

func main()  {
  modPtr := flag.Bool("mod", false, "a boolean for Go module support")
  flag.Parse()

  if (*modPtr) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter module name (github.com/username/repo): ")
    text, _ := reader.ReadString('\n')
    out, err := exec.Command("sh", "./scripts/mod.sh", text).Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%s", out)
  }

  out, err := exec.Command("sh", "./scripts/test.sh").Output()
  if err != nil {
      fmt.Println(err)
  }
  fmt.Printf("%s", out)
}
