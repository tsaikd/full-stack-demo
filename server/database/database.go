package database

import (
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/tsaikd/KDGoLib/errutil"
)

// errors
var (
	ErrDatabaseOpenFailed1 = errutil.NewFactory("open database connection failed: %q")
	ErrDatabasePingFailed1 = errutil.NewFactory("ping database connection failed: %q")
)

var db *sqlx.DB

// Database return connection instance
func Database() *sqlx.DB {
	return db
}

// Config init database connection
// remember defer Close()
func Config(connurl string, ssldir string) (err error) {
	if err = Close(); err != nil {
		return
	}

	u, err := url.Parse(connurl)
	if err != nil {
		return
	}

	query := u.Query()
	query.Set("sslmode", "require")
	u.RawQuery = query.Encode()

	driverName := u.Scheme
	u.Scheme = ""
	dataSourceName := strings.TrimPrefix(u.String(), "//")
	if driverName == "postgres" {
		dataSourceName = driverName + "://" + dataSourceName
	}
	if db, err = sqlx.Open(driverName, dataSourceName); err != nil {
		return ErrDatabaseOpenFailed1.New(err, connurl)
	}

	if err = db.Ping(); err != nil {
		return ErrDatabasePingFailed1.New(err, connurl)
	}

	db.SetMaxOpenConns(20)

	return
}

// Close release database connection
func Close() (err error) {
	if db != nil {
		err = db.Close()
		db = nil
	}
	return
}
