package config

import "time"

type NacosServerConfig struct {
	Addresses string
	Port      uint64
}

type NacosClientConfig struct {
	NamespaceId string
	Timeout     uint64
	LogDir      string
	CacheDir    string
	RotateTime  time.Duration
	MaxAge      int
	LogLevel    string
}

type KVConfig struct {
	DataId      string
	Group       string
	Content     string
	PostAddress string
	Interval    time.Duration
	Channel     int
}
