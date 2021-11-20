package db

// DatabaseList is struct for list of database
type DatabaseList struct {
	MovieImdb struct {
		Mysql Database
	}
}

// Database is struct for Database conf
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}
