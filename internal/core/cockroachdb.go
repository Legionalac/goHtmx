package libs

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)
type CockroachDb struct{
	ConnectionPool *pgxpool.Pool
}

func (db *CockroachDb) CreateCockRoachConnection() {
	logger := GetLogger()
	dbpool, err := pgxpool.New(context.Background(), viper.GetString("SQL_CONNECTION_STRING"))
	if err != nil {
		logger.Fatal().Msg("ERROR WHILE CONNECTING TO DATABASE" + err.Error())
	}
	db.ConnectionPool = dbpool
}

func (db *CockroachDb) CloseConnection(){
	db.ConnectionPool.Close()
}