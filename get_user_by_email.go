package system

import (
	"encoding/json"
	"os"
)

func GetUserByEmail(email string) (*User, error) {
	b, err := os.ReadFile("/data/users/email_index.json")
	if err != nil {
		return nil, err
	}
	emailIndex := map[string]string{}
	json.Unmarshal(b, &emailIndex)
	id, ok := emailIndex[email]
	if !ok {
		return nil, os.ErrNotExist
	}
	return GetUser(id)
}
