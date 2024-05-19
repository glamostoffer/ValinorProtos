package chat

type Config struct {
	Address string `validate:"required"`
	DontRun bool
}
