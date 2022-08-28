package domain

import "time"

type (
	User struct {
		UUID      string     `db:"uuid"`
		Email     string     `db:"email"`
		LastName  string     `db:"last_name"`
		FirstName string     `db:"first_name"`
		Gender    string     `db:"gender"`
		CreatedAt *time.Time `db:"created_at"`
		UpdatedAt *time.Time `db:"updated_at"`
		DeletedAt *time.Time `db:"deleted_at"`
	}
)
