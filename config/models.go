package config

type DatabaseConfig struct {
	DBName      string `json:"dbName"`
	DBType      string `json:"dbType"`
	IsTimeSeries  bool   `json:"isTimSeries"`
	IsLockEnabled bool   `json:"isLockEnabled"`
	TimelyConfig timelyConfig `json:"TimelyConfig"`
}

type timelyConfig struct {
	IsTimelyEnabled bool   `json:"isTimelyEnabled"`
	TimelyType rune  `json:"timelyType"`
}

func GetTimelyType() map[rune]string {
 return map[rune]string {
		'h': "HOUR",
		'm': "MINUTE",
		'd': "DAY",
		'w': "WEEK",
		'o': "MONTH",
	}
}

type Config struct {
	Databases []DatabaseConfig `json:"databases"`
}
