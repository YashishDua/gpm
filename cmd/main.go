package cmd

import (
  "gpm/internal"
  "fmt"
  "errors"

  "github.com/spf13/cobra"
  "github.com/spf13/pflag"
)

var internalFlags internal.Flags

func Exec() {
  var rootCmd = &cobra.Command{Use: "gpm"}

  var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of gpm",
    Long:  `All software has versions. This is gpm's`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("gpm - Go Package Manager v0.0.1")
    },
  }

  var buildCmd = &cobra.Command{
    Use:   "build",
    Short: "Build the project",
    Long:  `Helps building project using mod and vendor, and works inside and outside the GOPATH`,
    PreRunE: func(cmd *cobra.Command, args []string) error {
      flags := cmd.Flags()
      flagPresent := false
      flags.VisitAll(func(flag *pflag.Flag) {
        if flag.Shorthand == "v" && flag.Changed {
          flagPresent = true
        } else
        if flag.Shorthand == "m" && flag.Changed {
          flagPresent = true
        }
      })

      if !flagPresent {
        return errors.New("Build type required (vendor or modules)")
      }
      return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
      Build(internalFlags)
    },
  }

  buildCmd.Flags().BoolVarP(&internalFlags.Vendor, "vendor", "v", false, "Builds project using vendor")
  buildCmd.Flags().BoolVarP(&internalFlags.Modules, "modules", "m", false, "Builds project using modules")

  rootCmd.AddCommand(versionCmd)
  rootCmd.AddCommand(buildCmd)
  rootCmd.Execute()

  /*isFileExist, _ := internal.CheckFileExist(".gpm")
  command := args[1]

  if command == "init" {
    if isFileExist {
      internal.PrintStep("gpm already initialized")
    } else {
      internal.PrintDescribe("Initializing gpm...")
      Init()
    }
    return
  }

  if command == "help" {
    internal.PrintDescribe("Print about all commands...")
    Help()
    return
  }

  // For any other Command, .goboil must be present
  if !isFileExist {
    internal.PrintStep("gpm not initialized")
    internal.PrintStep("Use <gpm init> to initialize the project")
    return
  }

  switch command {
  case "create":
    internal.PrintDescribe("Setting up project structure...")
    SetupProject()

  case "mod":
    internal.PrintDescribe("Creating modules file...")
    SetupMod()

  case "vendor":
    internal.PrintDescribe("Creating vendor...")
    SetupVendor()

  case "update":
    internal.PrintDescribe("Updating Go version...")
    UpdateVersion(flags.Version)

  default:
    internal.PrintStep("No such command. Use help to see all available commands.")
  }*/
}
