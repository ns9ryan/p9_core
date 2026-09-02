package errorhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/ns9ryan/p9_core/api/internal/types"
	"github.com/ns9ryan/p9_core/pkg/i18n"
	"github.com/ns9ryan/p9_core/pkg/i18nkey"
	"github.com/ns9ryan/p9_core/pkg/validate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler 错误处理器
type Handler struct {
	trans *i18n.Translator // 翻译器
}

// New 创建错误处理器
func New(trans *i18n.Translator) *Handler {
	return &Handler{
		trans: trans,
	}
}

// Handle 处理 API 错误
func (h *Handler) Handle(ctx context.Context, err error) (int, any) {
	// 处理参数校验错误
	var validationError *validate.Error
	if errors.As(err, &validationError) {
		return h.response(http.StatusBadRequest, validationError.Error())
	}

	// 处理 gRPC 错误
	if grpcStatus, ok := status.FromError(err); ok {
		httpStatus := grpcCodeToHTTPStatus(grpcStatus.Code())
		message := h.grpcMessage(ctx, grpcStatus)

		return h.response(httpStatus, message)
	}

	// 其他请求解析错误统一按参数错误处理
	return h.response(
		http.StatusBadRequest,
		h.trans.Trans(ctx, i18nkey.InvalidRequest),
	)
}

// grpcMessage 获取 gRPC 错误提示
func (h *Handler) grpcMessage(ctx context.Context, grpcStatus *status.Status) string {
	switch grpcStatus.Code() {
	case codes.ResourceExhausted:
		return h.trans.Trans(ctx, i18nkey.TooManyRequests)

	case codes.Unavailable:
		return h.trans.Trans(ctx, i18nkey.ServiceUnavailable)

	case codes.DeadlineExceeded:
		return h.trans.Trans(ctx, i18nkey.RequestTimeout)

	default:
		return h.trans.Trans(ctx, grpcStatus.Message())
	}
}

// response 创建错误响应
func (h *Handler) response(code int, message string) (int, any) {
	return code, &types.BaseResponse{
		Code: code,
		Msg:  message,
	}
}

// grpcCodeToHTTPStatus 将 gRPC 状态码转换为 HTTP 状态码
func grpcCodeToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.InvalidArgument,
		codes.AlreadyExists,
		codes.FailedPrecondition,
		codes.OutOfRange:
		return http.StatusBadRequest

	case codes.Unauthenticated:
		return http.StatusUnauthorized

	case codes.PermissionDenied:
		return http.StatusForbidden

	case codes.NotFound:
		return http.StatusNotFound

	case codes.ResourceExhausted:
		return http.StatusTooManyRequests

	case codes.Unimplemented:
		return http.StatusNotImplemented

	case codes.Unavailable:
		return http.StatusServiceUnavailable

	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout

	default:
		return http.StatusInternalServerError
	}
}
