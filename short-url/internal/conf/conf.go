package conf

type ServerConfig struct {
	DataBase DataBase
}

type DataBase struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
	Pref string `json:"pref"` // 表前缀
}
