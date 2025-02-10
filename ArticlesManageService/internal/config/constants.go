package config

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var (
	ArticlesAccessService_url = /*os.Getenv("ArticlesAccessServiceURL")*/ "url"
)
