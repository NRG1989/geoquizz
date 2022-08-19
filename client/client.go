package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"europe/handlers"
)

func main() {
	//connecting to Server
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {

		//2 Client recives the country
		source, err := handlers.Recieve(conn)
		if err != nil {
			return
		}
		fmt.Print("What is the capital of ", source, "-")

		//3 Client insert the capital and sends it
		//var answer string

		inputReader := bufio.NewReader(os.Stdin)
		answer, err := inputReader.ReadString('\n')
		if err != nil {
			return
		}

		answer = strings.TrimSuffix(answer, "\n")
		if answer == "" {
			answer = " "
		}

		if _, err := conn.Write([]byte(answer)); err != nil {
			fmt.Println(err)
			return
		}
		//5 Client recieves the message if his answer was right

		answerFromServer, err := handlers.Recieve(conn)
		if err != nil {
			return
		}

		fmt.Println(answerFromServer)

	}
}
