package model

type (
	PlayerPassport struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	DoRenewToken struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	DoLogout struct {
		AcessToken   string `json:"accessToken" validate:"required"`
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	LoginResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int64  `json:"expiresIn"`
	}

	LogoutResponse struct {
		Message string `json:"message"`
	}
)
