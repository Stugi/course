package storage

// Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Интерфейс БД.
// Этот интерфейс позволяет абстрагироваться от конкретной СУБД.
// Можно создать реализацию БД в памяти для модульных тестов.
type Interface interface {
	Tasks() ([]Task, error)
	NewTask(Task) (int, error)
}
