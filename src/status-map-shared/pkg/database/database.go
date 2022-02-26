package database

import (
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service is the database service struct
type Service struct {
	isLocal bool
	conn    string
	db      *gorm.DB
}

// New returns a new database service
func New(cfg DatabaseConfig, isLocal bool) (*Service, error) {
	d := &Service{
		conn:    cfg.GetConnectionString(true),
		isLocal: isLocal,
	}

	db, err := gorm.Open(postgres.Open(d.conn), getGormConfig(d.isLocal))
	connectErr, ok := errors.Unwrap(err).(*pgconn.PgError)
	if ok && connectErr.Code == ERR_CODE_NO_DB {
		if err := d.CreateDatabase(cfg); err != nil {
			return nil, errors.Wrap(err, "Unable to create missing database")
		}
		db, err = gorm.Open(postgres.Open(d.conn), getGormConfig(d.isLocal))
	}

	if err != nil {
		return nil, fmt.Errorf("error opening to DB: %s", err)
	}

	d.db = db
	return d, nil
}

// CreateDatabase creates DB
func (d *Service) CreateDatabase(cfg DatabaseConfig) error {

	dbConn, err := gorm.Open(postgres.Open(cfg.GetConnectionString(false)), getGormConfig(d.isLocal))
	if err != nil {
		return err
	}

	tx := dbConn.Exec(fmt.Sprintf("CREATE DATABASE %s;", cfg.GetDatabaseName()))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetDB returns a pointer to the DB
func (c *Service) GetDB() *gorm.DB {
	return c.db
}

// Close closes connection to the DB
func (d *Service) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// MigrateTables runs table migrations
func (d *Service) MigrateTables(tables ...interface{}) error {
	if err := d.db.AutoMigrate(tables...); err != nil {
		return err
	}

	for _, table := range tables {
		for counter := 0; counter < 20 && !d.db.Migrator().HasTable(table); counter++ {
			time.Sleep(3 * time.Second)
		}
		if !d.db.Migrator().HasTable(table) {
			return fmt.Errorf("table '%s' creation took too long", table)
		}
	}

	return nil
}

// MakeRecreateConstraints drop and recreate constraints to perform migration
func (d *Service) MakeRecreateConstraints() func(table interface{}, constraint string) error {
	return func(table interface{}, constraint string) error {
		if err := d.db.Migrator().DropConstraint(table, constraint); err != nil {
			return err
		}
		if err := d.db.Migrator().CreateConstraint(table, constraint); err != nil {
			return err
		}
		return nil
	}
}
