package cmd

import (
  "fmt"

  "goboil/internal"
)

func Exec(args []string, vendorPtr *bool, modPtr *bool) {
  isFileExist, _ := internal.CheckFileExist(".goboil")

  if (args[1] == "init") {
    if (isFileExist) {
      fmt.Println("Goboil already initialized..")
      return
    }
    fmt.Println("Initializing goboil..")
    Init()
    return
  }

  // For any other Command, .goboil must be present
  if (!isFileExist) {
    fmt.Println("Goboil not initialized..")
    return
  }

  if (args[1] == "create") {
    fmt.Println("Setting up project structure..")
    SetupProject()
    return
  }

  if (args[1] == "mod") {
    fmt.Println("Creating mod file..")
    SetupMod()
    return
  }

  if (args[1] == "build") {
    fmt.Println("Building..")
    Build(vendorPtr, modPtr)
    return
  }

  if (args[1] == "vendor") {
    fmt.Println("Creating vendor..")
    SetupVendor()
    return
  }
}
