package internal

import (
  "fmt"
  "github.com/fatih/color"
)

func PrintDescribe(text string) {
  text = fmt.Sprintf(`# %s`, text)
  color.Black(text)
}

func PrintStep(text string) {
  text = fmt.Sprintf(`gpm: %s`, text)
  color.Black(text)
}

func PrintError(err error) {
  text := fmt.Sprintf(`gpm: [ERROR] %s`, err.Error())
  color.Red(text)
}
