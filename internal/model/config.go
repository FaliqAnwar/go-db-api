package model

type (
	Config struct {
		App App `json:"app"`
		DB  DB  `json:"db"`
	}

	App struct {
		Env  string `json:"env"`
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	}

	DB struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		SslMode  string `json:"ssl_mode"`
	}
)
