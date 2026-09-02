package rpcerror

import (
	"context"
	"strings"

	"google.golang.org/grpc"
)

// UnaryClientInterceptor 记录 RPC 调用来源
func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req any,
	reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err == nil {
		return nil
	}

	return Wrap(strings.TrimPrefix(method, "/"), err)
}
