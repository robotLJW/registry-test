package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once           sync.Once
	NacosClientCfg *NacosClientConfig
	NacosServerCfg *NacosServerConfig
	NacosKVCfg     *KVConfig
	CalculateKVCfg *KVConfig
)

func ReadNacosConfig() {
	once.Do(func() {
		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		err = readConfigFile("nacos", "yaml", currentDirectory+"/configs")
		if err != nil {
			log.Fatal(err)
		}
		NacosServerCfg = &NacosServerConfig{
			Addresses: viper.GetString("nacosConfig.serverConfig.addresses"),
			Port:      viper.GetUint64("nacosConfig.serverConfig.port"),
		}
		NacosClientCfg = &NacosClientConfig{
			NamespaceId: viper.GetString("nacosConfig.clientConfig.namespaceId"),
			Timeout:     viper.GetUint64("nacosConfig.clientConfig.timeout"),
			LogDir:      viper.GetString("nacosConfig.clientConfig.logDir"),
			CacheDir:    viper.GetString("nacosConfig.clientConfig.cacheDir"),
			RotateTime:  viper.GetDuration("nacosConfig.clientConfig.rotateTime"),
			MaxAge:      viper.GetInt("nacosConfig.clientConfig.maxAge"),
			LogLevel:    viper.GetString("nacosConfig.clientConfig.logLevel"),
		}
		NacosKVCfg = &KVConfig{
			DataId:      viper.GetString("nacosConfig.kvConfig.dataId"),
			Group:       viper.GetString("nacosConfig.kvConfig.group"),
			PostAddress: viper.GetString("nacosConfig.kvConfig.postAddress"),
		}
	})
}

func ReadCalculateConfig() {
	once.Do(func() {
		currentDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		err = readConfigFile("calculate", "yaml", currentDirectory+"/configs")
		if err != nil {
			log.Fatal(err)
		}
		NacosServerCfg = &NacosServerConfig{
			Addresses: viper.GetString("nacosConfig.serverConfig.addresses"),
			Port:      viper.GetUint64("nacosConfig.serverConfig.port"),
		}
		NacosClientCfg = &NacosClientConfig{
			NamespaceId: viper.GetString("nacosConfig.clientConfig.namespaceId"),
			Timeout:     viper.GetUint64("nacosConfig.clientConfig.timeout"),
			LogDir:      viper.GetString("nacosConfig.clientConfig.logDir"),
			CacheDir:    viper.GetString("nacosConfig.clientConfig.cacheDir"),
			RotateTime:  viper.GetDuration("nacosConfig.clientConfig.rotateTime"),
			MaxAge:      viper.GetInt("nacosConfig.clientConfig.maxAge"),
			LogLevel:    viper.GetString("nacosConfig.clientConfig.logLevel"),
		}
		CalculateKVCfg = &KVConfig{
			DataId:      viper.GetString("nacosConfig.kvConfig.dataId"),
			Group:       viper.GetString("nacosConfig.kvConfig.group"),
			Content:     viper.GetString("nacosConfig.kvConfig.content"),
			Interval:    viper.GetDuration("nacosConfig.kvConfig.interval"),
			PostAddress: viper.GetString("nacosConfig.kvConfig.postAddress"),
			Channel:     viper.GetInt("nacosConfig.kvConfig.channel"),
		}
	})
}

func readConfigFile(fileName string, fileType string, filePath string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return err
}
