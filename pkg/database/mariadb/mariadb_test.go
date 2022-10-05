package mariadb

import (
	"log"
	"testing"

	"gorm.io/gorm"
)

func TestDBConnection(t *testing.T) {
	log.Println("TestDBConnection")

	db := NewDatabaseConnection("authUser", "authUser1234", "127.0.0.1", 3306, "AuthenticationDB")

	err := db.Connect()
	if err != nil {
		t.Error("Error connectiong to database: ", err)
		return
	}

	log.Println("Database connected!")

	err = db.Close()
	if err != nil {
		t.Error("Error closing database connection: ", err)
		return
	}

	log.Println("Database connection closed.")
	log.Println("TestDBConnection OK.")
}

type TesteStruct struct {
	gorm.Model
	Login    string `gorm:"unique" json:"Login"`
	Password string `json:"-"`
}

func TestMigration(t *testing.T) {
	log.Println("TestMigration")

	db := NewDatabaseConnection("authUser", "authUser1234", "127.0.0.1", 3306, "AuthenticationDB")

	err := db.Connect()
	if err != nil {
		t.Error("Error connectiong to database: ", err)
		return
	}

	log.Println("Database connected!")

	err = db.MigrateStruct(TesteStruct{})
	if err != nil {
		t.Error("Error migrating struct to DB.")
		return
	}

	log.Println("Migrationg sucessfull.")

	err = db.Close()
	if err != nil {
		t.Error("Error closing database connection: ", err)
		return
	}

	log.Println("Database connection closed.")

	log.Println("TestMigration OK.")
}
