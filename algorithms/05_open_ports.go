package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "127.0.0.1"

	for port := 1; port <= 65535; port++ {
		go func(p int) {
			addr := fmt.Sprintf("%s:%d", target, p)
			conn, err := net.DialTimeout("tcp", addr, 200*time.Millisecond)
			if err == nil {
				fmt.Printf("Port %d is OPEN\n", p)
				conn.Close()
			}
		}(port)

		// Крошечная пауза, чтобы Федора не заблокировала сканер за спам
		time.Sleep(100 * time.Microsecond)
	}

	// Ждем завершения последних проверок
	time.Sleep(2 * time.Second)
}
