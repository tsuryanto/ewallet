package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *SqlConn

// SqlConn is a wrapper around the gorm.DB object for MySQL
type SqlConn struct {
	db *gorm.DB
}

// DB returns the current active DB connection
func (s *SqlConn) DB() *gorm.DB {
	return s.db
}

// New returns the existing SqlConn if already initialized
func New() *SqlConn {
	return conn
}

// NewDB creates and returns a new MySQL database connection if it's not already established
func NewDB(driver, host, port, user, password, name string) (*SqlConn, error) {
	if conn == nil {
		// MySQL DSN (Data Source Name)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, name)

		// Open the MySQL connection
		gormConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		// Initialize the global conn variable with the established connection
		conn = &SqlConn{
			db: gormConnection,
		}

		// err = conn.db.AutoMigrate(&model.User{}, &model.UserBalance{}, &model.TopUpTransaction{})
		// if err != nil {
		// 	return nil, err
		// }
	}

	// Return the existing connection
	return conn, nil
}
