package cpumeter_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"path/filepath"
	"testing"

	"github.com/noncepad/solpipe-market/go/cpumeter"
	pbr "github.com/noncepad/solpipe-market/go/proto/relay"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCpu(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := grpc.NewServer()
	tmpdir := t.TempDir()
	path := filepath.Join(tmpdir, "server.sock")

	l, err := net.Listen("unix", path)
	if err != nil {
		t.Fatal(err)
	}
	err = cpumeter.Add(ctx, s)
	if err != nil {
		t.Fatal(err)
	}
	go s.Serve(l)

	conn, err := grpc.Dial(fmt.Sprintf("unix://%s", path), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	client := pbr.NewCapacityClient(conn)

	stream, err := client.OnStatus(ctx, &pbr.Empty{})
	if err != nil {
		t.Fatal(err)
	}
out:
	for {
		var msg *pbr.CapacityStatus
		msg, err = stream.Recv()
		if err == io.EOF {
			err = nil
			break out
		} else if err != nil {
			break out
		}
		log.Printf("utlization ratio: %f", msg.UtilizationRatio)
	}
	if err != nil {
		t.Fatal(err)
	}
}
