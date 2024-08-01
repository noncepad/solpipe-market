package remote

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	pbj "github.com/noncepad/solpipe-market/go/proto/solanajsonrpc"
)

func (e1 external) Post(stream pbj.JsonRpc_PostServer) error {
	ctx, cancel := context.WithCancelCause(stream.Context())
	doneC := ctx.Done()
	reqC := make(chan *http.Request, 1)
	client := &http.Client{}
	go e1.loopPostRead(ctx, cancel, stream, reqC)
	var req *http.Request
	select {
	case <-doneC:
		return ctx.Err()
	case req = <-reqC:
	}

	resp, err := client.Do(req)
	if err != nil {
		cancel(err)
		return err
	}
	defer resp.Body.Close()
	err = stream.Send(&pbj.Response{
		Payload: &pbj.Response_Code{
			Code: int32(resp.StatusCode),
		},
	})
	if err != nil {
		cancel(err)
		return err
	}
	respheader := new(pbj.Header)
	respheader.Path = ""
	respheader.Headers = make(map[string]string)
	for k, v := range resp.Header {
		respheader.Headers[k] = strings.Join(v, "\n")
	}
	respheader.Parameters = make(map[string]string)
	err = stream.Send(&pbj.Response{
		Payload: &pbj.Response_Header{
			Header: respheader,
		},
	})
	if err != nil {
		cancel(err)
		return err
	}
	finishC := make(chan struct{}, 1)
	go e1.loopClose(ctx, cancel, finishC, stream, resp.Body)
	select {
	case <-doneC:
		return ctx.Err()
	case <-finishC:
	}
	cancel(nil)
	return nil
}

func (e1 external) loopClose(ctx context.Context, cancel context.CancelCauseFunc, finishC chan<- struct{}, stream pbj.JsonRpc_PostServer, body io.Reader) {
	var err error
	var buf [2048]byte
	var i int
out:
	for {
		i, err = body.Read(buf[:])
		if err == io.EOF {
			err = nil
			break out
		} else if err != nil {
			break out
		}
		b := new(pbj.Body)
		b.Body = make([]byte, i)
		copy(b.Body, buf[0:i])
		err = stream.Send(&pbj.Response{
			Payload: &pbj.Response_Body{
				Body: b,
			},
		})
		if err != nil {
			break out
		}
	}
	if err != nil {
		cancel(err)
	} else {
		// tell the other end that we have an EOF
		err = stream.Send(&pbj.Response{
			Payload: &pbj.Response_Body{
				Body: &pbj.Body{
					Body: []byte{},
				},
			},
		})
		if err != nil {
			cancel(err)
		} else {
			finishC <- struct{}{}
		}
	}
}

func (e1 external) loopPostRead(
	ctx context.Context,
	cancel context.CancelCauseFunc,
	stream pbj.JsonRpc_PostServer,
	reqC chan<- *http.Request,
) {
	doneC := ctx.Done()
	method := "POST"
	var err error
	var req *http.Request
	var n int
	var i int
	rdr, wtr := io.Pipe()
out:
	for {
		var msg *pbj.Request
		msg, err = stream.Recv()
		if err == io.EOF {
			err = nil
			break out
		} else if err != nil {
			break out
		}
		switch xmsg := msg.Payload.(type) {
		case *pbj.Request_Header:
			if req != nil {
				err = errors.New("cannot receive multiple headers")
				break out
			}
			header := xmsg.Header
			if header == nil {
				err = errors.New("blank header")
				break out
			}
			req, err = e1.config.Remote.Request(method, header.Path, header.Parameters, rdr)
			if err != nil {
				break out
			}
			// TODO: put in headers; there needs to be some security check here
			select {
			case <-doneC:
				break out
			case reqC <- req:
			}
		case *pbj.Request_Body:
			if req == nil {
				err = errors.New("cannot receive body before header")
				break out
			}
			if xmsg.Body == nil {
				err = errors.New("blank body")
				break out
			}
			if xmsg.Body.Body == nil {
				err = errors.New("blank body")
				break out
			}
			if len(xmsg.Body.Body) == 0 {
				// this is our EOF
				break out
			}
			n = 0
			for n < len(xmsg.Body.Body) {
				i, err = wtr.Write(xmsg.Body.Body[n:])
				if err == io.EOF && len(xmsg.Body.Body) <= n+i {
					err = nil
					break out
				} else if err != io.EOF {
					break out
				}
				n += i
			}
		default:
			err = fmt.Errorf("unknown message type %T", xmsg)
		}

	}
	if err != nil {
		cancel(err)
	}
}

func loopPostWrite(
	ctx context.Context,
	cancel context.CancelCauseFunc,
	stream pbj.JsonRpc_PostServer,
) {
}
