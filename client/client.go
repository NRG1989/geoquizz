package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0:n])
		fmt.Print("What is the capital of ", source, "-")

		//3 клиент вводит ответ и отправляет серверу
		var answer string

		inputReader := bufio.NewReader(os.Stdin)
		answer, _ = inputReader.ReadString('\n')
		if err != nil {
			continue
		}

		if n1, err := conn.Write([]byte(answer)); n1 == 0 || err != nil {
			fmt.Println(err)
			return
		}
		//5 Клиент получает ответ верно или нет
		input1 := make([]byte, (1024 * 4))
		n2, err := conn.Read(input1)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source1 := string(input1[0:n2])

		fmt.Println(source1)

	}
}
