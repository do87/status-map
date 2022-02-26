package config

import (
	"errors"
	"fmt"
	"strings"
)

// Config is the app configuration
type Config struct {
	Server   Server
	Database Database
	IAP      IAP
}

// Server setup details
type Server struct {
	BindAddr string
	Hostname string
	HttpPort int
}

// Database db connection information
type Database struct {
	Host    string
	User    string
	Pass    string
	Name    string
	SSLMode string
	Local   bool
}

type IAP struct {
	ClinetID string
}

// FromEnv loads environment variables into app config
func FromEnv() (*Config, error) {
	h := NewEnvHandler()

	cfg := &Config{}
	cfg.Server.BindAddr = h.GetString("ODJ_EE_BIND_ADDRESS", "0.0.0.0")
	cfg.Server.HttpPort = h.GetInt("ODJ_EE_HTTP_PORT", 8080)
	cfg.Server.Hostname = h.GetOsHostname()

	cfg.Database.Name = h.GetString("ODJ_DEP_STATUS_MAP_DB_DATABASE")
	cfg.Database.Host = h.GetString("ODJ_DEP_STATUS_MAP_DB_HOST")
	cfg.Database.User = h.GetString("ODJ_DEP_STATUS_MAP_DB_USER")
	cfg.Database.Pass = h.GetString("ODJ_DEP_STATUS_MAP_DB_PASSWORD")
	cfg.Database.SSLMode = h.GetString("DB_SSL_MODE", "disable")
	cfg.Database.Local = cfg.IsLocal()

	cfg.IAP.ClinetID = h.GetString("ODJ_DEP_ODJ_CORE_API_AUDIENCE", "")

	if len(h.Errors) > 0 {
		return nil, errors.New(strings.Join(h.Errors, "\n"))
	}

	return cfg, nil
}

func (c *Config) IsLocal() bool {
	h := NewEnvHandler()
	return h.GetString("LOCAL", "false") == "true"
}

// GetConnectionString build db connection string
// set useProvidedDBName to false in cases where it's needed to connect to the default 'postgres' DB
func (d Database) GetConnectionString(useProvidedDBName bool) string {
	user := d.User
	name := d.Name
	if !useProvidedDBName {
		name = "postgres"
	}

	c := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host,
		user,
		d.Pass,
		name,
		d.SSLMode,
	)

	return c
}

// GetDatabaseName returns database name
func (d Database) GetDatabaseName() string {
	return d.Name
}
