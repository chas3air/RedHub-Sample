package config

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var (
	UsersAccessServiceURL = /*os.Getenv("UsersAccessServiceURL")*/ "url"
)
