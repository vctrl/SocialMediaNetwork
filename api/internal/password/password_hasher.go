package password

type PasswordHasher interface {
	HashPass(plainPass string) ([]byte, error)
	CheckPass(passHash []byte, plain string) error
}

type PasswordHasherImpl struct {
}
