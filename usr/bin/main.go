package main

import (
  "fmt"
  "time"
  "os"
  "github.com/fatih/color"
  "errors"
  "github.com/yuin/gopher-lua"
)


type Color struct {
  Red int
  Green int
  Blue int
}

type Config struct {
  Name string
  Color Color
  FPS int
}


func main() {
  cfg := read_config()
  file_name := cfg.Name
  color := cfg.Color
  fps := cfg.FPS
   
  draw(file_name, color, fps)
}



func is_err (err error) bool{
  if err != nil {
    fmt.Println(err)
    return true
  }else{
    return false
  }
}



func draw(file_name string, color Color, fps int) {
  id := 1
  fps = 10/fps*100
  red, green, blue := color.Red, color.Green, color.Blue
  for {

    path := pathfinder(file_name, id)
  
    err := print(red,green,blue,path)
    
    if errors.Is(err, os.ErrNotExist) {
      clear()
      id = 1
      path := pathfinder(file_name, id)
      print(red,green,blue,path)
    
    }else{
      if err != nil {
        is_err(err)
      }
    }

    time.Sleep(time.Duration(fps) * time.Millisecond)
    clear()
    id += 1
  
  }
}



func print(red, green, blue int, path string) (err error) {
  file, err := os.ReadFile(path)
  
  if is_err(err) == false {
    color.RGB(red, green, blue).Println(string(file))
  }
  return
}



func pathfinder(name string, id int) (file string) {
  home, _ := os.UserHomeDir()
  file = fmt.Sprintf("%v/.config/meowfetch/animations/%v/%v", home, name, id)
  return 
}



func clear() {
  fmt.Fprint(os.Stdout, "\x1b[H\x1b[2J\x1b[3J")
}



func read_config() (cfg Config) {
  L := lua.NewState()
  defer L.Close()
  
  home, _ := os.UserHomeDir()
  file := fmt.Sprintf("%v/.config/meowfetch/config.lua", home)
  is_err(L.DoFile(file))
  
  cfg = Config{}
  configTable := L.GetGlobal("config")
  if tbl, ok := configTable.(*lua.LTable); ok {
    
    if name := tbl.RawGetString("name"); name != lua.LNil {
      cfg.Name = name.String()
    }

    if fps := tbl.RawGetString("fps"); fps != lua.LNil {
      if n, ok := fps.(lua.LNumber); ok {
        cfg.FPS = int(n)
      }
    }

    if colorTbl, ok := tbl.RawGetString("color").(*lua.LTable); ok {
      if red := colorTbl.RawGetString("red"); red != lua.LNil {
        cfg.Color.Red = int(red.(lua.LNumber))
      }
      if green := colorTbl.RawGetString("green"); green != lua.LNil {
        cfg.Color.Green = int(green.(lua.LNumber))
      }
      if blue := colorTbl.RawGetString("blue"); blue != lua.LNil {
        cfg.Color.Blue = int(blue.(lua.LNumber))
      }
    }


  }else{
    fmt.Println("config fehler")
    cfg.Name = "arch"
    cfg.Color.Red = 255
    cfg.Color.Green = 0
    cfg.Color.Blue = 0
    cfg.FPS = 2
  }
  return
}

