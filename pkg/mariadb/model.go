package mariadb

import "database/sql"

// DataBaseConfig - Struct containing the necessary information to stabish a coonection to a server.
type DataBaseConfig struct {
	User            string `json:"user"`
	Password        string `json:"password"`
	URL             string `json:"url"`
	Port            string `json:"port"`
	DBName          string `json:"dbName"`
	MaxConnections  int    `json:"maxConnections"`
	IdleTimeMinutes int    `json:"idleTimeMinutes"`
	DB              *sql.DB
}
