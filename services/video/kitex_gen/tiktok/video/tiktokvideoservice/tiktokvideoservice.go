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
		"putVideo":           kitex.NewMethodInfo(putVideoHandler, newPutVideoArgs, newPutVideoResult, false),
		"deleteVideo":        kitex.NewMethodInfo(deleteVideoHandler, newDeleteVideoArgs, newDeleteVideoResult, false),
		"getOneVideoInfo":    kitex.NewMethodInfo(getOneVideoInfoHandler, newGetOneVideoInfoArgs, newGetOneVideoInfoResult, false),
		"downloadOneVideo":   kitex.NewMethodInfo(downloadOneVideoHandler, newDownloadOneVideoArgs, newDownloadOneVideoResult, false),
		"downloadMultiVideo": kitex.NewMethodInfo(downloadMultiVideoHandler, newDownloadMultiVideoArgs, newDownloadMultiVideoResult, false),
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

func getOneVideoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.GetOneVideoInfoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.TiktokVideoService).GetOneVideoInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetOneVideoInfoArgs:
		success, err := handler.(video.TiktokVideoService).GetOneVideoInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetOneVideoInfoResult)
		realResult.Success = success
	}
	return nil
}
func newGetOneVideoInfoArgs() interface{} {
	return &GetOneVideoInfoArgs{}
}

func newGetOneVideoInfoResult() interface{} {
	return &GetOneVideoInfoResult{}
}

type GetOneVideoInfoArgs struct {
	Req *video.GetOneVideoInfoRequest
}

func (p *GetOneVideoInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.GetOneVideoInfoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetOneVideoInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetOneVideoInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetOneVideoInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetOneVideoInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetOneVideoInfoArgs) Unmarshal(in []byte) error {
	msg := new(video.GetOneVideoInfoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetOneVideoInfoArgs_Req_DEFAULT *video.GetOneVideoInfoRequest

func (p *GetOneVideoInfoArgs) GetReq() *video.GetOneVideoInfoRequest {
	if !p.IsSetReq() {
		return GetOneVideoInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetOneVideoInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetOneVideoInfoResult struct {
	Success *video.GetOneVideoInfoResponse
}

var GetOneVideoInfoResult_Success_DEFAULT *video.GetOneVideoInfoResponse

func (p *GetOneVideoInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.GetOneVideoInfoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetOneVideoInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetOneVideoInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetOneVideoInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetOneVideoInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetOneVideoInfoResult) Unmarshal(in []byte) error {
	msg := new(video.GetOneVideoInfoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetOneVideoInfoResult) GetSuccess() *video.GetOneVideoInfoResponse {
	if !p.IsSetSuccess() {
		return GetOneVideoInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetOneVideoInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.GetOneVideoInfoResponse)
}

func (p *GetOneVideoInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func downloadOneVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.DownloadOneVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.TiktokVideoService).DownloadOneVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DownloadOneVideoArgs:
		success, err := handler.(video.TiktokVideoService).DownloadOneVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DownloadOneVideoResult)
		realResult.Success = success
	}
	return nil
}
func newDownloadOneVideoArgs() interface{} {
	return &DownloadOneVideoArgs{}
}

func newDownloadOneVideoResult() interface{} {
	return &DownloadOneVideoResult{}
}

type DownloadOneVideoArgs struct {
	Req *video.DownloadOneVideoRequest
}

func (p *DownloadOneVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.DownloadOneVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DownloadOneVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DownloadOneVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DownloadOneVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DownloadOneVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DownloadOneVideoArgs) Unmarshal(in []byte) error {
	msg := new(video.DownloadOneVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DownloadOneVideoArgs_Req_DEFAULT *video.DownloadOneVideoRequest

func (p *DownloadOneVideoArgs) GetReq() *video.DownloadOneVideoRequest {
	if !p.IsSetReq() {
		return DownloadOneVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DownloadOneVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type DownloadOneVideoResult struct {
	Success *video.DownloadOneVideoResponse
}

var DownloadOneVideoResult_Success_DEFAULT *video.DownloadOneVideoResponse

func (p *DownloadOneVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.DownloadOneVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DownloadOneVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DownloadOneVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DownloadOneVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DownloadOneVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DownloadOneVideoResult) Unmarshal(in []byte) error {
	msg := new(video.DownloadOneVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DownloadOneVideoResult) GetSuccess() *video.DownloadOneVideoResponse {
	if !p.IsSetSuccess() {
		return DownloadOneVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DownloadOneVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.DownloadOneVideoResponse)
}

func (p *DownloadOneVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func downloadMultiVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.DownloadMultiVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.TiktokVideoService).DownloadMultiVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DownloadMultiVideoArgs:
		success, err := handler.(video.TiktokVideoService).DownloadMultiVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DownloadMultiVideoResult)
		realResult.Success = success
	}
	return nil
}
func newDownloadMultiVideoArgs() interface{} {
	return &DownloadMultiVideoArgs{}
}

func newDownloadMultiVideoResult() interface{} {
	return &DownloadMultiVideoResult{}
}

type DownloadMultiVideoArgs struct {
	Req *video.DownloadMultiVideoRequest
}

func (p *DownloadMultiVideoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.DownloadMultiVideoRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DownloadMultiVideoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DownloadMultiVideoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DownloadMultiVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DownloadMultiVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DownloadMultiVideoArgs) Unmarshal(in []byte) error {
	msg := new(video.DownloadMultiVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DownloadMultiVideoArgs_Req_DEFAULT *video.DownloadMultiVideoRequest

func (p *DownloadMultiVideoArgs) GetReq() *video.DownloadMultiVideoRequest {
	if !p.IsSetReq() {
		return DownloadMultiVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DownloadMultiVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type DownloadMultiVideoResult struct {
	Success *video.DownloadMultiVideoResponse
}

var DownloadMultiVideoResult_Success_DEFAULT *video.DownloadMultiVideoResponse

func (p *DownloadMultiVideoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.DownloadMultiVideoResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DownloadMultiVideoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DownloadMultiVideoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DownloadMultiVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DownloadMultiVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DownloadMultiVideoResult) Unmarshal(in []byte) error {
	msg := new(video.DownloadMultiVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DownloadMultiVideoResult) GetSuccess() *video.DownloadMultiVideoResponse {
	if !p.IsSetSuccess() {
		return DownloadMultiVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DownloadMultiVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.DownloadMultiVideoResponse)
}

func (p *DownloadMultiVideoResult) IsSetSuccess() bool {
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

func (p *kClient) GetOneVideoInfo(ctx context.Context, Req *video.GetOneVideoInfoRequest) (r *video.GetOneVideoInfoResponse, err error) {
	var _args GetOneVideoInfoArgs
	_args.Req = Req
	var _result GetOneVideoInfoResult
	if err = p.c.Call(ctx, "getOneVideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadOneVideo(ctx context.Context, Req *video.DownloadOneVideoRequest) (r *video.DownloadOneVideoResponse, err error) {
	var _args DownloadOneVideoArgs
	_args.Req = Req
	var _result DownloadOneVideoResult
	if err = p.c.Call(ctx, "downloadOneVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadMultiVideo(ctx context.Context, Req *video.DownloadMultiVideoRequest) (r *video.DownloadMultiVideoResponse, err error) {
	var _args DownloadMultiVideoArgs
	_args.Req = Req
	var _result DownloadMultiVideoResult
	if err = p.c.Call(ctx, "downloadMultiVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
