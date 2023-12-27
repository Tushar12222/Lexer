package main

import (
    "fmt"
    "os"
    "lexer/utils"
    "lexer/models"
)


func main(){
  var filePath string = os.Args[1]
  err := utils.IsValid(filePath)

  // check if the file is of the valid type
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  
  // read the data from the file
  data , err := utils.ReadData(filePath)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  
  currSource := models.NewSource(data, 0, 0, 1)
  utils.Tokenize(&currSource)
  fmt.Println(currSource.ListOftokens)
}
