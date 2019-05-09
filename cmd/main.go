package cmd

import (
  "gpm/internal"
  "gpm/pkg/logger"
)

func Exec(args []string, flags internal.Flags) {
  isFileExist, _ := internal.CheckFileExist(".gpm")
  command := args[1]

  if command == "init" {
    if isFileExist {
      logger.PrintStep("gpm already initialized")
    } else {
      logger.PrintDescribe("Initializing gpm...")
      Init()
    }
    return
  }

  if command == "help" {
    logger.PrintDescribe("Print about all commands...")
    Help()
    return
  }

  // For any other Command, .goboil must be present
  if !isFileExist {
    logger.PrintStep("gpm not initialized")
    logger.PrintStep("Use <gpm init> to initialize the project")
    return
  }

  switch command {
  case "create":
    logger.PrintDescribe("Setting up project structure...")
    SetupProject()

  case "mod":
    logger.PrintDescribe("Creating modules file...")
    SetupMod()

  case "build":
    logger.PrintDescribe("Building...")
    Build(flags.Vendor, flags.Mod)

  case "vendor":
    logger.PrintDescribe("Creating vendor...")
    SetupVendor()

  case "update":
    logger.PrintDescribe("Updating Go version...")
    UpdateVersion(flags.Path)

  default:
    logger.PrintStep("No such command. Use help to see all available commands.")
  }
}
