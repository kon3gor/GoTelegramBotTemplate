package config

func Token() string {
	return config.telegram.token
}

func WriteLogs() bool {
	return config.logs.write
}

func LogUnknown() bool {
	return config.bot.log_unknown
}

func LogsPath() string {
	return config.logs.path
}
