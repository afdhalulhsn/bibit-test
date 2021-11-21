package server

type ServerList struct {
	Grpc Server
}

type Server struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Timeout  int    `yaml:"timeout"`
	RestHost string `yaml:"restHost"`
}
