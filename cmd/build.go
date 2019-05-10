package cmd

import (
  "fmt"

  "gpm/internal"
  "gpm/pkg/logger"
)

func Build(flags internal.Flags) {
  buildScript := getBuildScript(flags.Vendor, flags.Modules)
  if buildScript == "" {
    logger.PrintStep("Build failed")
    return
  }

  if scriptErr := internal.ConfigureScript(buildScript).Run(); scriptErr != nil {
    logger.PrintError(scriptErr)
    return
  }
  logger.PrintStep("Build successful")
}

func getBuildScript(vendorFlag bool, modFlag bool) string {
  buildScript := `go build`
  
  if modFlag {
    if modExist, _ := internal.CheckFileExist("go.mod"); modExist {
      dir, dirErr := internal.GetCurrentDir()
      if dirErr != nil {
        logger.PrintError(dirErr)
        return ""
      }

      if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath { // Inside GOPATH
        logger.PrintStep("Using modules to build inside GOPATH")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s`, buildScript)
      }

      return buildScript
    }

    logger.PrintStep("modules doesn't exist")
    return ""
  }

  if vendorFlag {
    if vendorExist, _ := internal.CheckFileExist("vendor"); vendorExist {
      dir, dirErr := internal.GetCurrentDir()
      if dirErr != nil {
        logger.PrintError(dirErr)
        return ""
      }

      if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath { // Inside GOPATH
        logger.PrintStep("Using vendor to build inside GOPATH")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s -mod=vendor`, buildScript)
      }

      return buildScript
    }

    logger.PrintStep("vendor doesn't exist")
    return ""
  }

  logger.PrintStep("No vendor or modules were present")
  return ""
}
