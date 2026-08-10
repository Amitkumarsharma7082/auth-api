package db

import (
	"auth-api/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("Connect DB")
	db, err := sql.Open(
		"mysql",                              // driver
		"root:@tcp(127.0.0.1:3306)/auth_api", // mysql -u root, tcp -. connection and 127(computer address)
	)
	/*
	   User: root
	   Password: empty
	   Host: 127.0.0.1
	   Port: 3306
	   Database: auth_api
	*/
	if err != nil {
		return nil, err
	}

	err = db.Ping() // mysql server is available

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateUser(database *sql.DB, user model.User) error {
	query := `
		INSERT INTO users(name, email, password)
		VALUES (?, ?, ?)
	`
	_, err := database.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

/*
SELECT id, name, email, password
FROM users
WHERE email = "aksharma@gmail.com"
*/
func GetUserByEmail(database *sql.DB, email string) (model.User, error) {
	var user model.User

	query := `
	SELECT id, name, email, password
	FROM users
	WHERE email = ?`

	err := database.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserById(database *sql.DB, id int) (model.User, error) {
	var user model.User

	query := `
	SELECT id, name, email, password
	FROM users
	WHERE id = ?
	`
	err := database.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return user, err
	}
	return user, nil
}


