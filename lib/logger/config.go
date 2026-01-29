package logger

type Config struct {
	Filename    string
	MaxSizeMB   int // 每份日志大小
	MaxBackups  int // 日志保留多少份
	MaxDays     int // 保留多少天的日志
	Compress    bool
	Level       string
	LogEncoding string
}
