package db

import (
	"database/sql"
	"fmt"
	"github.com/vin-oys/api-carpool/db/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("? unable to load environment configurations", err)
	}

	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	conn, err := sql.Open(config.DBDriver, dataSourceName)
	if err != nil {
		log.Fatal("? unable to connect to db", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
