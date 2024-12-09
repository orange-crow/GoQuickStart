package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Starts a P2P node as either a server or a client
func main() {
	// 检查命令行参数，决定是启动为服务器还是客户端。
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client] [host:port]")
		return
	}

	mode := os.Args[1]
	hostPort := os.Args[2]

	if mode == "server" {
		startServer(hostPort)
	} else if mode == "client" {
		startClient(hostPort)
	} else {
		fmt.Println("Invalid mode. Use 'server' or 'client'.")
	}
}

// Starts the server node
func startServer(hostPort string) {
	listener, err := net.Listen("tcp", hostPort) // 创建一个TCP监听器
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running on", hostPort)

	for {
		conn, err := listener.Accept() // 等待客户端连接
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn) // 创建协程(goroutine)的关键字: 接收到连接后，启动新协程处理通信。
	}
}

// Handles incoming connections to the server
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			return
		}
		fmt.Printf("Received: %s", message)

		response := strings.ToUpper(message)
		conn.Write([]byte(response)) // 写入消息: 发送字节数据到客户端
	}
}

// Starts the client node
func startClient(hostPort string) {
	conn, err := net.Dial("tcp", hostPort) // 使用 net.Dial 建立 TCP 连接
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server", hostPort)

	reader := bufio.NewReader(os.Stdin) // 读取消息: 使用 bufio.NewReader 包装连接
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message)) // 写入消息:将输入的消息发送给服务器

		response, err := bufio.NewReader(conn).ReadString('\n') //读取消息: 表示读取一行，直到遇到换行符 \\n
		if err != nil {
			fmt.Println("Server disconnected")
			return
		}
		fmt.Printf("Response: %s", response)
	}
}
