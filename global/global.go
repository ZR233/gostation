/*
@Time : 2019-07-17 10:30
@Author : zr
*/
package global

import "github.com/spf13/viper"

func Debug() bool {
	v := viper.GetViper()
	return v.GetBool("debug")
}

func ServerListen() string {
	v := viper.GetViper()
	return v.GetString("server.listen")
}

func ServerHost() string {
	v := viper.GetViper()
	return v.GetString("server.host")
}

func ServerPort() string {
	v := viper.GetViper()
	return v.GetString("server.port")
}

func WXLoginAppId() string {
	return ""
}

func WXLoginSecret() string {
	return ""
}
func StaticFileWebhdfsHost() string {
	v := viper.GetViper()
	return v.GetString("static-file.webhdfs.host")
}
func StaticFileWebhdfsPort() string {
	v := viper.GetViper()
	return v.GetString("static-file.webhdfs.port")
}

func StaticFilePathCamPic() string {
	v := viper.GetViper()
	return v.GetString("static-file.path.cam-pic")
}
func StaticFilePathApk() string {
	v := viper.GetViper()
	return v.GetString("static-file.path.apk")
}

func CamDetectServerHost() string {
	v := viper.GetViper()
	return v.GetString("cam-detect-server.host")
}

func CamDetectServerPort() string {
	v := viper.GetViper()
	return v.GetString("cam-detect-server.port")
}
func CamDetectServerSecure() bool {
	v := viper.GetViper()
	return v.GetBool("cam-detect-server.secure")
}

func LogPath() string {
	logConf := viper.GetStringMapString("log")
	logPath := logConf["path"]
	return logPath
}

var Test = false
