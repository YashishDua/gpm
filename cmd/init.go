package cmd

import (
	"fmt"

	"github.com/yashishdua/gpm/internal"
)

func Init() {
	internal.PrintDescribe("Initializing gpm...")

	initScript := `mkdir -p .gpm`
	if scriptErr := internal.ConfigureScript(initScript).Run(); scriptErr != nil {
		fmt.Println(scriptErr)
		return
	}

	internal.PrintStep("Initialized")
}
