package cmd

import (
  "fmt"

  "gpm/internal"
)

func Exec(args []string, flags internal.Flags) {
  isFileExist, _ := internal.CheckFileExist(".gpm")
  command := args[1]

  if command == "init" {
    if isFileExist {
      fmt.Println("gpm already initialized..")
    } else {
      fmt.Println("Initializing goboil..")
      Init()
    }
    return
  }

  if command == "help" {
    Help()
    return
  }

  // For any other Command, .goboil must be present
  if !isFileExist {
    fmt.Println("gpm not initialized..")
    return
  }

  switch command {
  case "create":
    fmt.Println("Setting up project structure..")
    SetupProject()

  case "mod":
    fmt.Println("Creating mod file..")
    SetupMod()

  case "build":
    fmt.Println("Building..")
    Build(flags.Vendor, flags.Mod)

  case "vendor":
    fmt.Println("Creating vendor..")
    SetupVendor()

  case "update":
    fmt.Println("Updating Go version..")
    UpdateVersion(flags.Path)

  default:
    fmt.Println("No such command. Use help to see all available commands.")
  }
}
