package cmd

import (
  "fmt"

  "gpm/internal"
  "gpm/pkg/logger"
)

func Init()  {
  initScript := `mkdir -p .gpm`
  if scriptErr := internal.ConfigureScript(initScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
    return
  }

  logger.PrintStep("Initialized")
}
