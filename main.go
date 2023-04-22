package main

import (
	"fmt"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/router"
	"github.com/alif-github/project-test/serverconfig"
	"github.com/alif-github/project-test/util"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"os"
)

func main() {

	//--- Generate Configuration
	environment := "development"
	args := os.Args
	if len(args) > 1 {
		environment = args[1]
		fmt.Println("Run in environment : ", environment)
	}

	config.GenerateConfiguration(environment)

	//--- Define Logger
	util.ConfigZap(config.ApplicationConfiguration.GetLogFile())

	//--- Set Server Attribute
	serverconfig.SetServerAttribute()

	//--- Auto Create Schema
	autoCreateSchema()

	//--- DB Migration
	dbMigration()

	//--- Info Starting Web
	logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
	logModel.Message = fmt.Sprintf(`Starting Port %s`, config.ApplicationConfiguration.GetServer().Port)
	util.LogInfo(logModel.LoggerZapFieldObject())

	//--- Router
	router.APIController()
}

func autoCreateSchema() {
	createSchema := fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %s;`, config.ApplicationConfiguration.GetPostgresql().DefaultSchema)
	_, errS := serverconfig.ServerAttribute.DBConnection.Exec(createSchema)
	if errS != nil {
		logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
		logModel.Code = 500
		logModel.Message = fmt.Sprintf(`Error auto create schema -> %s`, errS.Error())
		util.LogError(logModel.LoggerZapFieldObject())
		os.Exit(3)
	}
}

func dbMigration() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	if serverconfig.ServerAttribute.DBConnection != nil {
		n, err := migrate.Exec(serverconfig.ServerAttribute.DBConnection, "postgres", migrations, migrate.Up)
		if err != nil {
			logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
			logModel.Code = 500
			logModel.Message = fmt.Sprintf(`Error on migration -> %s`, err.Error())
			util.LogError(logModel.LoggerZapFieldObject())
			os.Exit(3)
		}

		logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
		logModel.Code = 200
		logModel.Message = fmt.Sprintf(`Has Applied %d Migrations`, n)
		util.LogInfo(logModel.LoggerZapFieldObject())
		return
	}

	logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
	logModel.Code = 500
	logModel.Message = fmt.Sprintf(`DB Connection Not Found`)
	util.LogError(logModel.LoggerZapFieldObject())
	os.Exit(3)
}
