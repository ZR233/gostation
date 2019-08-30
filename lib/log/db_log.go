/*
@Time : 2019-06-28 14:26
@Author : zr
*/
package log

import (
	_ "camdig/server/config"
	"camdig/server/model"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger *logrus.Logger
var apiWriter *APIWriter

const FileNamePrefix = "apit"

type APIWriter struct {
	SqlChan  chan *logrus.Entry
	FileChan chan *logrus.Entry
}

func newAPIWriter() *APIWriter {
	a := &APIWriter{}
	a.SqlChan = make(chan *logrus.Entry, 3000)
	return a
}

func init() {
	apiWriter = newAPIWriter()
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)
	Logger.AddHook(apiHook{})

	go func() {
		for {
			apiWriter.saveAPILogDB()
		}
	}()
	go func() {
		for {
			apiWriter.saveAPILogFile()
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			apiWriter.saveAPILogFileToDB()
		}
	}()
}

type apiHook struct {
}

func (a apiHook) Fire(entry *logrus.Entry) error {
	select {
	case apiWriter.SqlChan <- entry:
		return nil
	default:
		apiWriter.FileChan <- entry
	}
	return nil
}
func (a apiHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func APILogFromEntry(entry *logrus.Entry) (row *model.ApiLog) {

	execTime := entry.Data["execTime"].(time.Duration)
	apiName := entry.Data["apiName"].(string)
	optTime := entry.Data["optTime"].(time.Time)
	optUserid := entry.Data["optUserid"].(int)
	src := entry.Data["src"].(string)
	code := entry.Data["code"].(int)
	paramsD := entry.Data["params"].([]byte)
	//params := base64.StdEncoding.EncodeToString(paramsD)
	row = &model.ApiLog{
		OptTime:   optTime,
		ApiName:   apiName,
		Code:      code,
		ExecTime:  int(execTime),
		OptUserid: optUserid,
		Src:       src,
		Params:    string(paramsD),
		Msg:       entry.Message,
	}
	return
}
