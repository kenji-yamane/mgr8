package drivers

import (
	"fmt"

	"github.com/migratemgr8/mgr8/domain"
	"github.com/migratemgr8/mgr8/drivers/mysql"
	"github.com/migratemgr8/mgr8/drivers/postgres"
)

type DriverName string

const (
	MySql    DriverName = "mysql"
	Postgres DriverName = "postgres"
)

func GetDriver(driverName string) (domain.Driver, error) {
	driver := DriverName(driverName)
	switch driver {
	case Postgres:
		return postgres.NewPostgresDriver(), nil
	case MySql:
		return mysql.NewMySqlDriver(), nil
	}
	return nil, fmt.Errorf("inexistent driver %s", driverName)
}
