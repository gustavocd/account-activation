package models

import (
	"crypto/rand"
	"encoding/base64"
)

// ActivationToken represents user's token.
type ActivationToken struct {
	ID    int
	Token string
}

// GenerateToken generates a token for and specific user's id
func (at *ActivationToken) GenerateToken(lastInsertID int) error {
	//generate the token string
	token, _ := generateRandomString(32)
	// associate the token with user's id passed as a parameter
	query := `INSERT INTO activation_tokens(user_id, token) VALUES($1, $2);`
	_, err := DB.Exec(query, lastInsertID, token)
	if err != nil {
		return err
	}
	return nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
