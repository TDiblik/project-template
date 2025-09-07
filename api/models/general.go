package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Custom general types
type SQLNullString struct {
	sql.NullString
}

func (ns SQLNullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil) // or return `[]byte("null")` for explicit `null`
}
func (ns *SQLNullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.String = *s
		ns.Valid = true
	} else {
		ns.Valid = false
	}
	return nil
}

type SQLNullTime struct {
	sql.NullTime
}

func (nt SQLNullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	}
	return json.Marshal(nil) // or return `[]byte("null")` for explicit `null`
}
func (nt *SQLNullTime) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		nt.Time = *t
		nt.Valid = true
	} else {
		nt.Valid = false
	}
	return nil
}
