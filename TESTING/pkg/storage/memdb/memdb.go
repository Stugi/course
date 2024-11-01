package memdb

import "app/testing/pkg/storage"

// Пользовательский тип данных - реализация БД в памяти.
// Т.н. "заглушка".
type DB []storage.Task

// Выполнение контракта интерфейса storage.Interface
func (db *DB) Tasks() ([]storage.Task, error) {
	return *db, nil
}
func (db *DB) NewTask(task storage.Task) (int, error) {
	*db = append(*db, task)
	return 0, nil
}
