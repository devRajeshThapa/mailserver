package main

import (
    "fmt"
    "net"
    "bufio"
)

func main() {
    // Listen on TCP port 9000 on all interfaces
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        fmt.Println("Error starting TCP server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening on port 9000...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn) // handle each client concurrently
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Client connected:", conn.RemoteAddr())

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()
        fmt.Println("Received:", text)
        conn.Write([]byte("Echo: " + text + "\n")) // send response
    }
}

