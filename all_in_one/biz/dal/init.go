package dal

import (
	"hertz-examples/all_in_one/biz/dal/mysql"
	"hertz-examples/all_in_one/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
