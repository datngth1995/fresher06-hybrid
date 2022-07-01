package db

import (
	"encoding/json"
	"fmt"
	"os"
	"programming_fresher/common"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sQLConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Dbname    string `json:"dbname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc       string `json:"loc"`
}

func MigrationDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: common.MigrationPath,
	}

	n, err := migrate.Exec(sqlDB, common.DriverName, migrations, migrate.Down)

	if err != nil {
		return err
	}

	n, err = migrate.Exec(sqlDB, common.DriverName, migrations, migrate.Up)

	if err != nil {
		return err
	}

	fmt.Println("Applied migraions! ", n)

	return nil
}

func loadDBConfigs(filepath string) *sQLConfig {
	configFile, err := os.Open(filepath)

	defer configFile.Close()

	if err != nil {
		return nil
	}

	jsonParser := json.NewDecoder(configFile)
	config := sQLConfig{}
	jsonParser.Decode(&config)

	return &config
}

func ConnectToDB() (*gorm.DB, error) {
	config := loadDBConfigs(common.PathSQLConfig)

	mySQLInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		config.Username, config.Password, config.Host, config.Port, config.Dbname, config.Charset, config.ParseTime, config.Loc)

	db, err := gorm.Open(mysql.Open(mySQLInfo), &gorm.Config{})

	fmt.Println(db, err)

	if err != nil {
		return nil, err
	}

	//Turn on debug log
	db = db.Debug()

	//Migration & Seed into database
	if err := MigrationDB(db); err != nil {
		fmt.Println(fmt.Sprint(err))
		return nil, err
	}

	return db, nil
}
