package serverconfig

import (
	"database/sql"
	"fmt"
	"github.com/alif-github/project-test/config"
	"github.com/alif-github/project-test/model"
	"github.com/alif-github/project-test/util"
	_ "github.com/jackc/pgx/stdlib"
	"os"
	"sync"
)

var ServerAttribute serverAttribute

type dbInfo struct {
	instance      *sql.DB
	driver        string
	connectionStr string
	setParams     []string
}

type serverAttribute struct {
	DBConnection *sql.DB
}

var once sync.Once

func SetServerAttribute() {
	var (
		err      error
		dbConfig = config.ApplicationConfiguration.GetPostgresql()
	)

	dbInfoData := dbInfo{
		instance:      nil,
		driver:        dbConfig.Driver,
		connectionStr: dbConfig.Address,
		setParams:     []string{fmt.Sprintf(`search_path = '%s'`, dbConfig.DefaultSchema)},
	}

	once.Do(func() {
		if dbInfoData.setParams != nil && len(dbInfoData.setParams) > 0 {
			for _, param := range dbInfoData.setParams {
				dbInfoData.connectionStr += fmt.Sprintf(` %s`, param)
			}
		}

		dbInfoData.instance, err = sql.Open(dbInfoData.driver, dbInfoData.connectionStr)
		if err != nil {
			logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServer().Version, config.ApplicationConfiguration.GetServer().Application)
			logModel.Code = 500
			logModel.Message = err.Error()
			util.LogError(logModel.LoggerZapFieldObject())
			os.Exit(1)
		}

		ServerAttribute.DBConnection = dbInfoData.instance
	})
}
