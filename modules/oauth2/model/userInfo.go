package model

type (
	UserInfo struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}

	PlayerPassport struct {
		RefreshToken string `json:"refresh_token"`
	}

	CreatePlayerInfo struct {
		ID      string
		Email   string
		Name    string
		Picture string
		*PlayerPassport
	}
)
