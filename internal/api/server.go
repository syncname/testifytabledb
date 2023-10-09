package api

import (
	"database/sql"
	"github.com/syncname/testifyexample/internal/models"
)

type Server struct {
	db *sql.DB
}

func (s *Server) CreateUser(u *models.User) (*models.User, error) {

	createUserQuery := "insert into users(name,email) values($1, $2) returning *"

	stm, err := s.db.Prepare(createUserQuery)
	if err != nil {
		return nil, err
	}

	res := stm.QueryRow(u.Name, u.Email)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var resUser models.User
	err = res.Scan(&resUser.Name, &resUser.Email, &resUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}

func (s *Server) GetUser(name string) (*models.User, error) {

	getUserQuery := "select * from users where name = $1"

	stm, err := s.db.Prepare(getUserQuery)
	if err != nil {
		return nil, err
	}

	res := stm.QueryRow(name)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var resUser models.User
	err = res.Scan(&resUser.Name, &resUser.Email, &resUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}

func (s *Server) DeleteUser(name string) (*models.User, error) {

	deleteUserQuery := "delete from users where name = $1 returning *"

	stm, err := s.db.Prepare(deleteUserQuery)
	if err != nil {
		return nil, err
	}

	res := stm.QueryRow(name)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var resUser models.User
	err = res.Scan(&resUser.Name, &resUser.Email, &resUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}

func (s *Server) UpdateUserEmail(u *models.User) (*models.User, error) {

	updateUserQuery := "update users SET email = $1 where name = $2 RETURNING *"

	stm, err := s.db.Prepare(updateUserQuery)
	if err != nil {
		return nil, err
	}

	res := stm.QueryRow(u.Email, u.Name)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var resUser models.User
	err = res.Scan(&resUser.Name, &resUser.Email, &resUser.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}
