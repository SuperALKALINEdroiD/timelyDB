package nodes

import (
	"log"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
)

type Node struct {
	ID string
}

func NodeSeupTask(nodeID string) (*Node, error) {

	return &Node{
		ID: nodeID,
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
