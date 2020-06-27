package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	Name string	`json:"name"`
	Host string	`json:"host"`
	Port int	`json:"port"`
	Username string	`json:"username"`
	Password string	`json:"password"`
}
type RedisConfig struct {
	Host string	`json:"host"`
	Port int	`json:"port"`
	Pass string	`json:"pass"`
	DBIndex int	`json:"db_index"`
}
type Config struct {
	DEBUG bool
	Port int
	Host string
	DB *DBConfig
	Redis *RedisConfig
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
	return  m
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
		Host: "127.0.0.1",
		Port: 6379,
		Pass: "",
		DBIndex: 10,
	}
	fmt.Printf("db: %v  redis: %v", db, redis)
	config := &Config{
		DEBUG: false,
		Port:  10000,
		Host:  "",
		DB:    &db,
		Redis:  &redis,
	}
	var key string
	if key = os.Getenv("DEBUG"); key != "" {
		if b, err := strconv.ParseBool(key); err == nil {
			config.DEBUG = b
		}
	}
	if key = os.Getenv("HOST"); key != "" {
		config.Host = key
	}
	if key = os.Getenv("PORT"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			config.Port = i
		}
	}
	// mongo
	if key = os.Getenv("DB_HOST"); key != "" {
		config.DB.Host = key
	}
	if key = os.Getenv("DB_PORT"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			config.DB.Port = i
		}
	}
	if key = os.Getenv("DB_USER"); key != "" {
		config.DB.Username = key
	}
	if key = os.Getenv("DB_PWD"); key != "" {
		config.DB.Password = key
	}
	if key = os.Getenv("DB_NAME"); key != "" {
		config.DB.Name = key
	}
	// redis
	if key = os.Getenv("REDIS_HOST"); key != "" {
		config.Redis.Host = key
	}
	if key = os.Getenv("REDIS_PORT"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			config.Redis.Port = i
		}
	}
	if key = os.Getenv("REDIS_PWD"); key != "" {
		config.Redis.Pass = key
	}
	if key = os.Getenv("REDIS_DB"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			config.Redis.DBIndex = i
		}
	}
	
	ConfigData = config
	fmt.Println("ConfigData:", ConfigData)
	return config, nil
}

