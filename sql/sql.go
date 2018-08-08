package sql

import (
	"io/ioutil"
)

var Sql = map[string]map[string]string{
	"user": map[string]string{
		"getUser": stringifySQL("sql/user/getUser.sql"),
	},
}

func stringifySQL(pathToSql string) string {
	sql, err := ioutil.ReadFile(pathToSql)

	if err == nil {
		return string(sql)
	}

	return "err"
}
