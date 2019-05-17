package cmd

import (
	"errors"
	"fmt"

	"github.com/yashishdua/gpm/internal"
)

func Build(flags internal.Flags) {
	internal.PrintDescribe("Building...")

	buildScript := getBuildScript(flags.Vendor, flags.Modules)
	if buildScript == "" {
		internal.PrintError(errors.New("Build failed"))
		return
	}

	if scriptErr := internal.ConfigureScript(buildScript).Run(); scriptErr != nil {
		internal.PrintError(scriptErr)
		return
	}

	internal.PrintStep("Build successful")
}

func getBuildScript(vendorFlag bool, modFlag bool) string {
	buildScript := `go build`

	if modFlag {
		if modExist, _ := internal.CheckFileExist("go.mod"); modExist {
			dir, dirErr := internal.GetCurrentDir()
			if dirErr != nil {
				internal.PrintError(dirErr)
				return ""
			}

			if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath { // Inside GOPATH
				internal.PrintStep("Using modules to build inside GOPATH")
				buildScript = fmt.Sprintf(`GO111MODULE=on %s`, buildScript)
			}

			return buildScript
		}

		internal.PrintStep("modules doesn't exist")
		return ""
	}

	if vendorFlag {
		if vendorExist, _ := internal.CheckFileExist("vendor"); vendorExist {
			dir, dirErr := internal.GetCurrentDir()
			if dirErr != nil {
				internal.PrintError(dirErr)
				return ""
			}

			if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath { // Inside GOPATH
				internal.PrintStep("Using vendor to build inside GOPATH")
				buildScript = fmt.Sprintf(`GO111MODULE=on %s -mod=vendor`, buildScript)
			}

			return buildScript
		}

		internal.PrintStep("vendor doesn't exist")
		return ""
	}

	internal.PrintStep("no vendor or modules were present")
	return ""
}
