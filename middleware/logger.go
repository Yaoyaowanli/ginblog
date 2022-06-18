package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

//Logger 日志中间件
func Logger () gin.HandlerFunc {
	//配置日志中间件参数
	logFilePath := "log/ginblog"
	linkName := "latest_log.log"
	src,err := os.OpenFile(logFilePath,os.O_RDWR | os.O_CREATE,0755)
	if err != nil {
		fmt.Println("err:",err)
	}
	logger := logrus.New()
	//这里out是输出日志的方向
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//日志分割
	logWriter,err := retalog.New(
		logFilePath+"%Y%m%d.log",retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
		)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	hook := lfshook.NewHook(writeMap,&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(hook)


	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime :=fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		//获取主机名
		hostName,err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		//当前请求的HTTP响应状态代码。
		statusCode := c.Writer.Status()
		//获取请求端的ip地址
		clientIp := c.ClientIP()
		//客户端浏览器信息
		userAgent := c.Request.UserAgent()
		//请求正文的长度
		dataSize := c.Writer.Size()
		if dataSize <0 {
			dataSize = 0
		}
		//请求方法
		method := c.Request.Method
		//字符串格式的请求url
		path := c.Request.RequestURI


		entry := logger.WithFields(logrus.Fields{
			"HostName":hostName,
			"StatusCode":statusCode,
			"SpendTime":spendTime,
			"ClientIp":clientIp,
			"Method":method,
			"Path":path,
			"DataSize":dataSize,
			"UserAgent":userAgent,
		})
		//c.errors : 上下文的所有处理程序/中间件的错误列表,大于0说明系统有错误
		if len(c.Errors)>0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		}else if statusCode >= 400 {
			entry.Warn()   //警告
		}else {
			entry.Info()  //成功
		}
	}
}