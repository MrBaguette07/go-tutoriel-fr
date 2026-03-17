package models

import "fmt"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("le nom est requis")
	}
	if u.Email == "" {
		return fmt.Errorf("l'email est requis")
	}
	return nil
}


