package auth

type AuthService struct {
	Secret string
}

func New(secret string) *AuthService {
	return &AuthService{
		Secret: secret,
	}
}

func (a *AuthService) Authorize() {

}

func (a *AuthService) ValidateToken() {

}

func (a *AuthService) UserAuth() {

}

func (a *AuthService) AdminAuth() {

}

func (a *AuthService) GenerateToken() {

}

func (a *AuthService) HashPassword() {

}

func (a *AuthService) ComparePassword() {

}
