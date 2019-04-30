package cmd

import (
  "gpm/pkg/logger"
)

func Help() {
  logger.PrintDescribe("help: Print about all commands")
}
