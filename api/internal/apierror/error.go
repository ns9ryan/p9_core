package apierror

import (
	"errors"
	"net/http"

	"github.com/ns9ryan/p9_core/pkg/i18nkey"
)

// Error API内部错误
type Error struct {
	HTTPStatus int    // HTTP状态码
	MessageKey string // 多语言消息Key
	err        error  // 原始错误
}

// Internal 创建服务器内部错误
func Internal(err error) error {
	if err == nil {
		return nil
	}

	return &Error{
		HTTPStatus: http.StatusInternalServerError,
		MessageKey: i18nkey.InternalError,
		err:        err,
	}
}

// Error 返回原始错误信息
func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return e.MessageKey
}

// Unwrap 返回原始错误
func (e *Error) Unwrap() error {
	return e.err
}

// FromError 获取API内部错误
func FromError(err error) (*Error, bool) {
	var apiErr *Error
	if !errors.As(err, &apiErr) {
		return nil, false
	}

	return apiErr, true
}
