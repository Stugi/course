package main

import (
	bank "bank/type"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	// source := rand.NewSource(time.Now().Unix())
	// r := rand.New(source)
	rand.Seed(time.Now().Unix())
}

func main() {
	var command string
	client := &bank.Client{}

	var m sync.RWMutex
	var rs, ws int
	rsCh := make(chan int)
	wsCh := make(chan int)

	go func() {
		for {
			select {
			case n := <-rsCh:
				rs += n
			case n := <-wsCh:
				ws += n
			}
		}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addMoney(wsCh, &m, &wg, client, rand.Intn(10))
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go minusMoney(wsCh, &m, &wg, client, rand.Intn(10))
	}

	fmt.Println("You can enter commands: balance, deposit, withdrawal, exit")
	_, err := fmt.Scanf("%s", &command)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	switch command {
	case "balance":
		wg.Add(1)
		fmt.Println("Balance:", getBalance(rsCh, &m, &wg, client))
	case "withdrawal":
		fmt.Println("Enter withdrawal:")
		var withdrawal int
		_, err := fmt.Scanf("%v", &withdrawal)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		wg.Add(1)
		go minusMoney(wsCh, &m, &wg, client, withdrawal)
	case "deposit":
		fmt.Println("Enter deposit:")
		var deposit int
		_, err := fmt.Scanf("%v", &deposit)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		wg.Add(1)
		go addMoney(wsCh, &m, &wg, client, deposit)
	case "exit":
		return
	default:
		fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
	}

	// b := make([]byte, 1)
	// os.Stdin.Read(b)

	wg.Wait()
}

func sleep(min, max int) {
	time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)
}
func addMoney(c chan int, m *sync.RWMutex, wg *sync.WaitGroup, client *bank.Client, deposit int) {
	sleep(500, 1000)
	m.Lock()
	c <- 1
	client.Deposit(deposit)
	c <- -1
	m.Unlock()
	wg.Done()
}

func minusMoney(c chan int, m *sync.RWMutex, wg *sync.WaitGroup, client *bank.Client, withdrawal int) {
	sleep(500, 1000)
	m.Lock()
	c <- 1
	client.Withdrawal(withdrawal)
	c <- -1
	m.Unlock()
	wg.Done()
}

func getBalance(c chan int, m *sync.RWMutex, wg *sync.WaitGroup, client *bank.Client) int {
	defer func() {
		c <- -1
		m.RUnlock()
		wg.Done()
	}()

	m.RLock()
	c <- 1
	return client.Balance()
}
