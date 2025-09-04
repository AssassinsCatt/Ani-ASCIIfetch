package main

import (
  "fmt"
  "time"
  "os"
  "github.com/fatih/color"
)

func main() {
  file_name := "arch"
  id := 1
  
  draw(file_name, id)  
}

func is_err (err error) bool{
  if err != nil {
    fmt.Println(err)
    return true
  }else{
    return false
  }
}

func draw(file_name string, id int) {
  path := fmt.Sprintf("animations/%v/%v", file_name, id)
  file, err := os.ReadFile(path)
  
  if is_err(err) == false {
    color.Green(string(file))
  }else{
    fmt.Print("\033[H\033[2J") 
    id = 1
    path := fmt.Sprintf("animations/%v/%v", file_name, id)
    file, _ := os.ReadFile(path)
    color.Green(string(file))
  }

  time.Sleep(500 * time.Millisecond)
  fmt.Fprint(os.Stdout, "\x1b[H\x1b[2J\x1b[3J")
  draw(file_name, id+1)
}
