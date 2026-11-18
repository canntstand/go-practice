package main

import (
	"fmt"
)

type Server struct {
	Name   string
	Host   string
	Port   int32
	Online bool
}

func ShowAllServers(s []Server) {
	for i := 0; i < len(s); i++ {
		status := "DOWN"

		if s[i].Online == true {
			status = "UP"
		}

		fmt.Printf("[%d] %s | %s:%d | %s\n", i, s[i].Name, s[i].Host, s[i].Port, status)
	}
}

func AddServer(s []Server) []Server {
	var name string
	var host string
	var port int32
	var online bool

	_, err := fmt.Scan(&name, &host, &port, &online)

	if err != nil {
		fmt.Println("Ошибка")
		return s
	}

	s = append(s, Server{name, host, port, online})
	return s
}

func ChangeServerStatus(s []Server) {
	var number int8

	_, err := fmt.Scan(&number)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	s[number].Online = !s[number].Online
	return

}

func ShowStatistics(s []Server) {
	totalServers := len(s)
	online := 0
	offline := 0

	for i := 0; i < len(s); i++ {
		if s[i].Online == true {
			online++
		} else {
			offline++
		}
	}

	fmt.Printf(`Total servers: %d
Online: %d
Offline: %d`, totalServers, online, offline)
}

func main() {
	servers := []Server{
		{Name: "Web-01", Host: "192.168.1.10", Port: 443, Online: true},
		{Name: "DB-01", Host: "192.168.1.20", Port: 5432, Online: false},
		{Name: "VPN-01", Host: "10.0.0.1", Port: 51820, Online: true},
	}

	for {
		fmt.Printf(`===== Server Monitor =====
1. Show all servers
2. Add server
3. Change server status
4. Show statistics
5. Exit`)
		var option int8

		_, err := fmt.Scan(&option)

		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Ошибка: введите число от 1 до 5.")
			continue
		}

		switch option {
		case 1:
			ShowAllServers(servers)
		case 2:
			servers = AddServer(servers)
		case 3:
			ChangeServerStatus(servers)
		case 4:
			ShowStatistics(servers)
		case 5:
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Неизвестная команда. Выберите пункт от 1 до 5.")
		}
	}
}
