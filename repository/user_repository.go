package repository

import (
	"REST-API-BookCatalog-Gin/entity"
	"database/sql"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	AddUser(payload entity.User) error
	UpdateUser(payload entity.User) error
	// DeleteUser(id int) (*entity.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) GetUsers() ([]entity.User, error) {
	rows, err := u.DB.Query("select * from tbl_users")

	if err != nil {
		return nil, err
	}

	var users []entity.User

	for rows.Next() {
		var res entity.User
		_ = rows.Scan(&res.Id, &res.Username, &res.Email)
		users = append(users, res)
	}

	return users, nil
}

func (u *userRepository) GetUserByID(id int) (*entity.User, error) {
	sqlStatement := "SELECT * FROM tbl_users WHERE user_id = ?"
	row := u.DB.QueryRow(sqlStatement, id)
	var user entity.User
	err := row.Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (b *userRepository) AddUser(payload entity.User) error {
	_, err := b.DB.Exec("INSERT INTO tbl_users (username, email) VALUES (?, ?)", payload.Username, payload.Email)

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UpdateUser(payload entity.User) error {
	_, err := u.DB.Exec("UPDATE tbl_users SET username = ?, email = ? WHERE user_id = ?", payload.Username, payload.Email, payload.Id)

	if err != nil {
		return err
	}

	return nil
}


// func ( u *userRepository) DeleteUser(id int) (error) {
// 	sqlStatement := "DELETE FROM tbl_users WHERE user_id = ?"
// 	row := u.DB.QueryRow(sqlStatement, id)
// 	var user entity.User
// 	err := row.Scan(&user.Id, &user.Username, &user.Email)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }