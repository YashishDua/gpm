package main

import (
  "os"
  "strings"

  cmd "gpm/cmd"
  internal "gpm/internal"
)

func main() {
  args := os.Args
  if (len(args) == 1) {
    cmd.Help()
    return
  }

  flags := internal.Flags{}

  for _, arg := range args {
    if (strings.Contains(arg, "-")) {
      s := strings.Split(arg, "=")

      if (s[0] == "-vendor") {
        flags.Vendor = true
      } else
      if (s[0] == "-mod") {
        flags.Mod = true
      } else
      if (s[0] == "-path") {
        flags.Path = s[1]
      }
    }
  }

  cmd.Exec(args, flags)
}
