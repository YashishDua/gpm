package cmd

import (
  "flag"
)

func Exec()  {
  modPtr := flag.Bool("mod", false, "a boolean for Go module support")
  flag.Parse()

  SetupProject()

  if (*modPtr) {
    SetupMod()
  }
}
