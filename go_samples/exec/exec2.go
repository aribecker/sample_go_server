package main

import (
    "fmt"
    "os/exec"
    "sync"
	"strings"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  fmt.Println("command is ",cmd)
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
  wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(3)

    //x := []string{"echo newline >> foo.o", "echo newline >> f1.o", "echo newline >> f2.o"}
    x := []string{"dir"}
    go exe_cmd(x[0], wg)
    go exe_cmd(x[1], wg)
    go exe_cmd(x[2], wg)

    wg.Wait()
}
