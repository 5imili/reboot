package pkgs

import (
	"github.com/spf13/viper"
	"github.com/5imili/reboot/pkg/dao"
	"sync"
	"github.com/5imili/reboot/pkg/dao/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const(
	DaoConnection = "dao.connection"

)
var (
	mysqlDao dao.Storage
	mysqlOnce sync.Once
)

func init(){
	initDaoDefault()
}


func initDaoDefault(){
	viper.SetDefault(DaoConnection,"root:root@/reboot?charset=utf8&parseTime=true")
}

func GetDao() dao.Storage{
	mysqlOnce.Do(func(){
		mysqlDao = mysql.New(&mysql.Options{
			DbConnStr: viper.GetString(DaoConnection),
		})
		if mysqlDao == nil{
			panic("connect mysql failed")
		}
	})
	return mysqlDao
}