package ping

import (
	"database/sql"
	"fmt"
	"time"
)

var (
	// table name
	table = "ping"
)

// item defines the model
type Item struct {
	IP      string    `db:"ip"`
	Visited time.Time `db:"visited"`
}

// Service defines the database connection
type Service struct {
	DB Connection
}

type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// lists 10 most recent
func (s Service) ByRecent() ([]Item, bool, error) {
	var result []Item
	err := s.DB.Select(&result, fmt.Sprintf(`
        SELECT ip, visited
        FROM %v
        ORDER BY visited DESC
        LIMIT 10`,
		table))
	return result, err == sql.ErrNoRows, err
}

func (s Service) Create(ip string, visited time.Time) (sql.Result, error) {
	result, err := s.DB.Exec(fmt.Sprintf(`
        INSERT INTO %v
        ('ip, visited')
        values(?,?)
        `, table),
		ip, visited)
	return result, err
}
