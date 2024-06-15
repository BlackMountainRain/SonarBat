// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.21.12
// source: task/v1/task.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTaskCreateTask = "/api.task.v1.Task/CreateTask"
const OperationTaskDeleteTask = "/api.task.v1.Task/DeleteTask"
const OperationTaskGetTask = "/api.task.v1.Task/GetTask"
const OperationTaskGetTasks = "/api.task.v1.Task/GetTasks"
const OperationTaskOverwriteTask = "/api.task.v1.Task/OverwriteTask"
const OperationTaskUpdateTask = "/api.task.v1.Task/UpdateTask"

type TaskHTTPServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskReply, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskReply, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error)
	GetTasks(context.Context, *GetTasksRequest) (*GetTasksReply, error)
	OverwriteTask(context.Context, *OverwriteTaskRequest) (*OverwriteTaskReply, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskReply, error)
}

func RegisterTaskHTTPServer(s *http.Server, srv TaskHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/tasks", _Task_CreateTask0_HTTP_Handler(srv))
	r.PATCH("/api/v1/tasks/{id}", _Task_UpdateTask0_HTTP_Handler(srv))
	r.PUT("/api/v1/tasks/{id}", _Task_OverwriteTask0_HTTP_Handler(srv))
	r.DELETE("/api/v1/tasks/{id}", _Task_DeleteTask0_HTTP_Handler(srv))
	r.GET("/api/v1/tasks/{id}", _Task_GetTask0_HTTP_Handler(srv))
	r.GET("/api/v1/tasks", _Task_GetTasks0_HTTP_Handler(srv))
}

func _Task_CreateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTask(ctx, req.(*CreateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_OverwriteTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OverwriteTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskOverwriteTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OverwriteTask(ctx, req.(*OverwriteTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OverwriteTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_DeleteTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskDeleteTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTask(ctx, req.(*DeleteTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTask(ctx, req.(*GetTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTasks0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTasksRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTasks)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTasks(ctx, req.(*GetTasksRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTasksReply)
		return ctx.Result(200, reply)
	}
}

type TaskHTTPClient interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest, opts ...http.CallOption) (rsp *CreateTaskReply, err error)
	DeleteTask(ctx context.Context, req *DeleteTaskRequest, opts ...http.CallOption) (rsp *DeleteTaskReply, err error)
	GetTask(ctx context.Context, req *GetTaskRequest, opts ...http.CallOption) (rsp *GetTaskReply, err error)
	GetTasks(ctx context.Context, req *GetTasksRequest, opts ...http.CallOption) (rsp *GetTasksReply, err error)
	OverwriteTask(ctx context.Context, req *OverwriteTaskRequest, opts ...http.CallOption) (rsp *OverwriteTaskReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *UpdateTaskReply, err error)
}

type TaskHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskHTTPClient(client *http.Client) TaskHTTPClient {
	return &TaskHTTPClientImpl{client}
}

func (c *TaskHTTPClientImpl) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...http.CallOption) (*CreateTaskReply, error) {
	var out CreateTaskReply
	pattern := "/api/v1/tasks"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TaskHTTPClientImpl) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...http.CallOption) (*DeleteTaskReply, error) {
	var out DeleteTaskReply
	pattern := "/api/v1/tasks/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskDeleteTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TaskHTTPClientImpl) GetTask(ctx context.Context, in *GetTaskRequest, opts ...http.CallOption) (*GetTaskReply, error) {
	var out GetTaskReply
	pattern := "/api/v1/tasks/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TaskHTTPClientImpl) GetTasks(ctx context.Context, in *GetTasksRequest, opts ...http.CallOption) (*GetTasksReply, error) {
	var out GetTasksReply
	pattern := "/api/v1/tasks"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetTasks))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TaskHTTPClientImpl) OverwriteTask(ctx context.Context, in *OverwriteTaskRequest, opts ...http.CallOption) (*OverwriteTaskReply, error) {
	var out OverwriteTaskReply
	pattern := "/api/v1/tasks/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskOverwriteTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TaskHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*UpdateTaskReply, error) {
	var out UpdateTaskReply
	pattern := "/api/v1/tasks/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}