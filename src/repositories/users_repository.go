package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *usersRepository {
	return &usersRepository{db}
}

// Create belongs to the userRepository inserts a new user into the database and return its (id,error)
func (repository usersRepository) Create(user models.User) (uint64, error) {

	statement, err := repository.db.Prepare(
		"insert into users (name, nickname,email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository usersRepository) GetUsers(nickname string) ([]models.User, error) {
	formatedNick := fmt.Sprintf("%%%s%%", nickname) //"%nickname%"

	rows, err := repository.db.Query(
		"SELECT id, nickname, email, created_at FROM users WHERE nickname LIKE ?",
		formatedNick,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User

	// scan a row to a user model and appends it to the users slice
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
