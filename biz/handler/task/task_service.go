// Code generated by hertz generator.

package task

import (
	"context"
	"fmt"
	"hertz-mylist/biz/router/middleware"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	task "hertz-mylist/biz/model/task"
)

// CreateTask .
// @router /task/create [POST]
func CreateTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(task.TaskResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateTask .
// @router /task/update [PUT]
func UpdateTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(task.TaskResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteTask .
// @router /task/delete/:id [DELETE]
func DeleteTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(task.TaskResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetTaskDetail .
// @router /task/getDetail/:id [GET]
func GetTaskDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	claim, _ := c.Get(middleware.IdentityKey)
	fmt.Println(claim.(*middleware.Claim).ID)
	resp := new(task.TaskResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetTaskList .
// @router /task/getList/:pageSize/:pageNum [GET]
func GetTaskList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(task.TaskListResponse)

	c.JSON(consts.StatusOK, resp)
}
