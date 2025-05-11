// Implement a capacity grpc server using cpu and ram percentages for capacity.
package cpumeter

import (
	"context"
	"io"
	"log"
	"sync/atomic"
	"time"

	pbr "github.com/noncepad/solpipe-market/go/proto/relay"
	"google.golang.org/grpc"
)

type external struct {
	pbr.UnimplementedCapacityServer
	ctx         context.Context
	cancel      context.CancelCauseFunc
	isContainer bool
	subIndex    *uint32
	subC        chan<- subG
	internalC   chan<- func(*internal)
}

type subG struct {
	id     uint32
	ctx    context.Context
	cancel context.CancelFunc
	outC   chan<- float32
}

func Add(parentCtx context.Context, s *grpc.Server) error {
	var err error
	var isContainer bool
	isContainer, err = IsContainer()
	if err != nil {
		return err
	}
	subIndex := new(uint32)
	*subIndex = 0
	internalC := make(chan func(*internal), 10)
	subC := make(chan subG, 10)
	ctx, cancel := context.WithCancelCause(parentCtx)
	e1 := external{
		ctx:         ctx,
		cancel:      cancel,
		isContainer: isContainer,
		subIndex:    subIndex,
	}
	go loopInternal(ctx, cancel, internalC, subC, subIndex)
	pbr.RegisterCapacityServer(s, e1)
	return nil
}

const CHECK_INTERVAL time.Duration = 30 * time.Second

func (e1 external) OnStatus(req *pbr.CapacityRequest, stream pbr.Capacity_OnStatusServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	doneC := ctx.Done()
	outC := make(chan float32, 100)
	s := subG{
		id:     atomic.AddUint32(e1.subIndex, 1),
		ctx:    ctx,
		cancel: cancel,
		outC:   outC,
	}
	var err error
	select {
	case <-doneC:
		err = e1.ctx.Err()
	case e1.subC <- s:
	}
	if err != nil {
		return err
	}

	var ur float32
out:
	for {
		select {
		case <-doneC:
			break out
		case ur = <-outC:
		}
		err = stream.Send(&pbr.CapacityStatus{
			UtilizationRatio: ur,
		})
		if err == io.EOF {
			err = nil
			break out
		} else if err != nil {
			break out
		}
	}
	log.Print("exiting cpu meter subscription")
	return err
}
