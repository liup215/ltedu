package conf

import (
	"edu/lib/ai"
	"edu/lib/database/orm"
	"edu/lib/logger"
	"edu/lib/net/http/middleware/auth"
	"edu/lib/verification"
	"os"
	"strconv"
)

func init() {
	initConfig()
}

const (
	ENV_LTEDU_ORM_DIALECT                  = "LTEDU_DIALECT"
	ENV_LTEDU_MYSQL_HOST                   = "LTEDU_MYSQL_HOST"
	ENV_LTEDU_MYSQL_PORT                   = "LTEDU_MYSQL_PORT"
	ENV_LTEDU_MYSQL_USERNAME               = "LTEDU_MYSQL_USERNAME"
	ENV_LTEDU_MYSQL_PASSWORD               = "LTEDU_MYSQL_PASSWORD"
	ENV_LTEDU_MYSQL_CHARSET                = "LTEDU_MYSQL_CHARSET"
	ENV_LTEDU_MYSQL_DATABASE               = "LTEDU_MYSQL_DATABASE"
	ENV_LTEDU_SQLITE_DB_PATH               = "LTEDU_SQLITE_DB_PATH"
	ENV_LTEDU_HTTP_PORT                    = "LTEDU_HTTP_PORT"
	ENV_LTEDU_LOGGER_LEVEL                 = "LTEDU_LOGGER_LEVEL"
	ENV_LTEDU_LOGGER_MAX_SIZE_MB           = "LTEDU_LOGGER_MAX_SIZE_MB"
	ENV_LTEDU_LOGGER_MAX_BACKUPS           = "LTEDU_LOGGER_MAX_BACKUPS"
	ENV_LTEDU_LOGGER_MAX_DAYS              = "LTEDU_LOGGER_MAX_DAYS"
	ENV_LTEDU_LOGGER_COMPRESS              = "LTEDU_LOGGER_COMPRESS"
	ENV_LTEDU_LOGGER_ENCODING              = "LTEDU_LOGGER_ENCODING"
	ENV_LTEDU_AI_DIALECT                   = "LTEDU_AI_DIALECT"
	ENV_LTEDU_AI_ALI_BAILIAN_AGENT_KEY     = "LTEDU_AI_ALI_BAILIAN_AGENT_KEY"
	ENV_LTEDU_AI_ALI_BAILIAN_ACCESS_KEY    = "LTEDU_AI_ALI_BAILIAN_ACCESS_KEY"
	ENV_LTEDU_AI_ALI_BAILIAN_ACCESS_SECRET = "LTEDU_AI_ALI_BAILIAN_ACCESS_SECRET"
	ENV_LTEDU_AI_ALI_BAILIAN_APP_ID        = "LTEDU_AI_ALI_BAILIAN_APP_ID"
	ENV_LTEDU_SMTP_HOST                    = "LTEDU_SMTP_HOST"
	ENV_LTEDU_SMTP_PORT                    = "LTEDU_SMTP_PORT"
	ENV_LTEDU_SMTP_USERNAME                = "LTEDU_SMTP_USERNAME"
	ENV_LTEDU_SMTP_PASSWORD                = "LTEDU_SMTP_PASSWORD"
	ENV_LTEDU_SMTP_FROM_EMAIL              = "LTEDU_SMTP_FROM_EMAIL"
)

var Conf *Config

type Config struct {
	Orm             *orm.Config
	Auth            *auth.Config
	Http            *HttpConfig
	Smtp            *verification.SmtpConfig
	Domain          string // 应用域名
	Public          string // 目录
	UniofficeApiKey string
	Logger          *logger.Config
	SecretKey       string
	AiConfig        *ai.Config
}

type HttpConfig struct {
	Port string
}

func initConfig() {

	// Initialize the environment variables if not set
	ORM_DIALECT := os.Getenv(ENV_LTEDU_ORM_DIALECT)
	MYSQL_HOST := os.Getenv(ENV_LTEDU_MYSQL_HOST)
	MYSQL_PORT := os.Getenv(ENV_LTEDU_MYSQL_PORT)
	MYSQL_USERNAME := os.Getenv(ENV_LTEDU_MYSQL_USERNAME)
	MYSQL_PASSWORD := os.Getenv(ENV_LTEDU_MYSQL_PASSWORD)
	MYSQL_CHARSET := os.Getenv(ENV_LTEDU_MYSQL_CHARSET)
	MYSQL_DATABASE := os.Getenv(ENV_LTEDU_MYSQL_DATABASE)
	SQLITE_DB_PATH := os.Getenv(ENV_LTEDU_SQLITE_DB_PATH)
	HTTP_PORT := os.Getenv(ENV_LTEDU_HTTP_PORT)
	LOGGER_LEVEL := os.Getenv(ENV_LTEDU_LOGGER_LEVEL)
	LOGGER_MAX_SIZE_MB := os.Getenv(ENV_LTEDU_LOGGER_MAX_SIZE_MB)
	LOGGER_MAX_BACKUPS := os.Getenv(ENV_LTEDU_LOGGER_MAX_BACKUPS)
	LOGGER_MAX_DAYS := os.Getenv(ENV_LTEDU_LOGGER_MAX_DAYS)
	LOGGER_COMPRESS := os.Getenv(ENV_LTEDU_LOGGER_COMPRESS)
	LOGGER_ENCODING := os.Getenv(ENV_LTEDU_LOGGER_ENCODING)
	AI_DIALECT := os.Getenv(ENV_LTEDU_AI_DIALECT)
	ALI_BAILIAN_AGENT_KEY := os.Getenv(ENV_LTEDU_AI_ALI_BAILIAN_AGENT_KEY)
	ALI_BAILIAN_ACCESS_KEY := os.Getenv(ENV_LTEDU_AI_ALI_BAILIAN_ACCESS_KEY)
	ALI_BAILIAN_ACCESS_SECRET := os.Getenv(ENV_LTEDU_AI_ALI_BAILIAN_ACCESS_SECRET)
	ALI_BAILIAN_APP_ID := os.Getenv(ENV_LTEDU_AI_ALI_BAILIAN_APP_ID)
	SMTP_HOST := os.Getenv(ENV_LTEDU_SMTP_HOST)
	SMTP_PORT_STR := os.Getenv(ENV_LTEDU_SMTP_PORT)
	SMTP_USERNAME := os.Getenv(ENV_LTEDU_SMTP_USERNAME)
	SMTP_PASSWORD := os.Getenv(ENV_LTEDU_SMTP_PASSWORD)
	SMTP_FROM_EMAIL := os.Getenv(ENV_LTEDU_SMTP_FROM_EMAIL)

	// Set default values if environment variables are not set
	if ORM_DIALECT == "" {
		ORM_DIALECT = orm.DIALECT_SQLITE // Default to SQLite if not set
	}

	if ORM_DIALECT == orm.DIALECT_MYSQL {
		if MYSQL_HOST == "" {
			MYSQL_HOST = "localhost"
		}
		if MYSQL_PORT == "" {
			MYSQL_PORT = "3306"
		}
		if MYSQL_USERNAME == "" {
			MYSQL_USERNAME = "root"
		}
		if MYSQL_PASSWORD == "" {
			MYSQL_PASSWORD = "password"
		}
		if MYSQL_CHARSET == "" {
			MYSQL_CHARSET = "utf8mb4"
		}
		if MYSQL_DATABASE == "" {
			MYSQL_DATABASE = "ltedu" // Default MySQL database name
		}
	} else if ORM_DIALECT == orm.DIALECT_SQLITE {
		if SQLITE_DB_PATH == "" {
			SQLITE_DB_PATH = "./data/ltedu.db" // Default SQLite database path
		}
	} else {
		panic("Unsupported database dialect: " + ORM_DIALECT)
	}

	if HTTP_PORT == "" {
		HTTP_PORT = "80" // Default HTTP port
	}

	if LOGGER_LEVEL == "" {
		LOGGER_LEVEL = "debug" // Default logger level
	}
	if LOGGER_MAX_SIZE_MB == "" {
		LOGGER_MAX_SIZE_MB = "10" // Default max size in MB
	}

	if LOGGER_MAX_BACKUPS == "" {
		LOGGER_MAX_BACKUPS = "10" // Default max backups
	}
	if LOGGER_MAX_DAYS == "" {
		LOGGER_MAX_DAYS = "30" // Default max days
	}
	if LOGGER_COMPRESS == "" {
		LOGGER_COMPRESS = "false" // Default compression
	}
	if LOGGER_ENCODING == "" {
		LOGGER_ENCODING = "console" // Default log encoding
	}
	if AI_DIALECT == "" {
		AI_DIALECT = ai.ALI_BAILIAN // Default AI dialect
	}

	if SMTP_HOST == "" {
		SMTP_HOST = "smtp.qq.com"
	}
	if SMTP_PORT_STR == "" {
		SMTP_PORT_STR = "587"
	}
	smtpPort, err := strconv.Atoi(SMTP_PORT_STR)
	if err != nil {
		panic("Invalid SMTP port: " + SMTP_PORT_STR)
	}
	if SMTP_USERNAME == "" {
		SMTP_USERNAME = "your-email@qq.com" // Placeholder
	}
	if SMTP_PASSWORD == "" {
		SMTP_PASSWORD = "your-qq-smtp-password" // Placeholder for QQ SMTP authorization code
	}
	if SMTP_FROM_EMAIL == "" {
		SMTP_FROM_EMAIL = SMTP_USERNAME // Default from email to username
	}

	Conf = &Config{
		Smtp: &verification.SmtpConfig{
			Host:     SMTP_HOST,
			Port:     smtpPort,
			Username: SMTP_USERNAME,
			Password: SMTP_PASSWORD,
			From:     SMTP_FROM_EMAIL,
		},
		Orm: &orm.Config{
			Dialect: ORM_DIALECT,
			Mysql: orm.MysqlConfig{
				Host:     MYSQL_HOST,
				Port:     MYSQL_PORT,
				Username: MYSQL_USERNAME,
				Password: MYSQL_PASSWORD,
				Charset:  MYSQL_CHARSET,
				Database: MYSQL_DATABASE,
			},
			Sqlite:       orm.SqliteConfig{DBPath: SQLITE_DB_PATH},
			SQLLog:       true,
			MaxIdleConns: 2,
			MaxOpenConns: 4,
			Prefix:       "lt_",
		},
		Http: &HttpConfig{
			Port: HTTP_PORT,
		},
		Logger: &logger.Config{
			Level:       LOGGER_LEVEL,
			MaxSizeMB:   10,
			MaxBackups:  10,
			MaxDays:     30,
			Compress:    false,
			LogEncoding: LOGGER_ENCODING,
		},
		AiConfig: &ai.Config{
			Dialect: AI_DIALECT,
			AliBaiLian: ai.AliBaiLianConfig{
				AgentKey:        ALI_BAILIAN_AGENT_KEY,
				AccessKey:       ALI_BAILIAN_ACCESS_KEY,
				AccessSecretKey: ALI_BAILIAN_ACCESS_SECRET,
				AppId:           ALI_BAILIAN_APP_ID,
			},
		},
	}
}

// func production() {
// 	Conf = &Config{
// Orm: &orm.Config{
// 	Dialect: orm.DIALECT_MYSQL,
// 	Mysql: orm.MysqlConfig{
// 		Host:     "rm-2vc0jslmy5y0j6doi.mysql.cn-chengdu.rds.aliyuncs.com",
// 		Port:     "3306",
// 		Username: "ltedu",
// 		Password: "wyogx2pAR6bDTtmtQ4jUmBw26dOQdLWW",
// 		Charset:  "utf8",
// 		Database: "ltedu",
// 	},
// 	Sqlite:       orm.SqliteConfig{DBPath: "data/ltedu.db"},
// 	SQLLog:       true,
// 	MaxIdleConns: 2,
// 	MaxOpenConns: 4,
// 	Prefix:       "lt_",
// },
// 		// Badger: &badger.Config{Dir: "badgerDB"},
// 		Http: &HttpConfig{
// 			Port: "9002",
// 		},
// 		Domain:          "http://ltedu.alevel.icu",
// 		Public:          "public",
// 		UniofficeApiKey: "4438f4e27ac6e6c0ff52f8245f26335affb64ffba67d757fdbaf5924accadb2f",
// 		Logger: &logger.Config{
// 			Level:       "debug",
// 			Filename:    "logs/ltedu.log",
// 			MaxSizeMB:   10,
// 			MaxBackups:  10,
// 			MaxDays:     30,
// 			Compress:    false,
// 			LogEncoding: "console",
// 		},
// 		SecretKey: "cZG2xUj6gy3VAq03lC6m2tonh5phvvr7",
// 		AiConfig: &ai.Config{
// 			Dialect: ai.ALI_BAILIAN,
// 			AliBaiLian: ai.AliBaiLianConfig{
// 				AgentKey:        "f011305e68fa411eb71a801e74ac1c40_p_efm",
// 				AccessKey:       "LTAIRmRLjaIG5GUM",
// 				AccessSecretKey: "yA9WgYzVTtVXtegTaur9wtMVgDr7HQ",
// 				AppId:           "688a8f4d0200473ca3f5f64d7085fcf8",
// 			},
// 		},
// 	}
// }
