package cmd

import (
  "fmt"
  
  "gpm/internal"
  "gpm/pkg/logger"
)

func Build(vendorFlag bool, modFlag bool) {
  buildScript := getBuildScript(vendorFlag, modFlag)
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

  // Priority given to modules unless specified
  if modExist, _ := internal.CheckFileExist("go.mod"); modExist {
    dir, dirErr := internal.GetCurrentDir()
    if dirErr != nil {
      logger.PrintError(dirErr)
      return ""
    }

    if insideGoPath := internal.CheckInsideGoPath(dir); insideGoPath { // Inside GOPATH
      if (modFlag) {
        logger.PrintStep("Using modules to build inside GOPATH")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s`, buildScript)
      }
    } else { // Outside GOPATH
      if (vendorFlag) {
        logger.PrintStep("Using vendor to build inside GOPATH")
        buildScript = fmt.Sprintf(`GO111MODULE=on %s -mod=vendor`, buildScript)
      }
    }

    return buildScript
  }

  if vendorExist, _ := internal.CheckFileExist("vendor"); vendorExist {
    logger.PrintStep("Using vendor to build outside GOPATH")
    return buildScript
  }

  logger.PrintStep("No vendor or modules were present")
  return ""
}
