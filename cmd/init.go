package cmd

import (
  "fmt"

  "gpm/internal"
)

func Init()  {
  initScript := `mkdir -p .gpm`
  if scriptErr := internal.ConfigureScript(initScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
    return
  }

  internal.PrintStep("Initialized")
}
