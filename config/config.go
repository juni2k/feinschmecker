package config

type Config struct {
	Operator   string `json:"operator"`
	Repository string `json:"repository"`
	Workdir    string `json:"workdir"`
	Telegram   struct {
		Token string `json:"token"`
	} `json:"telegram"`
}
