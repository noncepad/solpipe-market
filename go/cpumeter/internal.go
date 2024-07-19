package cpumeter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type internal struct {
	ctx              context.Context
	closeSignalCList []chan<- error
	isContainer      bool
	utilizationRatio float32
	subIndex         *uint32
	subM             map[uint32]*subG
}

func loopInternal(
	ctx context.Context,
	cancel context.CancelCauseFunc,
	internalC <-chan func(*internal),
	subC <-chan subG,
	subIndex *uint32,
) {
	doneC := ctx.Done()

	timeC := time.After(1 * time.Second)
	in := new(internal)
	in.ctx = ctx
	in.closeSignalCList = make([]chan<- error, 0)
	in.subIndex = subIndex
	in.subM = make(map[uint32]*subG)
	var err error
out:
	for {
		select {
		case <-timeC:
			err = in.check()
			if err != nil {
				break out
			}
			// broadcast the values out
			for _, sub := range in.subM {
				select {
				case sub.outC <- in.utilizationRatio:
				default:
					log.Print("subscription %d buffer filled up", sub.id)
					sub.cancel()
					delete(in.subM, sub.id)
				}
			}
			timeC = time.After(CHECK_INTERVAL)
		case <-doneC:
			break out
		case req := <-internalC:
			req(in)
		case s := <-subC:
			in.subM[s.id] = &s
		}
	}
	in.finish(err)
	cancel(err)
}

func (in *internal) finish(err error) {
	log.Printf("exiting loopInternal cpumeter: %s", err)
	for _, errorC := range in.closeSignalCList {
		errorC <- err
	}
}

func (in *internal) check() error {
	var err error
	in.utilizationRatio, err = in.unsafeRatio()
	if err != nil {
		return err
	}
	return nil
}

func (in *internal) unsafeRatio() (float32, error) {
	var ratio float32
	var err error
	if in.isContainer {
		// TODO: find a way to get the container specific usage numbers, not the host
		ratio, err = in.unsafeRatioHost()
	} else {
		ratio, err = in.unsafeRatioHost()
	}
	if err != nil {
		return 1, err
	}
	return ratio / 100, nil
}

func (in *internal) unsafeRatioHost() (ratio float32, err error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		err = fmt.Errorf("failed to read memory: %s", err)
		return
	}
	x, err := cpu.PercentWithContext(in.ctx, CHECK_INTERVAL/10, false)
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
