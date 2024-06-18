package logger

import "database/sql"

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) StoreStatus() error {
	// _, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
	// if err != nil {
	// 	return err
	// }

	/*
		LOS Log Storage Model
		ID - Do I even need this? - I was thinking of adding a way to add comment via ui
		TIMESTAMP
		ISCONNECTED (TRUE/FALSE)
	*/

	return nil
}
