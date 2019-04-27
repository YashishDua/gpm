package cmd

import (
  "fmt"

  "goboil/internal"
)

func Init()  {
  initScript := `mkdir -p .goboil`
  if scriptErr := internal.ConfigureScript(initScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
