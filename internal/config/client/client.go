package client

type ClientList struct {
	Omdb Client
}

type Client struct {
	Endpoint string
	Url      string
	ApiKey   string
}
