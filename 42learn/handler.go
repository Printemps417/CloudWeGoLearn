package main

import (
	"context"
	api "demo/kitex_gen/api"
	"github.com/cloudwego/kitex/pkg/klog"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	klog.Info("echo called")
	return &api.Response{Message: req.Message}, nil
}
