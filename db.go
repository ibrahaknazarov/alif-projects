package main

import (
	"fmt"
	"log"
	"main/database"
)

type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type Order struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Amount int    `db:"amount"`
}

func main() {
	database.InitDB()
	db := database.DB

	// 1. Создание таблицы пользователей
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE
	);`
	db.MustExec(schema)

	// 2. Получить всех пользователей
	var users []User
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatal("Get all users:", err)
	}
	fmt.Println("Users:", users)

	// 3. Добавить нового пользователя
	_, err = db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", "Ali", "ali@example.com")
	if err != nil {
		log.Println("Insert error:", err)
	}

	// 4. Обновить email пользователя по ID
	_, err = db.Exec("UPDATE users SET email=$1 WHERE id=$2", "newemail@example.com", 1)
	if err != nil {
		log.Println("Update error:", err)
	}

	// 5. Удалить пользователя по email
	_, err = db.Exec("DELETE FROM users WHERE email=$1", "ali@example.com")
	if err != nil {
		log.Println("Delete error:", err)
	}

	// 6. Реализовать GetUserByEmail
	user, err := GetUserByEmail("newemail@example.com")
	if err != nil {
		log.Println("GetUserByEmail error:", err)
	} else {
		fmt.Println("Found user:", user)
	}

	// 7. Создание таблицы заказов и действия
	orderSchema := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		amount INTEGER
	);`
	db.MustExec(orderSchema)

	// Добавить заказ
	_, err = db.Exec("INSERT INTO orders (user_id, amount) VALUES ($1, $2)", 1, 250)
	if err != nil {
		log.Println("Insert order error:", err)
	}

	// Получить все заказы
	var orders []Order
	err = db.Select(&orders, "SELECT * FROM orders")
	if err != nil {
		log.Println("Get all orders error:", err)
	}
	fmt.Println("Orders:", orders)

	// Обновить заказ
	_, err = db.Exec("UPDATE orders SET amount=$1 WHERE id=$2", 300, 1)

	// Удалить заказ
	_, err = db.Exec("DELETE FROM orders WHERE id=$1", 1)
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}
