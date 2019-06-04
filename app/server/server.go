package server

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"crowns/config"
)

type SimpleServer struct {
	serverConfig *config.ServerConfig
	db           *sql.DB
}

func (s *SimpleServer) Init(serverConfig *config.ServerConfig) {
	s.serverConfig = serverConfig
	connPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		serverConfig.DatabaseConfig.Host,
		serverConfig.DatabaseConfig.Port,
		serverConfig.DatabaseConfig.User,
		serverConfig.DatabaseConfig.Password,
		serverConfig.DatabaseConfig.DbName,
	)
	var dbErr error
	s.db, dbErr = sql.Open("postgres", connPath)
	if dbErr == nil {
		dbErr = s.db.Ping()
	}
	if dbErr != nil {
		panic("could not establish database connection: " + dbErr.Error())
	}
}

func (s *SimpleServer) GetDbForTest() *sql.DB {
	return s.db
}

func (s *SimpleServer) Close() {
	s.db.Close()
}

func (s *SimpleServer) GetConfig() *config.ServerConfig {
	return s.serverConfig
}
