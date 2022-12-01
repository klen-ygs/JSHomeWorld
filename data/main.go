package main

import (
	"data/Config"
	"data/Servers"
)

var port = Config.GetString("port")
var ip = Config.GetString("ip")

func main() {

	err := Servers.Engine.Run(ip + ":" + port)
	if err != nil {
		panic(err)
	}
}
