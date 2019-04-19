package cmd

import (
  "fmt"
  "os/exec"
  "bufio"
  "strconv"
)

// Array cannot me made constant in Go
var dirs = []string{"cmd", "internal", "pkg", "scripts", "api", "test"}

func SetupProject() {
  for _, dir := range dirs {
    execSetupScript(dir)
  }
}

func execSetupScript(dir string) {
  scripts := getScripts(dir)
  script, countScript, keepScript := scripts[0], scripts[1], scripts[2]
  exec.Command("/bin/sh", "-c", script).Output()
  cmd := exec.Command("/bin/sh", "-c", countScript)
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    fmt.Println(err)
    return
  }

  if err := cmd.Start(); err != nil {
    fmt.Println(err)
    return
  }

  scanner := bufio.NewScanner(stdout)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    value, _ := strconv.Atoi(scanner.Text())
    if (value == 0) {
      exec.Command("/bin/sh", "-c", keepScript).Output()
      return
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}

func getScripts(dir string) []string {
  script := fmt.Sprintf(`mkdir -p %s`, dir)
  countScript := fmt.Sprintf(`cd %s && ls | wc -l`, dir)
  keepScript := fmt.Sprintf(`cd %s && touch .keep`, dir)
  return []string{script, countScript, keepScript}
}
