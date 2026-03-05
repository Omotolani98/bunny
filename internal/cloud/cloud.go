package cloud

type Cloud interface {
	Auth(token string) (string, error)
}
