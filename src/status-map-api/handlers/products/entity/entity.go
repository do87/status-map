package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/do87/status-map/src/status-map-shared/pkg/common"
	"github.com/lib/pq"
)

// Product model
type Product struct {
	ID           string `gorm:"primaryKey"`
	Org          string `gorm:"primaryKey;default:sit"`
	Platform     string
	Infras       pq.StringArray `gorm:"type:text[]"`
	Stages       pq.StringArray `gorm:"type:text[]"`
	StatusInfra  StatusMap
	StatusStages StatusMap
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// StatusMap represent a map of recent statuses
type StatusMap map[string]Status

// Status represent status of a run
type Status struct {
	CurrentStatus string
	UpdatedAt     time.Time
}

// Value for postgres StringArray
func (a StatusMap) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan parses StringArray value
func (a *StatusMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(common.ERR_WRONG_TYPE)
	}
	return json.Unmarshal(b, &a)
}

// TableName to ensure correct table naming
func (Product) TableName() string {
	return "products"
}
