package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	nameToHost := make(map[string]string, 0)
	for _, v := range os.Args[1:] {
		splitted := strings.Split(v, "=")
		nameToHost[splitted[0]] = splitted[1]
	}

	for name, host := range nameToHost {
		go func(name, host string) {
			conn, err := net.Dial("tcp", host)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			input := bufio.NewScanner(conn)
			for input.Scan() {
				fmt.Fprintf(os.Stdout, "%s\n %s \n", name, input.Text())
			}

		}(name, host)
	}
	for {
		//infinite
	}
}
