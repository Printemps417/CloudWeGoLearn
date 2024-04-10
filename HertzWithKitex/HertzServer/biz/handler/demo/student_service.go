// Code generated by hertz generator.

package demo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"hertz.demo/biz/model/demo"
	"hertz.demo/rpc/kitex_gen/demo/studentservice"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	student "hertz.demo/rpc/kitex_gen/demo"
)

var students sync.Map
var (
	cli studentservice.Client
)

func init() {
	var err error
	cli, err = studentservice.NewClient("kitex.demo", client.WithHostPorts("0.0.0.0:9999"))
	if err != nil {
		panic(err)
	}
}

// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req student.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Tracef("add student info:%v", req)
	resp, err := cli.Register(ctx, &req)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	//c.BindAndValidate不能跨类解析url中的数据，但能跨类解析body中的数据
	var reqc demo.QueryReq
	hlog.Tracef("Query receive:", c)
	err = c.BindAndValidate(&reqc)
	hlog.Tracef("Query request:", reqc)
	req := student.NewQueryReq()
	req.Id = reqc.ID
	//req := student.NewQueryReq()
	//hlog.Tracef("Query receive:", c)
	//err = c.BindAndValidate(req)
	//hlog.Tracef("Query request:", req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := cli.Query(ctx, req)
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(consts.StatusOK, resp)
}
