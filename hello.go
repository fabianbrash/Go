package main

import (
    "fmt"
    "os"
)


func main() {


    name, err := os.Hostname()
    
      if err != nil {
      
        panic(err)
      }
      
  fmt.Println("Hello, Fedora Iam running on host:", name)
 
 }
