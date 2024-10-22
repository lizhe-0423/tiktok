// Code generated by Kitex v0.4.4. DO NOT EDIT.
package tiktokauthservice

import (
	server "github.com/cloudwego/kitex/server"
	auth "github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler auth.TiktokAuthService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
