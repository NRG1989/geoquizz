package main

import (
	"bufio"
	"europe/handlers"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {

		//2 клиент получает страну
		source, err := handlers.Recieve(conn)
		if err != nil {
			return
		}
		fmt.Print("What is the capital of ", source, "-")

		//3 клиент вводит ответ и отправляет серверу
		var answer string

		inputReader := bufio.NewReader(os.Stdin)
		answer, err = inputReader.ReadString('\n')
		if err != nil {
			return
		}

		answer = strings.TrimSuffix(answer, "\n")

		if n1, err := conn.Write([]byte(answer)); n1 == 0 || err != nil {
			fmt.Println(err)
			return
		}
		//5 Клиент получает ответ верно или нет

		answerFromServer, err := handlers.Recieve(conn)
		if err != nil {
			return
		}

		fmt.Println(answerFromServer)

	}
}
