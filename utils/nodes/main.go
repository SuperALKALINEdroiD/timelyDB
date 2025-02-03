package nodes

import (
	"log"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
	"github.com/SuperALKALINEdroiD/timelyDB/utils/logs"
)

type Node struct {
	ID  string
	WAL *logs.WAL
}

func NodeSeupTask(nodeID string) (*Node, error) {
	wal, err := logs.SetupWAL(nodeID)
	if err != nil {
		return nil, err
	}

	return &Node{
		ID:  nodeID,
		WAL: wal,
	}, nil
}

func LoadNodes(config *config.DatabaseConfig) []*Node {
	if len(config.Nodes) == 0 || config.NodeCount == 0 {
		log.Fatalln("No node configuration found.")
		return nil
	}

	log.Println("Loading nodes...")

	var nodes []*Node = make([]*Node, len(config.Nodes))

	for i, node := range config.Nodes {
		log.Printf("Node %d: Endpoint ==> %s\n", i+1, node.Endpoint)

		var setupError error
		nodes[i], setupError = NodeSeupTask(string(rune(i)))
		if setupError != nil {
			log.Printf("Error setting up Node %d: %v\n", i+1, setupError)
			continue
		}
	}

	log.Println("Nodes are up and running.")
	return nodes
}
