package system

import (
	"encoding/json"
	"os"
)

func GetUser(id string) (*User, error) {
	path := "/data/users/" + id + ".json"
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = json.Unmarshal(b, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
