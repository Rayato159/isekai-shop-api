package model

type (
	LoginResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	LogoutResponse struct {
		Message string `json:"message"`
	}
)
