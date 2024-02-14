package model

import "time"

type (
	PlayerProfile struct {
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		Username  *string   `json:"username"`
		Name      string    `json:"name"`
		Avatar    string    `json:"picture"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	UpdatePlayerProfile struct {
		Username string `json:"username"`
	}
)
