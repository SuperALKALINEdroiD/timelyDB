package config

import "fmt"

type DatabaseConfig struct {
	StoreName     string       `json:"dbName"`
	Port          int          `json:"port"`
	IsLockEnabled bool         `json:"isLockEnabled"`
	TimelyConfig  TimelyConfig `json:"timelyConfig"`
	Nodes         []NodeConfig `json:"nodes"`
	NodeCount     int          `json:"nodeCount"`
}

type TimelyConfig struct {
	IsEnabled bool `json:"isEnabled"`
	Type      rune `json:"type"`
}

type NodeConfig struct {
	Endpoint string `json:"endpoint"`
}

func GenerateExampleConfig(nodeCount int, host string) DatabaseConfig {
	nodes := generateNodeConfig(nodeCount, host)

	return DatabaseConfig{
		StoreName:     "example_store",
		Port:          7001,
		IsLockEnabled: true,
		TimelyConfig: TimelyConfig{
			IsEnabled: true,
			Type:      'H',
		},
		Nodes:     nodes,
		NodeCount: nodeCount,
	}
}

func generateNodeConfig(nodeCount int, host string) (nodes []NodeConfig) {
	nodes = make([]NodeConfig, nodeCount)
	for i := 0; i < nodeCount; i++ {
		nodes[i] = NodeConfig{
			Endpoint: fmt.Sprintf("%s:%d", host, 50051+i), // TODO: fix a port number
		}
	}

	return
}
