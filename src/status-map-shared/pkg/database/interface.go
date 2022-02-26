package database

type DatabaseConfig interface {
	GetDatabaseName() string
	GetConnectionString(bool) string
}
