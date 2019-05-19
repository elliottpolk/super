package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/elliottpolk/super"
	"github.com/elliottpolk/super/config"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

func Serve(ctx context.Context, comp *config.Composition) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", comp.Server.RpcPort))
	if err != nil {
		return errors.Wrap(err, "unable to create tcp listener")
	}

	server := grpc.NewServer()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(comp.Db.ConnString()))
	if err != nil {
		return errors.Wrap(err, "unable to generate mongo client")
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "unable to verify connection to mongo")
	}

	// register services
	super.RegisterDopeServiceServer(server, super.NewDopeServer(comp, client))

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		// receiving an interrupt signal, similar to a 'Ctrl+C'
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(listener)
}
