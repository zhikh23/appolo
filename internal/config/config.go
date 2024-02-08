package config

type Config struct {
	Env  		string		`yaml:"env" env-default:"local"`
	Postgres 	Postgres	`yaml:"postgres"`
	Server      Server		`yaml:"server"`
}

type Server struct {
	Port 		int 		`yaml:"port"`
}

type Postgres struct {
	Port     	string		`yaml:"port"`
	Host     	string		`yaml:"host"`
	User    	string		`yaml:"user"`
	Password 	string		`env:"POSTGRES_PASSWORD"`
	DbName   	string		`yaml:"db_name"`
	SslMode  	string		`yaml:"ssl_mode"`
}