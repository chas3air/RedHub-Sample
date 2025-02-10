package config

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var (
	UsersAccessService_url = /*os.Getenv("UsersAccessServiceURL")*/ "url"
)
