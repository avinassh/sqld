package libsqldriver

import (
	"database/sql"
	"database/sql/driver"
	"strings"

	tursodriver "github.com/avinassh/sqld/packages/golang/libsql-client/internal/turso/sqldriver"
	"github.com/mattn/go-sqlite3"
)

type LibsqlDriver struct {
}

func (d *LibsqlDriver) Open(dbPath string) (driver.Conn, error) {
	if strings.HasPrefix(dbPath, "file:") {
		return (&sqlite3.SQLiteDriver{}).Open(dbPath)
	}
	return tursodriver.TursoConnect(dbPath), nil
}

func init() {
	sql.Register("libsql", &LibsqlDriver{})
}
