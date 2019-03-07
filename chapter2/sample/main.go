package main

import (
	"awesomeProject/chapter2/sample/search"
	"log"
	"os"
)

func init(){
	log.SetOutput(os.Stdout)
}


func main() {
	search.Run("nbcnews")
}
