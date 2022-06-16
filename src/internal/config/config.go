package config

type (
	Config struct {
		App   App   `yaml:"app"`
		Mysql Mysql `yaml:"mysql"`
	}

	App struct {
		Name    string `yaml:"name"`
		Address string `yaml:"port"`
		Secret  string `yaml:"secret"`
	}

	Mysql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	}
)
