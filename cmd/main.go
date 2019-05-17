package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yashishdua/gpm/internal"

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
			fmt.Println("gpm - Go Project Manager v0.0.1")
		},
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the project",
		Long:  `Initializes the project`,
		Run: func(cmd *cobra.Command, args []string) {
			if isFileExist, _ := internal.CheckFileExist(".gpm"); isFileExist {
				internal.PrintStep("gpm already initialized")
			} else {
				Init()
			}
		},
	}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates directory structure",
		Long:  `Create the recommended project directory structure`,
		Run: func(cmd *cobra.Command, args []string) {
			if preCheck() {
				SetupProject()
			}
		},
	}

	var modCmd = &cobra.Command{
		Use:   "mod",
		Short: "Creates modules file",
		Long:  `Creates modules file`,
		Run: func(cmd *cobra.Command, args []string) {
			if preCheck() {
				SetupMod()
			}
		},
	}

	var vendorCmd = &cobra.Command{
		Use:   "vendor",
		Short: "Creates vendor using modules",
		Long:  `Creates vendor using modules`,
		Run: func(cmd *cobra.Command, args []string) {
			if preCheck() {
				SetupVendor()
			}
		},
	}

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Updates Go version",
		Long:  `Updates Go version to a entered version`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			flags := cmd.Flags()
			version := ""
			flags.VisitAll(func(flag *pflag.Flag) {
				if flag.Shorthand == "v" && flag.Changed {
					version = flag.Value.String()
				}
			})

			if strings.Contains(version, "go") {
				return errors.New("version cannot contain 'go' keyword")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if preCheck() {
				UpdateVersion(internalFlags)
			}
		},
	}

	updateCmd.Flags().StringVarP(&internalFlags.Version, "version", "v", "", "Version Number")

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
				} else if flag.Shorthand == "m" && flag.Changed {
					flagPresent = true
				}
			})

			if !flagPresent {
				return errors.New("build type required (vendor or modules)")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if preCheck() {
				Build(internalFlags)
			}
		},
	}

	buildCmd.Flags().BoolVarP(&internalFlags.Vendor, "vendor", "v", false, "Builds project using vendor")
	buildCmd.Flags().BoolVarP(&internalFlags.Modules, "modules", "m", false, "Builds project using modules")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(modCmd)
	rootCmd.AddCommand(vendorCmd)

	if err := rootCmd.Execute(); err != nil {
		internal.PrintError(err)
	}
}

func preCheck() bool {
	isFileExist, _ := internal.CheckFileExist(".gpm")

	if !isFileExist {
		internal.PrintStep("gpm not initialized")
		internal.PrintStep("Use <gpm init> to initialize the project")
		return false
	}

	return true
}
