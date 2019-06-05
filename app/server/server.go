package server

import (
	"fmt"

	"crowns/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SimpleServer struct {
	serverConfig *config.ServerConfig
	db           *sqlx.DB
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
	s.db, dbErr = sqlx.Open("postgres", connPath)
	if dbErr == nil {
		dbErr = s.db.Ping()
	}
	if dbErr != nil {
		panic("could not establish database connection: " + dbErr.Error())
	}
}

func (s *SimpleServer) GetDbForTest() *sqlx.DB {
	return s.db
}

func (s *SimpleServer) Close() {
	s.db.Close()
}

func (s *SimpleServer) GetConfig() *config.ServerConfig {
	return s.serverConfig
}
