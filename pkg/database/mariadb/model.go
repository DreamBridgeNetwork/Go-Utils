package mariadb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MariaDB - Struct containing the necessary information to stabish a coonection to a mariaDB database.
type MariaDB struct {
	Config     *mysql.Config
	Connection *gorm.DB
}
