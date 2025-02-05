package nodes

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
	"github.com/SuperALKALINEdroiD/timelyDB/utils/storage"
	"google.golang.org/grpc"
)

type internalServer struct {
	UnimplementedNodeServiceServer
}

type Node struct {
	ID      string
	Address string
	Storage storage.KVStore
}

func nodeSetupTask(ctx context.Context, nodeID string, port string) (*Node, error) {
	listener, httpError := net.Listen("tcp", port)
	if httpError != nil {
		return nil, fmt.Errorf("failed to start listener: %v", httpError)
	}

	grpcServer := grpc.NewServer()
	dataStoreServer := &internalServer{}
	RegisterNodeServiceServer(grpcServer, dataStoreServer)

	stop := make(chan struct{})

	go func() {
		log.Printf("Node %s: gRPC server started", nodeID)
		if err := grpcServer.Serve(listener); err != nil {
			log.Printf("gRPC server error: %v", err)
			close(stop)
		}
	}()

	go func() {
		<-ctx.Done() // context cancelled
		log.Printf("Shutting down gRPC server for Node %s...", nodeID)
		grpcServer.GracefulStop()
		listener.Close()
		close(stop)
	}()

	return &Node{ID: nodeID, Address: port, Storage: nil}, nil
}

func LoadNodes(ctx context.Context, config *config.DatabaseConfig) []*Node {
	if len(config.Nodes) == 0 || config.NodeCount == 0 {
		log.Println("No node configuration found.")
		return []*Node{}
	}

	log.Println("Loading nodes...")

	grpcNodes := make([]*Node, len(config.Nodes))

	for i, node := range config.Nodes {
		log.Printf("Node %d: Endpoint ==> %s\n", i+1, node.Endpoint)

		var setupError error
		grpcNodes[i], setupError = nodeSetupTask(ctx, fmt.Sprintf("%d", i), node.Endpoint)
		if setupError != nil {
			log.Printf("Error setting up Node %d: %v\n", i+1, setupError)
			continue
		}
	}

	log.Println("Nodes are up and running.")
	return grpcNodes
}
