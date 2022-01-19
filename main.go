package main

import (
  "bufio"
	"fmt"
  "os"
  "strings"
)

func main(){
  fmt.Print("Print a single word as filename:")
  terminalReader:=bufio.NewReader(os.Stdin)
  filename,_:=terminalReader.ReadString('\n')
  filename=strings.TrimSpace(filename)
  filename +=".txt"

  fmt.Println("The filename is:", filename)
  }
