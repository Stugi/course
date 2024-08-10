package main

import (
	"fmt"
	"time"
)

// Semaphore — структура двоичного семафора
type Semaphore struct {
	// Семафор — абстрактный тип данных,
	// в нашем случае в основе его лежит канал
	sem chan int
	// Время ожидания основных операций с семафором, чтобы не
	// блокировать
	// операции с ним навечно (необязательное требование, зависит от
	// нужд программы)
	timeout time.Duration
}

// Acquire — метод захвата семафора
func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- 0:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("Не удалось захватить семафор")
	}
}

// Release — метод освобождения семафора
func (s *Semaphore) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("Не удалось освободить семафор")
	}
}

// NewSemaphore — функция создания семафора
func NewSemaphore(timeout time.Duration, counter int) *Semaphore {
	return &Semaphore{
		sem:     make(chan int, counter),
		timeout: timeout,
	}
}

func main() {
	s := NewSemaphore(3 * time.Second)

	// Начинаем работать с разделяемыми данными,
	// поэтому пытаемся захватить семафор
	if err := s.Acquire(); err != nil {
		// Код, в случае если кто-то уже работает с разделяемыми
		// данными
	}

	// Выполняем важную работу с разделяемым ресурсом

	// Освобождаем семафор
	if err := s.Release(); err != nil {
		// Код, в случае если по ошибке освобождаем свободный
		// семафор
	}
}
