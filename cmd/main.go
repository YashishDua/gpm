package cmd

import (
  "fmt"
  "goboil/internal"
)

func Exec(args []string, modPtr *bool) {
  isFileExist, _ := internal.CheckFileExist(".goboil")

  if (args[1] == "init") {
    if (isFileExist) {
      fmt.Println("Goboil already initialized..")
    }
    Init()
    return
  }

  // For any other Command, .goboil must be present
  if (!isFileExist) {
    fmt.Println("Goboil not initialized..")
    return
  }

  if (args[1] == "create") {
    SetupProject()

    if (*modPtr) {
      fmt.Println("Flag inserted..")
    }
    return
  }

  if (args[1] == "mod") {
    SetupMod()
    return
  }

  if (args[1] == "vendor") {
    fmt.Println("Adding vendor support..")
    return
  }
}
