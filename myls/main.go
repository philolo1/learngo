package main
import "fmt"
import "os/exec"

func main() {
  cmd := exec.Command("ls")

  stdout, err := cmd.CombinedOutput()

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Print(string(stdout))
}
