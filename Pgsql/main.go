package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// структуры обычно используются для
// описания модели данных сущностей,
// хранящихся в БД
type user struct {
	id   int
	name string
}

func main() {
	// Объект БД - пул подключений к СУБД.
	// БД - долгоживущий объект. Следует создавать только один объект для каждой БД.
	// Далее этот объект следует передавать как зависимость.
	var db *sql.DB
	var err error

	// Подключение к БД.
	// В зависимости от драйвера, sql.Open может не выполнять фактического подключения,
	// а только проверить параметры соединения с БД.
	pwd := os.Getenv("dbpass")
	db, err = sql.Open("mysql", "user:"+pwd+"@tcp(server.domain:3306)/tasks")
	if err != nil {
		log.Fatal(err)
	}
	// Не забываем очищать ресурсы.
	defer db.Close()
	// Проверка соединения с БД. На случай, если sql.Open этого не делает.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	data := []user{
		{name: "Rob Pike"},
		{name: "Ken Thompson"},
	}
	err = addUsers(db, data)
	if err != nil {
		log.Fatal(err)
	}
	users, err := users(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}

// addUsers создает таблицу пользователей и заполняет данными.
func addUsers(db *sql.DB, users []user) error {
	// запрос на создание таблицы
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
    		name TEXT NOT NULL
		);
	`)
	// не забываем проверять ошибки
	if err != nil {
		return err
	}
	for _, u := range users {
		// запрос на вставку данных
		_, err := db.Exec(`
		INSERT INTO users (name)
		VALUES (?);
		`,
			u.name,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// users возвращает всех пользователей.
func users(db *sql.DB) ([]user, error) {
	// запрос на выборку данных
	rows, err := db.Query(`
		SELECT * FROM users ORDER BY id;
	`)
	if err != nil {
		return nil, err
	}
	var users []user
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var u user
		err = rows.Scan(
			&u.id,
			&u.name,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		users = append(users, u)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return users, rows.Err()
}

// addUsersPrepared добавляет пользователей в БД.
// Используется подготовленный запрос.
func addUsersPrepared(db *sql.DB, users []user) error {
	// подготовка запроса
	stmt, err := db.Prepare(`INSERT INTO users (name)VALUES (?);`)
	if err != nil {
		return err
	}
	for _, u := range users {
		// выполнение подготовленного запроса на вставку данных
		_, err = stmt.Exec(u.name)
		if err != nil {
			return err
		}
	}
	return nil
}

// Давайте сгруппируем операции по созданию пользователей в транзакцию.
// Таким образом мы сможем гарантировать, что либо все пользователи появятся в БД, либо ни одного.
// addUsersTx добавляет пользователей в БД.
// Используется транзакция.
func addUsersTx(db *sql.DB, users []user) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// tx - объект транзакции; позволяет управлять ее работой
	for _, u := range users {
		// запрос на вставку данных
		_, err := tx.Exec(`
		INSERT INTO users (name)
		VALUES (?);
		`,
			u.name,
		)
		if err != nil {
			// откат транзакции в случае ошибки
			tx.Rollback()
			return err
		}
	}
	// фиксация (подтверждение) транзакции
	tx.Commit()
	return nil
}
