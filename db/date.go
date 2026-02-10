package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Date stores a calendar date (YYYY-MM-DD) and scans/values cleanly with SQLite.
// It intentionally ignores time-of-day.
//
// Stored format: 2006-01-02
type Date time.Time

const dateLayout = "2006-01-02"

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	if t.IsZero() {
		return nil, nil
	}
	return t.Format(dateLayout), nil
}

func (d *Date) Scan(src any) error {
	if src == nil {
		*d = Date(time.Time{})
		return nil
	}

	switch v := src.(type) {
	case time.Time:
		*d = Date(v)
		return nil
	case string:
		return d.scanString(v)
	case []byte:
		return d.scanString(string(v))
	default:
		return fmt.Errorf("db.Date: unsupported Scan type %T", src)
	}
}

func (d *Date) scanString(s string) error {
	if s == "" {
		*d = Date(time.Time{})
		return nil
	}

	// Most robust first: date-only
	if t, err := time.Parse(dateLayout, s); err == nil {
		*d = Date(t)
		return nil
	}

	// Fallback: RFC3339-ish timestamps (if any existing rows were stored that way)
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		*d = Date(t)
		return nil
	}
	if t, err := time.Parse(time.RFC3339Nano, s); err == nil {
		*d = Date(t)
		return nil
	}

	return fmt.Errorf("db.Date: unable to parse %q", s)
}
