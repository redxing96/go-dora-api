package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志实例
type logger struct {
	// 配置
	config Config
	// zap实例
	zapLogger *zap.Logger
	// 日志记录等级
	level zap.AtomicLevel
	// lumberjack实例
	writer *lumberjack.Logger
}

// 创建日志实例
func NewLogger(cfg Config) *logger {
	l := &logger{
		config: cfg,
	}
	err := l.checkConfig()
	if err != nil {
		log.Fatalln("log config check error: ", err)
		return nil
	}
	l.build()
	// go l.cronRotate() // 每天凌晨切割日志
	return l
}

// 设置日志记录等级
func (receiver *logger) SetLevel(level string) error {
	err := receiver.checkLevel(level)
	if err != nil {
		return err
	}
	l, err := zapcore.ParseLevel(level)
	if err != nil {
		return err
	}
	receiver.level.SetLevel(l)
	return nil
}

// 获取日志记录等级
func (receiver *logger) Level() string {
	return receiver.level.Level().String()
}

// 调试日志
func (receiver *logger) Debug(msg string, fields ...IField) {
	receiver.zapLogger.Debug(msg, receiver.buildFields(fields...)...)
}

// 信息日志
func (receiver *logger) Info(msg string, fields ...IField) {
	receiver.zapLogger.Info(msg, receiver.buildFields(fields...)...)
}

// 警告日志
func (receiver *logger) Warn(msg string, fields ...IField) {
	receiver.zapLogger.Warn(msg, receiver.buildFields(fields...)...)
}

// 错误日志
func (receiver *logger) Error(msg string, fields ...IField) {
	receiver.zapLogger.Error(msg, receiver.buildFields(fields...)...)
}

// 调试日志
func (receiver *logger) Debugf(template string, args ...interface{}) {
	receiver.zapLogger.Sugar().Debugf(template, args...)
}

// 信息日志
func (receiver *logger) Infof(template string, args ...interface{}) {
	receiver.zapLogger.Sugar().Infof(template, args...)
}

// 警告日志
func (receiver *logger) Warnf(template string, args ...interface{}) {
	receiver.zapLogger.Sugar().Warnf(template, args...)
}

// 错误日志
func (receiver *logger) Errorf(template string, args ...interface{}) {
	receiver.zapLogger.Sugar().Errorf(template, args...)
}

// 日志切片
func (receiver *logger) Rotate() error {
	return receiver.writer.Rotate()
}

// 日志刷盘(在程序结束之前记得进行日志刷盘，避免遗漏)
func (receiver *logger) Sync() error {
	return receiver.zapLogger.Sync()
}

// 构造zap.Field实例数组
func (receiver *logger) buildFields(fields ...IField) []zap.Field {
	fieldList := make([]zap.Field, 0)
	for _, v := range fields {
		fieldList = append(fieldList, zap.Any(v.GetKey(), v.GetValue()))
	}
	return fieldList
}

// 校验配置
func (receiver *logger) checkConfig() error {
	if err := receiver.checkFilename(receiver.config.Filename); err != nil {
		return err
	}
	if err := receiver.checkLevel(receiver.config.Level); err != nil {
		return err
	}
	if receiver.config.IsStack {
		if err := receiver.checkLevel(receiver.config.StackLevel); err != nil {
			return err
		}
	}

	return nil
}

// 检验日志记录等级
func (receiver *logger) checkLevel(level string) error {
	switch level {
	case "debug", "info", "warn", "error":
		return nil
	default:
		return errors.New("invalid level")
	}
}

// 校验文件路径
func (receiver *logger) checkFilename(filename string) error {
	if !IsValidFile(filename) {
		return errors.New("invalid filename")
	}
	return nil
}

// 构建*zap.Logger
func (receiver *logger) build() {
	// 日志编码配置
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        zapcore.OmitKey,
		CallerKey:      "file",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.0000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 日志写入器配置
	receiver.writer = &lumberjack.Logger{
		Filename:   receiver.config.Filename,
		MaxSize:    receiver.config.MaxSize,
		MaxAge:     receiver.config.MaxAge,
		MaxBackups: receiver.config.MaxBackups,
		Compress:   receiver.config.Compress,
		LocalTime:  true,
	}

	// 日志记录等级
	level, _ := zap.ParseAtomicLevel(receiver.config.Level)
	receiver.level = level

	// 日志写入器
	writeSyncer := zapcore.AddSync(receiver.writer)
	if receiver.config.IsOutput {
		writeSyncer = zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(receiver.writer))
	}

	// 创建zap.Logger实例
	receiver.zapLogger = zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(writeSyncer),
			receiver.level,
		),
		zap.AddCaller(),
		zap.AddCallerSkip(1+receiver.config.StackSkip),
	)
	// 堆栈信息配置
	if receiver.config.IsStack && receiver.config.StackLevel != "" {
		stackLevel, _ := zapcore.ParseLevel(receiver.config.StackLevel)
		receiver.zapLogger = receiver.zapLogger.WithOptions(zap.AddStacktrace(stackLevel))
	}
	// 模块信息配置
	if receiver.config.ModuleKey != "" && receiver.config.ModuleValue != "" {
		receiver.zapLogger = receiver.zapLogger.With(zap.String(receiver.config.ModuleKey, receiver.config.ModuleValue))
	}
}

// 定时日志轮转(每天0点执行)
func (receiver *logger) cronRotate() {
	c := cron.New(cron.WithSeconds(), cron.WithLogger(cron.DefaultLogger))
	_, err := c.AddFunc("0 0 0 ? * ?", func() {
		receiver.Sync()
		receiver.Rotate()
	})
	if err != nil {
		receiver.zapLogger.Error("log cron init failed: " + err.Error())
		panic(err)
	}
	defer c.Stop()
	c.Run()
}

// 默认日志实例
type DefaultLogger struct {
	level string
}

// 设置日志记录等级
func (d *DefaultLogger) SetLevel(level string) error {
	d.level = level
	return nil
}

// 获取日志记录等级
func (d DefaultLogger) Level() string {
	return d.level
}

// 调试日志
func (d DefaultLogger) Debug(msg string, fields ...IField) {
	d.log("DEBUG", msg, fields...)
}

// 信息日志
func (d DefaultLogger) Info(msg string, fields ...IField) {
	d.log("INFO", msg, fields...)
}

// 警告日志
func (d DefaultLogger) Warn(msg string, fields ...IField) {
	d.log("WARN", msg, fields...)
}

// 错误日志
func (d DefaultLogger) Error(msg string, fields ...IField) {
	d.log("ERROR", msg, fields...)
}

// 调试日志
func (d DefaultLogger) Debugf(template string, args ...interface{}) {
	d.logf("DEBUG", template, args...)
}

// 信息日志
func (d DefaultLogger) Infof(template string, args ...interface{}) {
	d.logf("INFO", template, args...)
}

// 警告日志
func (d DefaultLogger) Warnf(template string, args ...interface{}) {
	d.logf("WARN", template, args...)
}

// 错误日志
func (d DefaultLogger) Errorf(template string, args ...interface{}) {
	d.logf("ERROR", template, args...)
}

// 日志记录
func (d DefaultLogger) log(flag, msg string, fields ...IField) {
	data := make([]string, 0)
	data = append(data, "msg:"+msg)
	for _, v := range fields {
		data = append(data, fmt.Sprintf("%s:%v", v.GetKey(), v.GetValue()))
	}
	fmt.Println("【"+flag+"】", strings.Join(data, " "))
}

// 日志记录
func (d DefaultLogger) logf(flag, template string, args ...interface{}) {
	fmt.Println("【"+flag+"】", "msg:"+fmt.Sprintf(template, args...))
}

// 日志轮转
func (d DefaultLogger) Rotate() error {
	return nil
}

// 日志刷盘
func (d DefaultLogger) Sync() error {
	return nil
}
