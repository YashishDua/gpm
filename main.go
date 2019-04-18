package main

import (
  "os"
  "flag"
  cmd "goboil/cmd"
)

func main() {
  args := os.Args
  if (len(args) == 1) {
    cmd.Help()
    return
  }

  modPtr := flag.Bool("mod", false, "a boolean for Go module support")
  flag.Parse()

  cmd.Exec(args, modPtr)

}
