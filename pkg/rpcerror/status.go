package rpcerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Internal 创建内部错误
func Internal(message string) error {
	return status.Error(codes.Internal, message)
}

// InvalidArgument 创建参数错误
func InvalidArgument(message string) error {
	return status.Error(codes.InvalidArgument, message)
}

// NotFound 创建资源不存在错误
func NotFound(message string) error {
	return status.Error(codes.NotFound, message)
}
