package conn

import (
	"os"
	"fmt"
	"net"
)

// Just Test
func Connectp string)(net.Conn){
	connection, err := net.Dial("tcp", p)
	if err != nil {
		fmt.Println("Error connecting  with the client.")
	} else {
		fmt.Printf("Connection established with %s \n", connection.RemoteAddr())
	}
	return connection
}
