package mariadb

import (
	"errors"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabaseConnection - Return a new database ready to connect.
func NewDatabaseConnection(user, password, url string, port int, schema string) *MariaDB {
	config := mysql.Config{
		DSN:                       mountDSNString(user, password, url, port, schema), // data source name
		DefaultStringSize:         256,                                               // default size for string fields
		DisableDatetimePrecision:  false,                                             // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                              // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                              // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                             // auto configure based on currently MySQL version
	}

	newDB := MariaDB{
		Config:     &config,
		Connection: nil,
	}

	return &newDB
}

// mountDSNString - Mount the DSN string with the configuration setted for mariaDB database.
func mountDSNString(user, password, url string, port int, schema string) string {
	DSNString := user + ":" + password +
		"@tcp(" + url + ":" + strconv.Itoa(port) + ")/" +
		schema + "?charset=utf8&parseTime=True&loc=Local"
	return DSNString
}

// Connect - Connect to a MariaDB database.
func (db *MariaDB) Connect() error {

	if db.Connection == nil {
		connection, err := gorm.Open(mysql.New(*db.Config), &gorm.Config{})

		if err != nil {
			log.Println("mariadb.Connect - Error openning db connection.")
			return err
		}

		db.Connection = connection
	}

	err := db.TesteConnection()
	if err != nil {
		log.Println("mariadb.Connect - Error testing db connection")
		return err
	}

	return nil
}

// Close - Close the dabatase connection.
func (db *MariaDB) Close() error {
	if db.Connection != nil {
		sqldb, err := db.Connection.DB()
		if err != nil {
			log.Println("database.Close - Error getting DB connection.")
			return err
		}
		err = sqldb.Close()
		if err != nil {
			log.Println("database.Close - Error closing DB.")
			return err
		}
	}

	db.Connection = nil
	return nil
}

// MigrateStruct - Migrate structs to database tables.
func (db *MariaDB) MigrateStruct(structToMigrate interface{}) error {
	if db.Connection == nil {
		log.Println("database.MigrateStruct - Database not connected.")
		return errors.New("database.MigrateStruct - Database not connected")
	}

	err := db.Connection.AutoMigrate(structToMigrate)
	if err != nil {
		log.Println("database.InitialMigration - Error migrating struct to DB.")
		return err
	}
	return nil
}

func (db *MariaDB) TesteConnection() error {
	if db.Connection == nil {
		errMsg := "database not connected"
		log.Println("mariadb.TesteConnection - " + errMsg + ".")

		return errors.New(errMsg)
	}

	sqldb, err := db.Connection.DB()
	if err != nil {
		db.Connection = nil
		log.Println("mariadb.TesteConnection - Error opening DB connection.")
		return err
	}

	err = sqldb.Ping()
	if err != nil {
		db.Connection = nil
		log.Println("database.TesteConnection - Error testing db connection.")
		return err
	}

	return nil
}
