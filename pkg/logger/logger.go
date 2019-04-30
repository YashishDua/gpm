package logger

import (
  "fmt"
  "github.com/fatih/color"
)

func PrintDescribe(text string) {
  color.Blue("[ ] " + text)
}

func PrintStep(text string) {
  text = fmt.Sprintf(`    %s`, text)
  fmt.Println(text)
}
