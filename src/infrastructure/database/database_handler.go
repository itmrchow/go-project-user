package database

type DatabaseHandler interface {
	Connect() error
	Migrate() error
}
