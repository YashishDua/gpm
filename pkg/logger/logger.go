package logger

import (
  "fmt"
  "github.com/fatih/color"
)

func PrintDescribe(text string) {
  text = fmt.Sprintf(`# %s`, text)
  color.Black(text)
}

func PrintStep(text string) {
  color.Black(text)
}

func PrintError(err error) {
  color.Red(err.Error())
}
