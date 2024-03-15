package exception

type Oauth2Exception struct{}

func (e *Oauth2Exception) Error() string {
	return "OAuth2 failed to authenticate user"
}
