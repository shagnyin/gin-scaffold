package logger

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/shagnyin/gin-scaffold/internal/config"
	"github.com/shagnyin/gin-scaffold/pkg/utils"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var cfg config.Log

func InitLog(config config.Log) {
	cfg = config

	hook := NewLfsHook(time.Duration(cfg.RotationTime)*time.Hour, cfg.RemainRotationCount)
	logrus.SetReportCaller(true)
	logrus.AddHook(hook)
}

func WithContext(ctx *gin.Context) *logrus.Entry {
	contextFromContext := trace.SpanContextFromContext(ctx.Request.Context())
	return logrus.WithFields(logrus.Fields{
		"span-id":  contextFromContext.SpanID(),
		"trace-id": contextFromContext.TraceID(),
		"ip":       utils.GetRealIP(ctx.Request),
		//"agent":    ctx.Request.Header.Get("User-Agent"),
	})
}

func NewLfsHook(rotationTime time.Duration, maxRemainNum uint) logrus.Hook {
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum, "debug"),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum, "info"),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum, "warn"),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum, "error"),
	}, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			f := frame.File
			//处理文件名
			dir, _ := os.Getwd()
			return "", strings.Split(f, dir)[1] + ":" + strconv.Itoa(frame.Line)
		},
	})
	return lfsHook
}
func initRotateLogs(rotationTime time.Duration, maxRemainNum uint, level string) *rotatelogs.RotateLogs {
	writer, err := rotatelogs.New(
		cfg.StorageLocation+level+"_"+"%Y%m%d"+".log",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxRemainNum),
	)
	if err != nil {
		panic(err.Error())
	} else {
		return writer
	}
}
