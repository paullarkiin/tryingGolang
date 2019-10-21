package main

import (
	"net"
	"fmt"
)

var host string
var port string

func main()  {

	
	fmt.Println("Port Scanner")
	fmt.Println("------------")

	fmt.Println("Enter Host IP: ")
	fmt.Scan(&host)

	fmt.Println("Enter Port Number: ")
	fmt.Scan(&port)

	
	conn, err := net.Dial("tcp", host +":" +port)
	if err != nil {
		panic(err)
	}
	fmt.Print("\n","Port ", port, " Open")
	conn.Close()
}
