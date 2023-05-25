package main

import (
	"fmt"
	"task/cmd"
	"task/db"
)

func main() {
	err := db.Initdb()
	if err != nil {
		panic(err)
	}
	err = cmd.RootCmd.Execute()
	if err != nil {
		fmt.Print(err)
	}

}
