package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

// Сетевой адрес.
//
// Служба будет слушать запросы на всех IP-адресах
// компьютера на порту 12345.
// Например, 127.0.0.1:12345
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const proto = "tcp4"

var proverbs = []string{
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"Don't communicate by sharing memory, share memory by communicating.",
	"The bigger the interface, the weaker the abstraction.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Error handling is important.",
}

func main() {
	// Запуск сетевой службы по протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Подключения обрабатываются в бесконечном цикле.
	// Иначе после обслуживания первого подключения сервер
	//завершит работу.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

// Обработчик. Вызывается для каждого соединения.
func handleConn(conn net.Conn) {
	// Закрытие соединения.
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	for {
		proverb := proverbs[rand.Intn(len(proverbs))]
		conn.Write([]byte(proverb + "\n"))
		time.Sleep(3 * time.Second)
	}
}
