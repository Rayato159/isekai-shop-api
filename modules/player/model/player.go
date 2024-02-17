package model

import "time"

type (
	Player struct {
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		Username  *string   `json:"username"`
		Name      string    `json:"name"`
		Avatar    string    `json:"picture"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	CreatePlayerReq struct {
		ID     string
		Email  string
		Name   string
		Avatar string
	}

	EditPlayerReq struct {
		Username string `json:"username" validate:"omitempty,max=64"`
	}
)
