package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

func CreateConnection() *sql.DB {
	configFile, err := ioutil.ReadFile("/Users/divyabansal/GoLangProject/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.DBName)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}
	return dbConn
}
