package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// ProbeConnection n times as defined
func ProbeConnection(ip string, port string, maxRetries int) error {
	counter := 0
	var (
		conn net.Conn
		err  error
	)
	fmt.Printf("Probe connection at %s%s...", ip, port)
	for counter < maxRetries {
		fmt.Print(".")
		conn, err = net.DialTimeout("tcp", ip+port, time.Duration(500)*time.Millisecond)
		if err == nil {
			fmt.Print("Complete with success!!\n")
			return nil
		}
		counter++
		// sleep to avoid block your ip or something else
		time.Sleep(time.Duration(250) * time.Millisecond)
	}

	if conn != nil {
		// don't forget to close the connection
		conn.Close()
	}
	fmt.Print("Fail.\n")
	return err
}

func main() {
	err := ProbeConnection("54.175.219.8", ":22", 10)
	if err != nil {
		log.Fatal(err)
	}
}
