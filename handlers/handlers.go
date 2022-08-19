package handlers

import (
	"fmt"
	"net"
)

//Recieve recieves the bytes from client or server
func Recieve(conn net.Conn) (string, error) {
	input := make([]byte, (1024 * 4))
	n, err := conn.Read(input)
	if /*n == 0 ||*/ err != nil {
		fmt.Println("Read error from Recieve:", err)
		return "", err
	}
	return string(input[0:n]), nil
}
