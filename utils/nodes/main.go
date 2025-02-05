package nodes

import (
	context "context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
	"google.golang.org/grpc"
)

type internalServer struct {
	UnimplementedNodeServiceServer
}

type Node struct {
	ID string
}

func NodeSetupTask(ctx context.Context, nodeID string) (*Node, error) {
	listener, httpError := net.Listen("tcp", ":5000") // TODO
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
		}
		close(stop)
	}()

	go func() {
		<-ctx.Done() // context cancelled
		log.Printf("Shutting down gRPC server for Node %s...", nodeID)
		grpcServer.GracefulStop()
		listener.Close()
		close(stop)
	}()

	return &Node{ID: nodeID}, nil
}

func LoadNodes(config *config.DatabaseConfig) []*Node {
	if len(config.Nodes) == 0 || config.NodeCount == 0 {
		log.Println("No node configuration found.")
		return []*Node{}
	}

	log.Println("Loading nodes...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel) // capture all signals on channel

	go func() {
		<-signalChannel
		log.Println("Received shutdown signal. Stopping servers...")
		cancel() // context cancelled: shutdown the servers
	}()

	grpcNodes := make([]*Node, len(config.Nodes))

	for i, node := range config.Nodes {
		log.Printf("Node %d: Endpoint ==> %s\n", i+1, node.Endpoint)

		var setupError error
		grpcNodes[i], setupError = NodeSetupTask(ctx, fmt.Sprintf("%d", i))
		if setupError != nil {
			log.Printf("Error setting up Node %d: %v\n", i+1, setupError)
			continue
		}
	}

	log.Println("Nodes are up and running.")
	<-ctx.Done()

	log.Println("All gRPC servers have been shut down.")
	return grpcNodes
}
