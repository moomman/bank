package global

type config struct {
	DbName    string
	DbSource  string
	ServerUrl string
	Secret    string
}

var Config config
