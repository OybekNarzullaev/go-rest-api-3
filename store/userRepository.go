package store

import (
	"github.com/OybekNarzullaev/http-rest-api/model"
)

// user repository
type UserRepositry struct{
	store *Store
}

// create user..
func (r *UserRepositry) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	
	err := r.store.db.QueryRow(`INSERT INTO users (email, encrypted_password) VALUES($1, $2) RETURNING id`,
	u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}


	return u, nil	
}

// find user by email...
func (r *UserRepositry) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(`SELECT id, email, encrypted_password FROM users WHERE email=$1`, 
	email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)
	if err != nil {
		return nil, err
	}

	return u, nil	
}