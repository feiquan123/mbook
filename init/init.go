package init

import (
	_ "mbook/models"
)

func init() {
	dbInit(dbOpt{
		Model:  dbModelMaster,
		DBName: "mbook",
	},dbOpt{
		Model:  dbModelSlave,
		DBName: "mbook",
	})
	sysInit()
}
