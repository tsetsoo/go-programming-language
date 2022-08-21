package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	// currentPath := strings.TrimSuffix(filepath.Dir(ex), "/exe")
	currentPath := filepath.Dir(ex)
	for {
		closed := writeToClient(conn, fmt.Sprintf("%s: ", currentPath))
		if closed {
			return
		}
		input := bufio.NewScanner(conn)
		for input.Scan() {
			commandParts := strings.Split(input.Text(), " ")
			switch commandParts[0] {
			case "cd":
				if len(commandParts) == 1 {
					commandParts = append(commandParts, "/")
				}
				currentPath = commandParts[1]
				closed := writeToClient(conn, fmt.Sprintf("%s: ", currentPath))
				if closed {
					return
				}
			case "close":
				writeToClient(conn, "Bye\n")
				conn.Close()
				return
			case "ls":
				out, err := exec.Command("ls", currentPath).Output() // TODO add check if exists
				if err != nil {
					log.Fatal(err)
				}
				closed := writeToClientBytes(conn, out)
				if closed {
					return
				}
				closed = writeToClient(conn, fmt.Sprintf("%s: ", currentPath))
				if closed {
					return
				}
			case "get":
				sendFileToClient(conn, fmt.Sprintf("%s/%s", currentPath, commandParts[1]))
				closed = writeToClient(conn, fmt.Sprintf("%s: ", currentPath))
				if closed {
					return
				}
			}

		}
	}
}

func sendFileToClient(conn net.Conn, fileName string) {
	file, err := os.Open(strings.TrimSpace(fileName)) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // make sure to close the file even if we panic.
	n, err := io.Copy(conn, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, "bytes sent")
}

func writeToClient(conn net.Conn, toWrite string) bool {
	_, err := io.WriteString(conn, toWrite)
	if err != nil {
		if errors.Is(err, syscall.EPIPE) {
			log.Print("Client disconnected")
			return true
		}
		log.Fatal(err)
	}
	return false
}

func writeToClientBytes(conn net.Conn, toWrite []byte) bool {
	_, err := conn.Write(toWrite)
	if err != nil {
		if errors.Is(err, syscall.EPIPE) {
			log.Print("Client disconnected")
			return true
		}
		log.Fatal(err)
	}
	return false
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
