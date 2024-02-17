package model

type (
	// This struct is provided by Google OAuth2
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

	CreatePlayerInfo struct {
		ID     string
		Email  string
		Name   string
		Avatar string
	}

	CreateAdminInfo struct {
		ID     string
		Email  string
		Name   string
		Avatar string
	}
)
