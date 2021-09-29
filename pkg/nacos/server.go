package nacos

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"registry-test/pkg/config"
)

func Execute() error {
	client, err := getConfigClient()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = client.ListenConfig(vo.ConfigParam{
		DataId: config.NacosKVCfg.DataId,
		Group:  config.NacosKVCfg.Group,
		OnChange: func(namespace, group, dataId, data string) {
			receiveTime := time.Now().UnixNano()
			fmt.Println("-------------------------")
			fmt.Printf("data: %v\n", data)
			url := fmt.Sprintf("%s?receiveTime=%v", config.NacosKVCfg.PostAddress, receiveTime)
			fmt.Println(url)
			http.Post(url, "application/json", nil)
		},
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	for {
		time.Sleep(60 * time.Second)
	}
	return nil
}

func getConfigClient() (config_client.IConfigClient, error) {
	config.ReadNacosConfig()
	// create ServerConfig
	ipAddresses := strings.Split(config.NacosServerCfg.Addresses, ",")
	sc := make([]constant.ServerConfig, len(ipAddresses))
	for i := 0; i < len(ipAddresses); i++ {
		sc[i] = *constant.NewServerConfig(ipAddresses[i], config.NacosServerCfg.Port, constant.WithContextPath("/nacos"))
	}
	// create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(config.NacosClientCfg.NamespaceId),
		constant.WithTimeoutMs(config.NacosClientCfg.Timeout),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(config.NacosClientCfg.LogDir),
		constant.WithCacheDir(config.NacosClientCfg.CacheDir),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel(config.NacosClientCfg.LogLevel),
	)
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
		return nil, err
	}
	return client, err
}
