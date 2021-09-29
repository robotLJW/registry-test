package calculate

import (
	"log"
	"strings"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"registry-test/pkg/config"
)

var StartTime int64

func Execute() error {
	client, err := getConfigClient()
	if err != nil {
		log.Fatal(err)
		return err
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: config.CalculateKVCfg.DataId,
		Group:  config.CalculateKVCfg.Group,
	})
	if err != nil {
		content = config.CalculateKVCfg.Content
	} else {
		content = content + "1"
	}
	//tick := time.NewTicker(config.CalculateKVCfg.Interval)
	StartTime = time.Now().UnixNano()
	client.PublishConfig(vo.ConfigParam{
		DataId:  config.CalculateKVCfg.DataId,
		Group:   config.CalculateKVCfg.Group,
		Content: content,
	})
	//for {
	//	select {
	//	case <-tick.C:
	//		StartTime = time.Now().UnixNano()
	//		client.PublishConfig(vo.ConfigParam{
	//			DataId:  config.CalculateKVCfg.DataId,
	//			Group:   config.CalculateKVCfg.Group,
	//			Content: content,
	//		})
	//		content = content + "1"
	//	}
	//}

	return nil
}

func getConfigClient() (config_client.IConfigClient, error) {
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
