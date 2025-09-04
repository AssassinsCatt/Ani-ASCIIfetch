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
    return true
  }else{
    return false
  }
}



func draw(file_name string, id int) {
  path := fmt.Sprintf("animations/%v/%v", file_name, id)
  
  err := print(0,255,0,path)
  
  if err != nil {
    id = 1
  }

  time.Sleep(500 * time.Millisecond)
  fmt.Fprint(os.Stdout, "\x1b[H\x1b[2J\x1b[3J")
  draw(file_name, id+1)
}



func print(red, green, blue int, path string) (err error) {
  file, err := os.ReadFile(path)
  
  if is_err(err) == false {
    color.Green(string(file))

  }else{
    file, _ := os.ReadFile(path)
    color.Green(string(file))
  }

  color.RGB(red, green, blue).Println()
  return
}
