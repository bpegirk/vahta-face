package modules

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var Con net.Conn

func InitSocket() {
	// Подключаемся к сокету
	conString := fmt.Sprintf("%s:%v", Cfg.Socket.Host, Cfg.Socket.Port)
	Con, _ = net.Dial("tcp", conString)
	fmt.Println("Start listening...")
	fmt.Fprintf(Con, "\n")
	result := make([]byte, 128)
	for {
		_, err := bufio.NewReader(Con).Read(result)
		if err != nil {
			panic(err)
		}
		socketMessage(result)
	}
}

func socketMessage(result []byte) {
	stringBody := string(result)

	lastIndex := 0
	if strings.Contains(stringBody, "<START>") {
		lastIndex = strings.Index(stringBody, "<END>")
	} else {
	}

	var body []byte
	for i := 7; i < lastIndex; i++ {
		body = append(body, result[i])
	}
	fmt.Println("Raw data: ", body)
	ParseBody(body)
}
