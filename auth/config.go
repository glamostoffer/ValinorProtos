package auth

type Config struct {
	Address string `yaml:"address" validate:"required"`
	DontRun bool   `yaml:"dont_run"`
}
