package cmd

import (
  "fmt"
  "os"

  "gpm/internal"
)

func UpdateVersion(pathFlag string) {
  if (len(pathFlag) <= 0) {
    fmt.Println("Provide an absolute path of archive file")
    fmt.Println("You can download the latest version from here: https://golang.org/dl/")
    return
  }

  file, err := os.Open(pathFlag)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  if contentType, _ := internal.GetFileContentType(file); contentType != "application/x-gzip" {
    fmt.Println(pathFlag, "is not a valid zip file")
    return
  }

  uninstallScript := `sudo rm -rf /usr/local/go`
  extractScript := fmt.Sprintf(`sudo tar -C /usr/local -xzf %s`, pathFlag)
  setScript := `echo $PATH | grep "/usr/local/go/bin"`

  fmt.Println("Uninstalling previous version..")
  if scriptErr := internal.ConfigureScript(uninstallScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
    return
  }

  fmt.Println("Extracting archive..")
  if scriptErr := internal.ConfigureScript(extractScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
    return
  }

  fmt.Println("Setting path..")
  if scriptErr := internal.ConfigureScript(setScript).Run(); scriptErr != nil {
    fmt.Println(scriptErr)
  }
}
