package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

const addr = ":12345"
const network = "tcp4"

// Тип данных для RPC сервера.
// Может быть любым пользовательским типом.
// Все экспортируемые методы этого типа
// будут доступны для удаленного вызова.
type Server int

// Метод Time доступен для удаленного вызова.
func (s *Server) Time(msg string, resp *string) error {
	if msg != "time" {
		return errors.New("неизвестная команда")
	}
	*resp = time.Now().String()
	return nil
}

func main() {
	// Создаем указатель на переменную типа Server.
	srv := new(Server)
	// Регистрируем методы типа Server в службе RPC.
	rpc.Register(srv)
	// Регистрируем HTTP-обработчик для службы RPC.
	rpc.HandleHTTP()
	// Создаём сетевую службу.
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем HTTP-сервер поверх созданной сетевой службы.
	http.Serve(listener, nil)
}
