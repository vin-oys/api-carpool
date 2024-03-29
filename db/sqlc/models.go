// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Category string

const (
	CategoryAdult Category = "adult"
	CategoryChild Category = "child"
)

func (e *Category) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Category(s)
	case string:
		*e = Category(s)
	default:
		return fmt.Errorf("unsupported scan type for Category: %T", src)
	}
	return nil
}

type NullCategory struct {
	Category Category
	Valid    bool // Valid is true if Category is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCategory) Scan(value interface{}) error {
	if value == nil {
		ns.Category, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Category.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Category), nil
}

type Country string

const (
	CountryMalaysia  Country = "malaysia"
	CountrySingapore Country = "singapore"
)

func (e *Country) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Country(s)
	case string:
		*e = Country(s)
	default:
		return fmt.Errorf("unsupported scan type for Country: %T", src)
	}
	return nil
}

type NullCountry struct {
	Country Country
	Valid   bool // Valid is true if Country is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCountry) Scan(value interface{}) error {
	if value == nil {
		ns.Country, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Country.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCountry) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Country), nil
}

type UserRole string

const (
	UserRoleSuperAdministrator UserRole = "super_administrator"
	UserRoleAdministrator      UserRole = "administrator"
	UserRoleDriver             UserRole = "driver"
	UserRolePassenger          UserRole = "passenger"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Car struct {
	PlateID   string       `json:"plate_id"`
	Pax       int32        `json:"pax"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Schedule struct {
	ID             int32           `json:"id"`
	DepartureDate  time.Time       `json:"departure_date"`
	DepartureTime  time.Time       `json:"departure_time"`
	Pickup         json.RawMessage `json:"pickup"`
	DropOff        json.RawMessage `json:"drop_off"`
	PickupCountry  Country         `json:"pickup_country"`
	DropOffCountry Country         `json:"drop_off_country"`
	// When carpool confirmed
	DriverID sql.NullInt32 `json:"driver_id"`
	// When carpool confirmed
	PlateID   sql.NullString `json:"plate_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

type SchedulePassenger struct {
	ID          int32         `json:"id"`
	ScheduleID  sql.NullInt32 `json:"schedule_id"`
	PassengerID int32         `json:"passenger_id"`
	Category    Category      `json:"category"`
	Seat        sql.NullInt32 `json:"seat"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   sql.NullTime  `json:"updated_at"`
}

type User struct {
	ID            int32          `json:"id"`
	Username      string         `json:"username"`
	Password      string         `json:"password"`
	Firstname     sql.NullString `json:"firstname"`
	Lastname      sql.NullString `json:"lastname"`
	ContactNumber string         `json:"contact_number"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	RoleID        UserRole       `json:"role_id"`
}
