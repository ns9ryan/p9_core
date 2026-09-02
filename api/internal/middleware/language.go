package middleware

import (
	"net/http"

	"github.com/ns9ryan/p9_core/pkg/i18n"
)

// LanguageMiddleware 语言中间件
type LanguageMiddleware struct{}

// NewLanguageMiddleware 创建语言中间件
func NewLanguageMiddleware() *LanguageMiddleware {
	return &LanguageMiddleware{}
}

// Handle 处理请求语言
func (m *LanguageMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取请求语言并写入上下文
		language := r.Header.Get("Accept-Language")
		ctx := i18n.WithLanguage(r.Context(), language)

		next(w, r.WithContext(ctx))
	}
}
