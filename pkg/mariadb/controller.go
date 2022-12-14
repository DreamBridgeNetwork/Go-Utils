package mariadb

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/DreamBridgeNetwork/Go-Utils/pkg/jsonfile"
	_ "github.com/go-sql-driver/mysql"
)

// CloseDBConnection - Closes and finalize a DB connection.
func (database *DataBaseConfig) CloseDBConnection() error {
	if database.DB == nil {
		log.Println("mariadb.CloseDBConnection - Null pointer to connection.")
		return errors.New("null pointer to connection")
	}
	err := database.DB.Close()

	if err != nil {
		log.Println("mariadb.CloseDBConnection - Error closing connection to DB.")
		return err
	}

	database.DB = nil

	return nil
}

// Connect - Connects to a data base.
func (database *DataBaseConfig) Connect() error {
	log.Println("mariadb.Connect")

	connectionURL := database.mountConnectionURL()

	db, err := sql.Open("mysql", connectionURL)

	if err != nil {
		log.Println("mariadb.Connect - Error opening connection to DB.")
		return err
	}

	db.SetMaxIdleConns(database.MaxConnections) // Número máximo de conexões abertas sem uso
	db.SetMaxOpenConns(database.MaxConnections) // Número máximo de conexões abertas simultaneamente

	var tempoDuracao = time.Minute * time.Duration(database.IdleTimeMinutes)
	db.SetConnMaxLifetime(tempoDuracao) // Tempo máximo que uma conexão pode ser reutilizada. 0 = para sempre

	err = testDBConnection(db)

	if err != nil {
		db = nil
		log.Printf("mariadb.Connect - Fail testing DB connection.")
		return err
	}

	database.DB = db

	return nil
}

// testDBConnection - Sends a PING for the DB to verify if we are connecterd.
func testDBConnection(db *sql.DB) error {

	if db == nil {
		log.Printf("mariadb.testDBConnection - Connection point is null.")
		return errors.New("connection point is null")
	}

	err := db.Ping()

	if err != nil {
		log.Println("mariadb.testDBConnection - Error testing DB connection.")
		return err
	}

	return nil
}

func (database *DataBaseConfig) mountConnectionURL() string {
	log.Println("mariadb.mountConnectionURL")

	//formato da URL => user:password@tcp(127.0.0.1:3306)/database

	if database.URL == "localhost" || database.URL == "127.0.0.1" {
		return database.User + ":" + database.Password + "@tcp(:" + database.Port + ")/" + database.DBName
	}

	return database.User + ":" + database.Password + "@tcp(" + database.URL + ":" + database.Port + ")/" + database.DBName

}

// SelectSingleRow - Executes a SELECT operation in the data base, that returns only one row
func (database *DataBaseConfig) SelectSingleRow(query string) *sql.Row {

	row := database.DB.QueryRow(query)

	return row
}

// Select - Executes a SELECT operation in the data base
func (database *DataBaseConfig) Select(query string, args ...any) (*sql.Rows, error) {
	// Prepara a query
	log.Println("Query: " + query)
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		log.Println("mariadb.Select - Error preparing the query: " + query)
		return nil, err
	}

	defer stmt.Close()

	// Executs the query in the data base.
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Printf("mariadb.Select - Error executing query: " + query)
		return nil, err
	}

	if rows == nil {
		log.Println("mariadb.Select - null rows .")
		return nil, errors.New("null rows")
	}

	return rows, nil
}

// ExecutaUpdateInsertDelete - Execute a delete, update or insert in the data base
func (database *DataBaseConfig) UpdateInsertDelete(query string) error {
	// Prepare query
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		log.Println("mariadb.UpdateInsertDelete - Error preparing query: " + query)
		return err
	}
	defer stmt.Close()

	// Executes the query in the data base
	_, err = stmt.Exec()

	if err != nil {
		log.Println("mariadb.UpdateInsertDelete - Error executing query: " + query)
		return err
	}

	return nil
}

// LoadDbConfig - Loads the configuration of a new data base connection
func LoadDbConfig(fileName, filepath string) (*DataBaseConfig, error) {
	log.Println("mariadb.loadDbConfig")

	var database DataBaseConfig

	err := jsonfile.ReadJSONFile2(filepath, fileName, &database)

	if err != nil {
		log.Println("mariadb.loadDbConfig - Error reading configuration file.")
		return nil, err
	}

	confJson, err := json.MarshalIndent(database, "", "    ")

	if err != nil {
		log.Println("mariadb.loadDbConfig - Error prointing Json.")
		return nil, err
	}

	log.Println("Database configuration loaded:\n", string(confJson))

	return &database, nil
}
