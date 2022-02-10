package main

import (
	"fmt"
	module "github.com/bpegirk/vahta-face/modules"
)

func main() {
	fmt.Println("Init config...")
	module.InitConfig()
	fmt.Println("Init fortnet sockets ...")
	module.InitSocket()
	//fmt.Println("Init SocketIo...")
	//module.InitSocketIo()

	fmt.Println("Exit")
}
