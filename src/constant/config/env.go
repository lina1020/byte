package config

import (
	"fmt"

	//"github.com/caarlos0/env/v6"
	//"github.com/joho/godotenv"
	//log "github.com/sirupsen/logrus"
)

var EnvCfg envConfig

type envConfig struct {
	ConsulAddr                string  `env:"CONSUL_ADDR" envDefault:"localhost:8500"`
	LoggerLevel               string  `env:"LOGGER_LEVEL" envDefault:"INFO"`
	LoggerWithTraceState      string  `env:"LOGGER_OUT_TRACING" envDefault:"disable"`
	TiedLogging               string  `env:"TIED" envDefault:"NONE"`
	MySQLHost            string  `env:"MYSQL_HOST"`
	MySQLPort            string  `env:"MYSQL_PORT"`
	MySQLUser            string  `env:"MYSQL_USER"`
	MySQLPassword        string  `env:"MYSQL_PASSWORD"`
	MySQLDataBase        string  `env:"MYSQL_DATABASE"`
	StorageType               string  `env:"STORAGE_TYPE" envDefault:"fs"`
	FileSystemStartPath       string  `env:"FS_PATH" envDefault:"/tmp"`
	FileSystemBaseUrl         string  `env:"FS_BASEURL" envDefault:"http://localhost/"`
	RedisAddr                 string  `env:"REDIS_ADDR"`
	RedisPassword             string  `env:"REDIS_PASSWORD" envDefault:""`
	RedisDB                   int     `env:"REDIS_DB" envDefault:"0"`
	TracingEndPoint           string  `env:"TRACING_ENDPOINT"`
	PyroscopeState            string  `env:"PYROSCOPE_STATE" envDefault:"false"`
	PyroscopeAddr             string  `env:"PYROSCOPE_ADDR"`
	RedisPrefix               string  `env:"REDIS_PREFIX" envDefault:"GUGUTIK"`
	PostgreSQLSchema          string  `env:"POSTGRESQL_SCHEMA" envDefault:""`
	RedisMaster               string  `env:"REDIS_MASTER"`
	ConsulAnonymityPrefix     string  `env:"CONSUL_ANONYMITY_NAME" envDefault:""`
	RabbitMQUsername          string  `env:"RABBITMQ_USERNAME" envDefault:"guest"`
	RabbitMQPassword          string  `env:"RABBITMQ_PASSWORD" envDefault:"guest"`
	RabbitMQAddr              string  `env:"RABBITMQ_ADDRESS" envDefault:"localhost"`
	RabbitMQPort              string  `env:"RABBITMQ_PORT" envDefault:"5672"`
	RabbitMQVhostPrefix       string  `env:"RABBITMQ_VHOST_PREFIX" envDefault:""`
	ChatGPTAPIKEYS            string  `env:"CHATGPT_API_KEYS"`
	PodIpAddr                 string  `env:"POD_IP" envDefault:"localhost"`
	GorseAddr                 string  `env:"GORSE_ADDR"`
	GorseApiKey               string  `env:"GORSE_APIKEY"`
	MagicUserId               uint32  `env:"MAGIC_USER_ID" envDefault:"1"`
	ChatGptProxy              string  `env:"CHATGPT_PROXY"`
	MySQLPrefix          string  `env:"MYSQL_PREFIX" envDefault:""`
	MySQLReplicaState    string  `env:"MYSQL_REPLICA"`
	MySQLReplicaAddress  string  `env:"MYSQL_REPLICA_ADDR"`
	MySQLReplicaUsername string  `env:"MYSQL_REPLICA_USER"`
	MySQLReplicaPassword string  `env:"MYSQL_REPLICA_PASSWORD"`
	OtelState                 string  `env:"TRACING_STATE" envDefault:"enable"`
	OtelSampler               float64 `env:"TRACING_SAMPLER" envDefault:"0.01"`
	AnonymityUser             string  `env:"ANONYMITY_USER" envDefault:"114514"`
	ElasticsearchUrl          string  `env:"ES_ADDR"` 
}

func init() {
	EnvCfg = envConfig{
		MySQLHost:     "localhost",
		MySQLPort:     "3306",
		MySQLUser:     "root",
		MySQLPassword: "",
		MySQLDataBase: "bytesmart",
	}
	fmt.Printf("%+v\n", EnvCfg)
/* 	if err := godotenv.Load(); err != nil {
		log.Errorf("Can not read env from file system, please check the right this program owned.")
	}

	EnvCfg = envConfig{}

	if err := env.Parse(&EnvCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	} */
}