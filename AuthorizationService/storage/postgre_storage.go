package storage

import (
	"authorization-service/models"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

// Реализуем метод CreateUser для структуры PostgresStorage
func (pg *PostgresStorage) CreateUser(user models.User) {
	// Используем параметризованный запрос для безопасности
	query := "INSERT INTO Users (id, login, name, pass, role) VALUES ($1, $2, $3, $4, $5)"
	_, err := pg.db.Exec(query, user.Id, user.Login, user.Name, user.Password, user.Role)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	fmt.Printf("User %s created successfully\n", user.Name)
}

func NewAuthStorage(host, port, user, password, dbname string) (UserStorage, error) {
	var ps PostgresStorage
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	var err error
	ps.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return &ps, nil
}

func (ps *PostgresStorage) FindUserById(id string) (user models.User, exis bool) {
	qery := fmt.Sprintf(`select * from Users where Id = '%s'`, id)
	row := ps.db.QueryRow(qery)

	u := models.User{}
	fmt.Println(qery)
	err := row.Scan(&u.Id, &u.Login, &u.Password, &u.Name, &u.Role)

	if err != nil {
		return u, false
	}

	return u, true
}

func (ps *PostgresStorage) FindUserByLogin(login string) (user models.User, exis bool) {
	qery := fmt.Sprintf(`select * from Users where Login = '%s'`, login)
	row := ps.db.QueryRow(qery)

	u := models.User{}
	fmt.Println(qery)
	err := row.Scan(&u.Id, &u.Login, &u.Password, &u.Name, &u.Role)

	if err != nil {
		return u, false
	}

	return u, true
}

func (ps *PostgresStorage) GetAllUsersJSON() (string, bool) {
	var jsonData string
	err := ps.db.QueryRow(`SELECT json_agg(t) FROM (SELECT * FROM Users) t`).Scan(&jsonData)
	if err != nil {
		fmt.Println("Пиздец, а не все пользователи!")
		return jsonData, false
	}
	fmt.Println(jsonData)
	return jsonData, true
}

func (ps *PostgresStorage) DeleteUser(id int) {
	fmt.Println(id)
	qery := fmt.Sprintf("Delete from Users where id = %s", strconv.Itoa(id))
	ps.db.Exec(qery)
	return
}
