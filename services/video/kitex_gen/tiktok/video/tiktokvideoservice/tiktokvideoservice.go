// Code generated by Kitex v0.4.4. DO NOT EDIT.

package tiktokvideoservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return tiktokVideoServiceServiceInfo
}

var tiktokVideoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "tiktokVideoService"
	handlerType := (*video.TiktokVideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"putVideo":    kitex.NewMethodInfo(putVideoHandler, newPutVideoArgs, newPutVideoResult, false),
		"deleteVideo": kitex.NewMethodInfo(deleteVideoHandler, newDeleteVideoArgs, newDeleteVideoResult, false),
		"getVideo":    kitex.NewMethodInfo(getVideoHandler, newGetVideoArgs, newGetVideoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func putVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PutVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.TiktokVideoService).PutVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PutVideoArgs:
		success, err := handler.(video.TiktokVideoService).PutVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PutVideoResult)
		realResult.Success = success
	}
	return nil
}
func newPutVideoArgs() interface{} {
	return &PutVideoArgs{}
}

func newPutVideoResult() interface{} {
	return &PutVideoResult{}
}

type PutVideoArgs struct {
	Req *video.PutVideoRequest
}

func (p *PutVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.PutVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PutVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PutVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PutVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PutVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PutVideoArgs) Unmarshal(in []byte) error {
	msg := new(video.PutVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PutVideoArgs_Req_DEFAULT *video.PutVideoRequest

func (p *PutVideoArgs) GetReq() *video.PutVideoRequest {
	if !p.IsSetReq() {
		return PutVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PutVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type PutVideoResult struct {
	Success *video.PutVideoResponse
}

var PutVideoResult_Success_DEFAULT *video.PutVideoResponse

func (p *PutVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.PutVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PutVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PutVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PutVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PutVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PutVideoResult) Unmarshal(in []byte) error {
	msg := new(video.PutVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PutVideoResult) GetSuccess() *video.PutVideoResponse {
	if !p.IsSetSuccess() {
		return PutVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PutVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PutVideoResponse)
}

func (p *PutVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.DeleteVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.TiktokVideoService).DeleteVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteVideoArgs:
		success, err := handler.(video.TiktokVideoService).DeleteVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteVideoResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteVideoArgs() interface{} {
	return &DeleteVideoArgs{}
}

func newDeleteVideoResult() interface{} {
	return &DeleteVideoResult{}
}

type DeleteVideoArgs struct {
	Req *video.DeleteVideoRequest
}

func (p *DeleteVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.DeleteVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteVideoArgs) Unmarshal(in []byte) error {
	msg := new(video.DeleteVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteVideoArgs_Req_DEFAULT *video.DeleteVideoRequest

func (p *DeleteVideoArgs) GetReq() *video.DeleteVideoRequest {
	if !p.IsSetReq() {
		return DeleteVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteVideoResult struct {
	Success *video.DeleteVideoResponse
}

var DeleteVideoResult_Success_DEFAULT *video.DeleteVideoResponse

func (p *DeleteVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.DeleteVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteVideoResult) Unmarshal(in []byte) error {
	msg := new(video.DeleteVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteVideoResult) GetSuccess() *video.DeleteVideoResponse {
	if !p.IsSetSuccess() {
		return DeleteVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.DeleteVideoResponse)
}

func (p *DeleteVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.GetVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.TiktokVideoService).GetVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetVideoArgs:
		success, err := handler.(video.TiktokVideoService).GetVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetVideoResult)
		realResult.Success = success
	}
	return nil
}
func newGetVideoArgs() interface{} {
	return &GetVideoArgs{}
}

func newGetVideoResult() interface{} {
	return &GetVideoResult{}
}

type GetVideoArgs struct {
	Req *video.GetVideoRequest
}

func (p *GetVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.GetVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetVideoArgs) Unmarshal(in []byte) error {
	msg := new(video.GetVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetVideoArgs_Req_DEFAULT *video.GetVideoRequest

func (p *GetVideoArgs) GetReq() *video.GetVideoRequest {
	if !p.IsSetReq() {
		return GetVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetVideoResult struct {
	Success *video.GetVideoResponse
}

var GetVideoResult_Success_DEFAULT *video.GetVideoResponse

func (p *GetVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.GetVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetVideoResult) Unmarshal(in []byte) error {
	msg := new(video.GetVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetVideoResult) GetSuccess() *video.GetVideoResponse {
	if !p.IsSetSuccess() {
		return GetVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.GetVideoResponse)
}

func (p *GetVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PutVideo(ctx context.Context, Req *video.PutVideoRequest) (r *video.PutVideoResponse, err error) {
	var _args PutVideoArgs
	_args.Req = Req
	var _result PutVideoResult
	if err = p.c.Call(ctx, "putVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteVideo(ctx context.Context, Req *video.DeleteVideoRequest) (r *video.DeleteVideoResponse, err error) {
	var _args DeleteVideoArgs
	_args.Req = Req
	var _result DeleteVideoResult
	if err = p.c.Call(ctx, "deleteVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideo(ctx context.Context, Req *video.GetVideoRequest) (r *video.GetVideoResponse, err error) {
	var _args GetVideoArgs
	_args.Req = Req
	var _result GetVideoResult
	if err = p.c.Call(ctx, "getVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}