package factory

import "github.com/google/uuid"

func NewImageID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
