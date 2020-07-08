package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type RedisConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Pass    string `json:"pass"`
	DBIndex int    `json:"db_index"`
}
type QQMail struct {
	Token    string
	Sender   string
	Nickname string // 发送者
}
type Config struct {
	DEBUG  bool
	Port   int
	Host   string
	DB     *DBConfig
	Redis  *RedisConfig
	QQMail *QQMail
}

var ConfigData *Config

func (config *Config) GetPort() int {
	return config.Port
}
func (config *Config) GetHost() string {
	return config.Host
}
func (config *Config) GetDEBUG() bool {
	return config.DEBUG
}
func (config *Config) GetDB() map[string]interface{} {
	var aMap map[string]interface{}
	aByte, _ := json.Marshal(*config.DB)
	json.Unmarshal(aByte, &aMap)
	return aMap
}
func (config *Config) GetRedis() map[string]interface{} {
	data, err := json.Marshal(config.Redis)
	if err != nil {
		log.Fatalf("Config.getDB err: %s", err)
	}
	m := make(map[string]interface{})
	if err = json.Unmarshal(data, &m); err != nil {
		log.Fatalf("Config.getDB []byte 2 map err: %s", err)
	}
	return m
}

func InitConfig() (*Config, error) {
	db := DBConfig{
		Name:     "novel",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
	}
	redis := RedisConfig{
		Host:    "127.0.0.1",
		Port:    6379,
		Pass:    "",
		DBIndex: 10,
	}
	fmt.Printf("db: %v  redis: %v", db, redis)
	config := &Config{
		DEBUG: false,
		Port:  10000,
		Host:  "",
		DB: &DBConfig{
			Name:     "novel",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "123456",
		},
		Redis: &RedisConfig{},
		QQMail: &QQMail{
			Token:    "",
			Sender:   "767838865@qq.com",
			Nickname: "aker",
		},
	}
	var value string
	if value = os.Getenv("DEBUG"); value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			config.DEBUG = b
		}
	}
	if value = os.Getenv("HOST"); value != "" {
		config.Host = value
	}
	if value = os.Getenv("PORT"); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			config.Port = i
		}
	}
	// mongo
	if value = os.Getenv("DB_HOST"); value != "" {
		config.DB.Host = value
	}
	if value = os.Getenv("DB_PORT"); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			config.DB.Port = i
		}
	}
	if value = os.Getenv("DB_USER"); value != "" {
		config.DB.Username = value
	}
	if value = os.Getenv("DB_PWD"); value != "" {
		config.DB.Password = value
	}
	if value = os.Getenv("DB_NAME"); value != "" {
		config.DB.Name = value
	}
	// redis
	if value = os.Getenv("REDIS_HOST"); value != "" {
		config.Redis.Host = value
	}
	if value = os.Getenv("REDIS_PORT"); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			config.Redis.Port = i
		}
	}
	if value = os.Getenv("REDIS_PWD"); value != "" {
		config.Redis.Pass = value
	}
	if value = os.Getenv("REDIS_DB"); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			config.Redis.DBIndex = i
		}
	}
	if value = os.Getenv("QQ_MAIL_TOKEN"); value != "" {
		config.QQMail.Token = value
	}
	if value = os.Getenv("QQ_MAIL_SENDER"); value != "" {
		config.QQMail.Sender = value
	}
	if value = os.Getenv("QQ_MAIL_NICKNAME"); value != "" {
		config.QQMail.Nickname = value
	}

	ConfigData = config
	fmt.Println("ConfigData:", ConfigData)
	return config, nil
}
