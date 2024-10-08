package main

//Напишите структуру, которая будет реализовывать клиент для клиента банковского приложения. Этот клиент должен реализовывать следующий интерфейс:

// type BankClient interface {
// 	// Deposit deposits given amount to clients account
// 	Deposit(amount int)

// 	// Withdrawal withdraws given amount from clients account.
// 	// return error if clients balance less the withdrawal amount
// 	Withdrawal(amount int) error

// 	// Balance returns clients balance
// 	Balance() int
// }
// Используя этот клиент, создайте консольное приложение, которое:

// В момент старта запускает 10 горутин, каждая из которых с промежутком от 0.5 секунд до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10.
// Так же запускается 5 горутин, которые с промежутком 0.5 секунд до 1 секунды снимают с клиента случайную сумму от 1 до 5. Если снятие невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.
// Если пользователь введет в консоль слово balance — приложение выведет в консоль текущий баланс клиента.
// deposit — запрашивается сумма (целое число) — которая добавляется на счёт пользователя
// withdrawal — запрашивается сумма (целое число) — которая выводится со счёта пользователя.Если запрашиваемая сумма больше текущего баланса пользователя, то пользователю должно быть показано сообщение о том, что его баланс недостаточен (и, очевидно, операция не должна быть выполнена).
// Если пользователь введет слово exit — приложение завершит работу.
// При вводе любой другой команды приложение выведет сообщение "Unsupported command. You can use commands: balance, deposit, withdrawal, exit".

func main() {

}

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account. Снимает деньги со счета клиента
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}
