package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Новые параметры подключения к базе данных PostgreSQL
// POSTGRESQL_HOST=45.10.43.153
// POSTGRESQL_PORT=5432
// POSTGRESQL_USER=gen_user
// POSTGRESQL_PASSWORD=g!AVY93W<$}d&x
// POSTGRESQL_DBNAME=default_db

func ConnectDB() (*sql.DB, error) {
	// Строка подключения к базе данных PostgreSQL
	connectionString := "postgres://gen_user:g%21AVY93W%3C%24%7Dd%26x@45.10.43.153:5432/default_db"

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Проверяем, что соединение установлено успешно
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Возвращаем объект соединения
	return db, nil
}
