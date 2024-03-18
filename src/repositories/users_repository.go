package repositories

import (
	"api/src/models"
	"database/sql"
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
