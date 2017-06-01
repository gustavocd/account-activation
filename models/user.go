package models

var db = getConnection()

// User represents user's account.
type User struct {
	ID       int
	Name     string
	LastName string
	Email    string
	Active   bool
}

// Create insert a user in the db and returns the last id inserted.
func (u *User) Create() (int, error) {
	query := "INSERT INTO users(name, lastname,email) VALUES($1, $2, $3) RETURNING id;"
	lastInsertID := 0
	err := db.QueryRow(query, u.Name, u.LastName, u.Email).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}
