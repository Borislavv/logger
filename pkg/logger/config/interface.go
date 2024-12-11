package loggerconfig

type Configurator interface {
	GetBufferCapacity() int
	GetLoggerLevel() string
	GetLoggerOutput() string
	GetLoggerFormatter() string
	GetLoggerLogsDir() string
	GetLoggerContextExtraFields() []string
}
