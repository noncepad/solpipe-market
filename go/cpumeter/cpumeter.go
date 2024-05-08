// Implement a capacity grpc server using cpu and ram percentages for capacity.
package cpumeter

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	pbr "github.com/noncepad/solpipe-market/go/proto/relay"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
)

type external struct {
	pbr.UnimplementedCapacityServer
	ctx              context.Context
	m                *sync.Mutex
	isContainer      bool
	utilizationRatio float32
	subIndex         int
	subM             map[int]*subG
}

type subG struct {
	id     int
	ctx    context.Context
	cancel context.CancelFunc
	outC   chan<- float32
}

func Add(ctx context.Context, s *grpc.Server) error {
	var err error

	e1 := new(external)
	e1.isContainer, err = IsContainer()
	if err != nil {
		return err
	}
	e1.ctx = ctx
	e1.m = &sync.Mutex{}
	e1.subIndex = 0
	e1.subM = make(map[int]*subG)
	e1.utilizationRatio = 0
	go e1.loopInternal()

	pbr.RegisterCapacityServer(s, e1)
	return nil
}

const CHECK_INTERVAL time.Duration = 30 * time.Second

func (e1 *external) loopInternal() {
	doneC := e1.ctx.Done()
	timeC := time.After(CHECK_INTERVAL)
	utilizationRatio := float32(0)
	isFirst := true
	var err error
out:
	for e1.ctx.Err() == nil {
		if !isFirst {
			select {
			case <-doneC:
				break out
			case <-timeC:
				timeC = time.After(CHECK_INTERVAL)
			}
		} else {
			isFirst = false
		}

		utilizationRatio, err = e1.unsafeRatio()
		if err != nil {
			log.Printf("failed to get ratio: %s", err)
			err = nil
			continue
		}

		e1.m.Lock()
		e1.utilizationRatio = utilizationRatio
		for _, sub := range e1.subM {
			if sub.ctx.Err() == nil {
				select {
				case sub.outC <- utilizationRatio:
				default:
					sub.cancel()
					delete(e1.subM, sub.id)
				}
			} else {
				sub.cancel()
				delete(e1.subM, sub.id)
			}
		}
		e1.m.Unlock()

	}
}

func (e1 *external) unsafeRatio() (float32, error) {
	var ratio float32
	var err error
	if e1.isContainer {
		// TODO: find a way to get the container specific usage numbers, not the host
		ratio, err = e1.unsafeRatioHost()
	} else {
		ratio, err = e1.unsafeRatioHost()
	}
	if err != nil {
		return 1, err
	}
	return ratio / 100, nil
}

func (e1 *external) unsafeRatioHost() (ratio float32, err error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		err = fmt.Errorf("failed to read memory: %s", err)
		return
	}
	x, err := cpu.PercentWithContext(e1.ctx, CHECK_INTERVAL/10, false)
	if err != nil {
		err = fmt.Errorf("failed to read cpu: %s", err)
		return
	}
	if len(x) != 1 {
		err = fmt.Errorf("cpu reading is wrong size: %d", len(x))
		return
	}
	if v.UsedPercent < x[0] {
		ratio = float32(x[0])
	} else {
		ratio = float32(v.UsedPercent)
	}
	log.Printf("averaging percent usage RAM and CPU: %f %f => %f", v.UsedPercent, x[0], ratio)

	return
}

func (e1 *external) OnStatus(req *pbr.Empty, stream pbr.Capacity_OnStatusServer) error {
	ctx, cancel := context.WithCancel(stream.Context())
	doneC := ctx.Done()
	e1.m.Lock()
	e1.subIndex++
	sub := new(subG)
	sub.id = e1.subIndex
	outC := make(chan float32, 10)
	sub.outC = outC
	sub.ctx = ctx
	sub.cancel = cancel
	e1.subM[sub.id] = sub
	e1.m.Unlock()
	var ur float32
	var err error
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

	e1.m.Lock()
	delete(e1.subM, sub.id)
	e1.m.Unlock()
	return err
}
