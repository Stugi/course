package db

import "sync"

// Заказ на доставку товаров.
type Order struct {
	ID              int       // номер заказа
	IsOpen          bool      // открыт/закрыт
	DeliveryTime    int64     // срок доставки
	DeliveryAddress string    // адрес доставки
	Products        []Product // состав заказа
}

// Товар.
type Product struct {
	ID    int     // артикул товара
	Name  string  // название
	Price float64 // цена
}

// База данных заказов.
type DB struct {
	m     sync.Mutex    //мьютекс для синхронизации доступа
	id    int           // текущее значение ID для нового заказа
	store map[int]Order // БД заказов
}

// Orders возвращает все заказы.
func (db *DB) Orders() []Order {
	db.m.Lock()
	defer db.m.Unlock()
	var data []Order
	for _, v := range db.store {
		data = append(data, v)
	}
	return data
}

// NewOrder создает новый заказ.
func (db *DB) NewOrder(o Order) int {
	db.m.Lock()
	defer db.m.Unlock()
	o.ID = db.id
	db.store[o.ID] = o
	db.id++
	return o.ID
}

// UpdateOrder обновляет данные заказа по ID.
func (db *DB) UpdateOrder(o Order) {
	db.m.Lock()
	defer db.m.Unlock()
	db.store[o.ID] = o
}

// DeleteOrder удаляет заказ по ID.
func (db *DB) DeleteOrder(id int) {
	db.m.Lock()
	defer db.m.Unlock()
	delete(db.store, id)
}
