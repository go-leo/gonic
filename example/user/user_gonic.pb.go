// Code generated by protoc-gen-gonic. DO NOT EDIT.

package user

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gonic "github.com/go-leo/gonic"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

type UserGonicService interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error)
	ModifyUser(ctx context.Context, request *ModifyUserRequest) (*ModifyUserResponse, error)
	UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error)
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
	ListUser(ctx context.Context, request *ListUserRequest) (*ListUserResponse, error)
}

func AppendUserGonicRoute[Router gin.IRoutes](router Router, service UserGonicService, opts ...gonic.Option) Router {
	options := gonic.NewOptions(opts...)
	handler := userGonicHandler{
		service: service,
		decoder: userGonicRequestDecoder{
			unmarshalOptions:        options.UnmarshalOptions(),
			shouldFailFast:          options.ShouldFailFast(),
			onValidationErrCallback: options.OnValidationErrCallback(),
		},
		encoder: userGonicEncodeResponse{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: gonic.DefaultEncodeError,
	}
	router.Match([]string{"POST"}, "/v1/user", gonic.Chain(handler.CreateUser(), options.Middlewares()...)...)
	router.Match([]string{"DELETE"}, "/v1/user/:id", gonic.Chain(handler.DeleteUser(), options.Middlewares()...)...)
	router.Match([]string{"PUT"}, "/v1/user/:id", gonic.Chain(handler.ModifyUser(), options.Middlewares()...)...)
	router.Match([]string{"PATCH"}, "/v1/user/:id", gonic.Chain(handler.UpdateUser(), options.Middlewares()...)...)
	router.Match([]string{"GET"}, "/v1/user/:id", gonic.Chain(handler.GetUser(), options.Middlewares()...)...)
	router.Match([]string{"GET"}, "/v1/users", gonic.Chain(handler.ListUser(), options.Middlewares()...)...)
	return router
}

type userGonicHandler struct {
	service      UserGonicService
	decoder      userGonicRequestDecoder
	encoder      userGonicEncodeResponse
	errorEncoder gonic.ErrorEncoder
}

func (h userGonicHandler) CreateUser() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		in, err := h.decoder.CreateUser(ctx)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		out, err := h.service.CreateUser(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		if err := h.encoder.CreateUser(ctx, out); err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
	})
}

func (h userGonicHandler) DeleteUser() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		in, err := h.decoder.DeleteUser(ctx)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		out, err := h.service.DeleteUser(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		if err := h.encoder.DeleteUser(ctx, out); err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
	})
}

func (h userGonicHandler) ModifyUser() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		in, err := h.decoder.ModifyUser(ctx)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		out, err := h.service.ModifyUser(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		if err := h.encoder.ModifyUser(ctx, out); err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
	})
}

func (h userGonicHandler) UpdateUser() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		in, err := h.decoder.UpdateUser(ctx)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		out, err := h.service.UpdateUser(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		if err := h.encoder.UpdateUser(ctx, out); err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
	})
}

func (h userGonicHandler) GetUser() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		in, err := h.decoder.GetUser(ctx)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		out, err := h.service.GetUser(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		if err := h.encoder.GetUser(ctx, out); err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
	})
}

func (h userGonicHandler) ListUser() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		in, err := h.decoder.ListUser(ctx)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		out, err := h.service.ListUser(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
		if err := h.encoder.ListUser(ctx, out); err != nil {
			h.errorEncoder(ctx, err, ctx.Writer)
			return
		}
	})
}

type userGonicRequestDecoder struct {
	unmarshalOptions        protojson.UnmarshalOptions
	shouldFailFast          bool
	onValidationErrCallback gonic.OnValidationErrCallback
}

func (decoder userGonicRequestDecoder) CreateUser(ctx *gin.Context) (*CreateUserRequest, error) {
	r := ctx.Request
	req := &CreateUserRequest{}
	if ok, err := gonic.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	if err := gonic.DecodeRequest(ctx, r, req, decoder.unmarshalOptions); err != nil {
		return nil, err
	}
	return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}
func (decoder userGonicRequestDecoder) DeleteUser(ctx *gin.Context) (*DeleteUserRequest, error) {
	r := ctx.Request
	req := &DeleteUserRequest{}
	if ok, err := gonic.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	vars := gonic.FormFromParams(ctx.Params)
	var varErr error
	req.Id, varErr = gonic.DecodeForm[int64](varErr, vars, "id", gonic.GetInt64)
	if varErr != nil {
		return nil, varErr
	}
	return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}
func (decoder userGonicRequestDecoder) ModifyUser(ctx *gin.Context) (*ModifyUserRequest, error) {
	r := ctx.Request
	req := &ModifyUserRequest{}
	if ok, err := gonic.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	if err := gonic.DecodeRequest(ctx, r, req, decoder.unmarshalOptions); err != nil {
		return nil, err
	}
	vars := gonic.FormFromParams(ctx.Params)
	var varErr error
	req.Id, varErr = gonic.DecodeForm[int64](varErr, vars, "id", gonic.GetInt64)
	if varErr != nil {
		return nil, varErr
	}
	return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}
func (decoder userGonicRequestDecoder) UpdateUser(ctx *gin.Context) (*UpdateUserRequest, error) {
	r := ctx.Request
	req := &UpdateUserRequest{}
	if ok, err := gonic.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	if req.Item == nil {
		req.Item = &UserItem{}
	}
	if err := gonic.DecodeRequest(ctx, r, req.Item, decoder.unmarshalOptions); err != nil {
		return nil, err
	}
	vars := gonic.FormFromParams(ctx.Params)
	var varErr error
	req.Id, varErr = gonic.DecodeForm[int64](varErr, vars, "id", gonic.GetInt64)
	if varErr != nil {
		return nil, varErr
	}
	return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}
func (decoder userGonicRequestDecoder) GetUser(ctx *gin.Context) (*GetUserRequest, error) {
	r := ctx.Request
	req := &GetUserRequest{}
	if ok, err := gonic.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	vars := gonic.FormFromParams(ctx.Params)
	var varErr error
	req.Id, varErr = gonic.DecodeForm[int64](varErr, vars, "id", gonic.GetInt64)
	if varErr != nil {
		return nil, varErr
	}
	return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}
func (decoder userGonicRequestDecoder) ListUser(ctx *gin.Context) (*ListUserRequest, error) {
	r := ctx.Request
	req := &ListUserRequest{}
	if ok, err := gonic.CustomDecodeRequest(ctx, r, req); ok && err != nil {
		return nil, err
	} else if ok && err == nil {
		return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
	}
	queries := r.URL.Query()
	var queryErr error
	req.PageNum, queryErr = gonic.DecodeForm[int64](queryErr, queries, "page_num", gonic.GetInt64)
	req.PageSize, queryErr = gonic.DecodeForm[int64](queryErr, queries, "page_size", gonic.GetInt64)
	if queryErr != nil {
		return nil, queryErr
	}
	return req, gonic.ValidateRequest(ctx, req, decoder.shouldFailFast, decoder.onValidationErrCallback)
}

type userGonicEncodeResponse struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer gonic.ResponseTransformer
}

func (encoder userGonicEncodeResponse) CreateUser(ctx *gin.Context, resp *CreateUserResponse) error {
	return gonic.EncodeResponse(ctx, ctx.Writer, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder userGonicEncodeResponse) DeleteUser(ctx *gin.Context, resp *DeleteUserResponse) error {
	return gonic.EncodeResponse(ctx, ctx.Writer, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder userGonicEncodeResponse) ModifyUser(ctx *gin.Context, resp *ModifyUserResponse) error {
	return gonic.EncodeResponse(ctx, ctx.Writer, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder userGonicEncodeResponse) UpdateUser(ctx *gin.Context, resp *UpdateUserResponse) error {
	return gonic.EncodeResponse(ctx, ctx.Writer, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder userGonicEncodeResponse) GetUser(ctx *gin.Context, resp *GetUserResponse) error {
	return gonic.EncodeResponse(ctx, ctx.Writer, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder userGonicEncodeResponse) ListUser(ctx *gin.Context, resp *ListUserResponse) error {
	return gonic.EncodeResponse(ctx, ctx.Writer, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
