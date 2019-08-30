/*
@Time : 2019-08-19 17:32
@Author : zr
*/
package log

import (
	"bufio"
	_ "camdig/server/config"
	dao "camdig/server/dao/gorm"
	"camdig/server/model"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync/atomic"
	"time"
)

var logIter uint32 = 0
var logDir = path.Join("temp", "log")

//日志保存至数据库
func (a *APIWriter) saveAPILogDB() {
	ticker := time.NewTicker(5 * time.Second)

	var rows []model.ApiLog
	var entries []*logrus.Entry

	for i := 0; i < 1000; i++ {

		select {
		case entry := <-a.SqlChan:
			logRow := model.ApiLog{}

			logRow = *APILogFromEntry(entry)

			rows = append(rows, logRow)
			entries = append(entries, entry)
		case <-ticker.C:
			i = 1000

		}
	}
	if len(rows) > 0 {
		logDao := dao.GetApiLogDAO()
		if err := logDao.Save(rows); err != nil {
			logrus.Warn(err)
			for _, entry := range entries {
				a.FileChan <- entry
			}
		}
	}
}

//日志缓存溢出时，保存至文件
func (a *APIWriter) saveAPILogFile() {
	ticker := time.NewTicker(2 * time.Second)
	logName := path.Join(logDir, "writing")

	_, err := os.Stat(logDir)
	if err != nil {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			logrus.Panic(err)
		}
	}

	//以读写方式打开文件，如果不存在，则创建
	file, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logrus.Panic(err)
	}
	defer func() {
	_:
		file.Close()
	}()

	for i := 0; i < 1000; i++ {
		select {
		case entry := <-a.FileChan:
			logRow := APILogFromEntry(entry)
			rowData, err := json.Marshal(logRow)
			str := string(rowData) + "\n"
			if err != nil {
				logrus.Warn(err)
				continue
			}
			_, err = file.Write([]byte(str))
			if err != nil {
				logrus.Warn(err)
				continue
			}

		case <-ticker.C:
			i = 1000

		}
	}

	err = file.Sync()
	if err != nil {
		logrus.Warn(err)
		return
	}

	stat, err := file.Stat()
	if err != nil {
		logrus.Warn(err)
		return
	}
	if stat.Size() > 0 {
		iter := atomic.AddUint32(&logIter, 1)
		newLogName := path.Join(logDir, fmt.Sprintf("apit%dn%d.log", time.Now().Unix(), iter))
		//创建新文件
		newFile, err := os.Create(newLogName)
		if err != nil {
			logrus.Warn(err)
			return
		}
		defer func() {
		_:
			newFile.Close()
		}()

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			logrus.Warn(err)
			return
		}
		//文件复制
		_, err = io.Copy(newFile, file)
		if err != nil {
			logrus.Warn(err)
			return
		}
	}

	err = file.Truncate(0)
	if err != nil {
		logrus.Warn(err)
		return
	}
}

//日志文件转存至数据库
func (a *APIWriter) saveAPILogFileToDB() {

	rd, err := ioutil.ReadDir(logDir)
	if err != nil {
		logrus.Warn("SaveLog read dir fail:", err)
		return
	}

	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			fileName := fi.Name()
			indexNum := strings.Index(fileName, FileNamePrefix)
			if indexNum != 0 {
				continue
			}

			logFilePath := path.Join(logDir, fileName)
			f, err := os.Open(logFilePath)
			if err != nil {
				logrus.Warn(err)
				continue
			}

			s := bufio.NewScanner(f)

			var rows []model.ApiLog
			for s.Scan() {

				text := s.Text()
				logrus.Info(text)
				logRow := model.ApiLog{}
				if err := json.Unmarshal([]byte(text), &logRow); err != nil {
					logrus.Warn(err)
				_:
					f.Close()
					continue
				}
				rows = append(rows, logRow)
			}
			err = s.Err()
			if err != nil {
				logrus.Warn(err)
			_:
				f.Close()
				continue
			}
		_:
			f.Close()
			if err = os.Remove(logFilePath); err != nil {
				logrus.Warn(err)
			}

			logDao := dao.GetApiLogDAO()
			if err = logDao.Save(rows); err != nil {
				logrus.Warn(err)
			}

		}
	}
}

func Clean() {
	timeBefore := time.Now().Add(-time.Hour * 24 * 5)

	logDao := dao.GetApiLogDAO()
	if err := logDao.DeleteByOptTimeBefore(timeBefore); err != nil {
		logrus.Warn(err)
	}
}
