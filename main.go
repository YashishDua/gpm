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

  vendorPtr := flag.Bool("vendor", false, "a boolean for using vendor while build")
  modPtr := flag.Bool("mod", false, "a boolean for using mod file while build")
  flag.Parse()

  cmd.Exec(args, vendorPtr, modPtr)
}
