package main

// TODO refactor

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	gc "google.golang.org/grpc/codes"
	gs "google.golang.org/grpc/status"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	h2r "github.com/takanoriyanagitani/go-conv-helper/examples/external/net/grpc/python/color/proto/color/hsv/v1"
)

type Hsv struct {
	hue        float32
	saturation float32
	value      float32
}

func (h Hsv) toProto() *h2r.Hsv {
	return &h2r.Hsv{
		Hue:        h.hue,
		Saturation: h.saturation,
		Value:      h.value,
	}
}

func (h Hsv) toRequest() *h2r.FromHsvRequest {
	var hsv *h2r.Hsv = h.toProto()
	return &h2r.FromHsvRequest{
		Hsv: hsv,
	}
}

func (r Rgb) fromProto(p *h2r.Rgb) Rgb {
	return Rgb{
		red:   p.GetRed(),
		blue:  p.GetBlue(),
		green: p.GetGreen(),
	}
}

func (r Rgb) fromResponse(p *h2r.FromHsvResponse) Rgb {
	var rgb *h2r.Rgb = p.GetRgb()
	return r.fromProto(rgb)
}

type Rgb struct {
	red   float32
	green float32
	blue  float32
}

func (r Rgb) convert() *h2r.Rgb {
	return &h2r.Rgb{
		Red:   r.red,
		Green: r.green,
		Blue:  r.blue,
	}
}

type HsvCmdSvcServer struct {
	agent *Agent
	wait  time.Duration
	h2r.UnimplementedHsvCmdServiceServer
}

func (c HsvCmdSvcServer) Convert(
	ctx context.Context,
	req *h2r.HsvCmd_ConvertRequest,
) (*h2r.HsvCmd_ConvertResponse, error) {
	_, ereq := c.validate(req)
	if nil != ereq {
		return nil, ereq
	}
	var m *h2r.Meta = req.GetMeta()
	var id *h2r.Uuid = m.GetReplyId()
	var h *h2r.Hsv = req.GetHsv()
	log.Printf("req: %v\n", req)
	c.agent.SetHsv(id, h)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	r, econv := c.waitRgb(
		ctx,
		req,
		c.wait,
	)
	var reply *h2r.HsvCmd_ConvertResponse = &h2r.HsvCmd_ConvertResponse{Rgb: r}
	return reply, errors.Join(ereq, econv)
}

func (c HsvCmdSvcServer) waitRgb(
	ctx context.Context,
	req *h2r.HsvCmd_ConvertRequest,
	wait time.Duration,
) (*h2r.Rgb, error) {
	var m *h2r.Meta = req.GetMeta()
	var id *h2r.Uuid = m.GetReplyId()
	for {
		if nil != ctx.Err() {
			log.Printf("ctx err: %v\n", ctx.Err())
			return nil, gs.Errorf(
				gc.Canceled,
				"detail: %v",
				ctx.Err(),
			)
		}
		r := c.agent.GetRgb(id)
		if nil != r {
			return r, nil
		}
		time.Sleep(wait)
	}
}

func (c HsvCmdSvcServer) validate(req *h2r.HsvCmd_ConvertRequest) (*h2r.HsvCmd_ConvertRequest, error) {
	_, e := dummyValidateMeta(req.GetMeta())
	if nil != e {
		return nil, gs.Errorf(
			gc.InvalidArgument,
			"invalid meta: %v",
			e,
		)
	}
	var h *h2r.Hsv = req.GetHsv()
	if nil == h {
		return nil, gs.Errorf(
			gc.InvalidArgument,
			"invalid hsv: %v\n",
			h,
		)
	}
	return req, e
}

func (c HsvCmdSvcServer) toHsv(h *h2r.Hsv) Hsv {
	return Hsv{
		hue:        h.GetHue(),
		saturation: h.GetSaturation(),
		value:      h.GetValue(),
	}
}

func (c HsvCmdSvcServer) fromRgb(r Rgb) *h2r.Rgb {
	return &h2r.Rgb{
		Red:   r.red,
		Blue:  r.blue,
		Green: r.green,
	}
}

func dummyValidateUuid(u *h2r.Uuid) (*h2r.Uuid, error) {
	var id uint64 = dummyUuidConverter(u.GetHi(), u.GetLo())
	switch id {
	case 0:
		return nil, gs.Errorf(
			gc.InvalidArgument,
			"invalid uuid: %v",
			u,
		)
	default:
		return u, nil
	}
}

func dummyValidateMeta(m *h2r.Meta) (*h2r.Meta, error) {
	_, e := dummyValidateUuid(m.GetReplyId())
	return m, e
}

func (c HsvCmdSvcServer) AsIf() h2r.HsvCmdServiceServer { return c }

type HsvEvtSvcServer struct {
	agent *Agent
	h2r.UnimplementedHsvEvtServiceServer
}

type Uuid struct{ lh uint64 }

func (u Uuid) convert() *h2r.Uuid {
	return &h2r.Uuid{
		Hi: 0xffffffff,
		Lo: u.lh,
	}
}

type Meta struct{ reply Uuid }

func MetaNew(m *h2r.Meta) (Meta, error) {
	var rep *h2r.Uuid = m.GetReplyId()
	if nil == rep {
		return Meta{}, gs.Errorf(
			gc.InvalidArgument,
			"invalid meta: %v",
			m,
		)
	}
	var raw uint64 = dummyUuidConverter(rep.GetHi(), rep.GetLo())
	if 0 == raw {
		return Meta{}, gs.Errorf(
			gc.InvalidArgument,
			"invalid meta: %v",
			m,
		)
	}
	return Meta{
		reply: Uuid{
			lh: raw,
		},
	}, nil
}

type ConvertedEvt struct {
	Meta
	Rgb
}

func RgbNew(r *h2r.Rgb) Rgb {
	return Rgb{
		red:   r.GetRed(),
		green: r.GetGreen(),
		blue:  r.GetBlue(),
	}
}

func ConvertedEvtNew(evt *h2r.HsvEvt_ConvertedRequest) (valid ConvertedEvt, e error) {
	var r *h2r.Rgb = evt.GetConverted()
	if nil == r {
		return valid, gs.Errorf(
			gc.InvalidArgument,
			"invalid request: %v",
			evt,
		)
	}
	var m *h2r.Meta = evt.GetMeta()
	mv, em := MetaNew(m)
	return ConvertedEvt{
		Meta: mv,
		Rgb:  RgbNew(r),
	}, em
}

func (e HsvEvtSvcServer) Converted(
	ctx context.Context,
	evt *h2r.HsvEvt_ConvertedRequest,
) (*h2r.HsvEvt_ConvertedResponse, error) {
	valid, err := ConvertedEvtNew(evt)
	if nil != err {
		return nil, err
	}
	var r Rgb = valid.Rgb
	var hr *h2r.Rgb = r.convert()

	var id Uuid = valid.Meta.reply
	var hi *h2r.Uuid = id.convert()
	e.agent.SetRgb(hi, hr)
	return &h2r.HsvEvt_ConvertedResponse{
		Sent: &h2r.Timestamp{
			UnixtimeMicros: time.Now().UnixMicro(),
		},
	}, nil
}

func (e HsvEvtSvcServer) AsIf() h2r.HsvEvtServiceServer { return e }

type HsvCmdRecvSvcServer struct {
	agent *Agent
	h2r.UnimplementedHsvCmdRecvServiceServer
}

func (r HsvCmdRecvSvcServer) GetCmd(
	ctx context.Context,
	req *h2r.HsvOrder_GetCmdRequest,
) (*h2r.HsvOrder_GetCmdResponse, error) {
	var hasReq bool = 0 < r.agent.HsvSize()
	if !hasReq {
		return nil, gs.Errorf(
			gc.NotFound,
			"no commands. request id: %v",
			req.GetRequestId(),
		)
	}
	var h *h2r.Hsv = r.agent.PopHsv()
	if nil == h {
		return nil, gs.Errorf(
			gc.Unknown,
			"req found: %v",
			hasReq,
		)
	}
	var reply *h2r.HsvOrder_GetCmdResponse = &h2r.HsvOrder_GetCmdResponse{
		Commands: &h2r.HsvOrder_Commands{
			Cmd: &h2r.HsvOrder_Commands_Convert{
				Convert: &h2r.HsvCmd_ConvertRequest{
					Meta: &h2r.Meta{
						ReplyId: &h2r.Uuid{
							Hi: 0xffff,
							Lo: 0xf00d,
						},
					},
					Hsv: h,
				},
			},
		},
	}
	return reply, nil
}

func (r HsvCmdRecvSvcServer) AsIf() h2r.HsvCmdRecvServiceServer { return r }

func dummyUuidConverter(hi uint64, lo uint64) uint64 { return hi & lo }

type HsvAgent struct {
	msg map[uint64]*h2r.Hsv
}

type RgbAgent struct {
	msg map[uint64]*h2r.Rgb
}

type Agent struct {
	hsv HsvAgent
	rgb RgbAgent
	lck sync.RWMutex
}

func (a *Agent) GetHsv(id *h2r.Uuid) *h2r.Hsv {
	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.RLock()
	hsv, ok := a.hsv.msg[i]
	a.lck.RUnlock()

	if ok {
		return hsv
	}
	return nil
}

func (a *Agent) PopHsv() *h2r.Hsv {
	a.lck.Lock()
	defer a.lck.Unlock()
	for id, hsv := range a.hsv.msg {
		var ret *h2r.Hsv = hsv
		delete(a.hsv.msg, id)
		if nil == ret {
			continue
		}
		return ret
	}
	return nil
}

func (a *Agent) GetRgb(id *h2r.Uuid) *h2r.Rgb {
	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.RLock()
	rgb, ok := a.rgb.msg[i]
	a.lck.RUnlock()

	if ok {
		return rgb
	}
	return nil
}

func (a *Agent) HasRgb(id *h2r.Uuid) (found bool) {
	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.RLock()
	_, ok := a.rgb.msg[i]
	a.lck.RUnlock()
	return ok
}

func (a *Agent) HasHsv(id *h2r.Uuid) (found bool) {
	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.RLock()
	_, ok := a.hsv.msg[i]
	a.lck.RUnlock()
	return ok
}

func (a *Agent) DelRgb(id *h2r.Uuid) (found bool) {
	found = a.HasRgb(id)
	if !found {
		return false
	}

	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.Lock()
	delete(a.rgb.msg, i)
	a.lck.Unlock()
	return true
}

func (a *Agent) DelHsv(id *h2r.Uuid) (found bool) {
	found = a.HasHsv(id)
	if !found {
		return false
	}

	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.Lock()
	delete(a.hsv.msg, i)
	a.lck.Unlock()
	return true
}

func (a *Agent) SetHsv(id *h2r.Uuid, hsv *h2r.Hsv) {
	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.Lock()
	a.hsv.msg[i] = hsv
	a.lck.Unlock()
}

func (a *Agent) SetRgb(id *h2r.Uuid, rgb *h2r.Rgb) {
	var i uint64 = dummyUuidConverter(id.GetHi(), id.GetLo())

	a.lck.Lock()
	a.rgb.msg[i] = rgb
	a.lck.Unlock()
}

func (a *Agent) HsvSize() int {
	a.lck.RLock()
	var sz int = len(a.hsv.msg)
	a.lck.RUnlock()
	return sz
}

var hsv2rgb ch.Converter[Hsv, Rgb] = func(_ context.Context, h Hsv) (Rgb, error) {
	return Rgb{}, gs.Errorf(gc.PermissionDenied, "request: %v", h)
}

const addr string = "localhost:2151"

func main() {
	lis, e := net.Listen("tcp", addr)
	if nil != e {
		log.Fatal(e)
	}
	defer lis.Close()

	var server *grpc.Server = grpc.NewServer()

	var agent *Agent = &Agent{
		hsv: HsvAgent{
			msg: make(map[uint64]*h2r.Hsv),
		},
		rgb: RgbAgent{
			msg: make(map[uint64]*h2r.Rgb),
		},
	}

	h2r.RegisterHsvCmdRecvServiceServer(server, HsvCmdRecvSvcServer{
		agent: agent,
	})

	var wait time.Duration = 1 * time.Second

	h2r.RegisterHsvCmdServiceServer(server, HsvCmdSvcServer{
		agent: agent,
		wait:  wait,
	})

	h2r.RegisterHsvEvtServiceServer(server, HsvEvtSvcServer{
		agent: agent,
	})

	e = server.Serve(lis)
	if nil != e {
		log.Fatal(e)
	}
}
