package system

import (
	"encoding/json"
	"os"
)

func GetUserByPhone(phone string) (*User, error) {
	b, err := os.ReadFile("/data/users/phone_index.json")
	if err != nil {
		return nil, err
	}
	phoneIndex := map[string]string{}
	json.Unmarshal(b, &phoneIndex)
	id, ok := phoneIndex[phone]
	if !ok {
		return nil, os.ErrNotExist
	}
	return GetUser(id)
}
