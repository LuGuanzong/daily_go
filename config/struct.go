package config

type basic struct {
	Env  string `json:"env" yaml:"env"`
	Port string `json:"port" yaml:"port"`
}

type mysql struct {
	Driver        string `json:"driver" yaml:"driver"`
	Host          string `json:"host" yaml:"host"`
	Port          int    `json:"port" yaml:"port"`
	User          string `json:"user" yaml:"user"`
	Password      string `json:"password" yaml:"password"`
	Db            string `json:"db" yaml:"db"`
	DBPoolMaxIdle int    `json:"dbPoolMaxIdle" yaml:"db_pool_max_idle"`
	DBPoolMax     int    `json:"dbPoolMax" yaml:"db_pool_max"`
}

type sqlite struct {
	Pwd string `json:"pwd" yaml:"pwd"`
}

type log struct {
	Dir        string `json:"dir" yaml:"dir"`
	Filename   string `json:"filename"  yaml:"filename"`
	Level      string `json:"level" yaml:"level"`
	MaxSize    int    `json:"MaxSize"  yaml:"max_size"`
	MaxBackups int    `json:"MaxBackups"  yaml:"max_backups"`
	MaxAge     int    `json:"MaxAge"  yaml:"max_age"`
	Compress   bool   `json:"Compress"  yaml:"compress"`
}

type Config struct {
	Basic  basic  `json:"basic" yaml:"basic"`
	Log    log    `json:"log" yaml:"log"`
	Mysql  mysql  `json:"mysql" yaml:"mysql"`
	Sqlite sqlite `json:"sqlite" yaml:"sqlite"`
}
