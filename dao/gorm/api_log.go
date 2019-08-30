/*
@Time : 2019-07-08 10:57
@Author : zr
*/
package gorm

import (
	"camdig/server/model"
	"camdig/server/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ApiLogDAO struct {
	*BaseSqlDAO
}

func GetApiLogDAO() ApiLogDAO {
	m := ApiLogDAO{
		BaseSqlDAO: newBase(),
	}
	return m
}

func (a *ApiLogDAO) Save(apiLog []model.ApiLog) error {
	sql := "INSERT INTO " + model.ApiLog{}.TableName() + " (opt_time, api_name, code, exec_time, opt_userid, src,params,msg,trace)VALUES "

	for _, v := range apiLog {
		//params, _ := base64.StdEncoding.DecodeString(v.Params)

		s := fmt.Sprintf("('%s', '%s', %s, %s, %s, '%s', '%s', '%s', '%s'),", v.OptTime.Format(utils.TimeFormatter()),
			v.ApiName, strconv.Itoa(v.Code), strconv.Itoa(v.ExecTime), strconv.Itoa(v.OptUserid),
			v.Src, v.Params, v.Msg, v.Trace)
		sql += s
	}
	sql = strings.TrimRight(sql, ",")
	return a.db.Exec(sql).Error
}

func (a *ApiLogDAO) DeleteByOptTimeBefore(time time.Time) error {
	return a.db.Delete(model.ApiLog{}, "opt_time < ?", time).Error
}
