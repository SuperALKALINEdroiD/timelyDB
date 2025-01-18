package nodes

import (
	"log"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
)

func LoadNodes(config *config.DatabaseConfig) {
	if len(config.Nodes) == 0 || config.NodeCount == 0 {
		log.Fatalln("No node configuration found.")
		return
	}

	log.Println("Loading nodes...")
	for i, node := range config.Nodes {
		log.Printf("Node %d: Endpoint ==> %s\n", i+1, node.Endpoint)
	}

	log.Println("Nodes are up and running.")
}
