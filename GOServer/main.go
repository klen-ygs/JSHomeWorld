package main

import (
	"GOServer/Serverlves"
	"fmt"
)

func main() {
	err := Serverlves.Engine.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
