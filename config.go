package auth

type Config struct {
	Address string `validate:"required"`
	DontRun bool
}
