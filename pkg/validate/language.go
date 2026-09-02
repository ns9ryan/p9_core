package validate

import "strings"

const (
	langZh = "zh"
	langEn = "en"
)

var supportedLanguages = map[string]string{
	"zh":    langZh,
	"zh-cn": langZh,
	"en":    langEn,
	"en-us": langEn,
}

// resolveLanguage 获取支持的校验语言
func resolveLanguage(value, fallback string) string {
	if lang, ok := lookupLanguage(value); ok {
		return lang
	}

	if lang, ok := lookupLanguage(fallback); ok {
		return lang
	}

	return langZh
}

// lookupLanguage 查找支持的校验语言
func lookupLanguage(value string) (string, bool) {
	value = strings.ToLower(strings.TrimSpace(value))
	lang, ok := supportedLanguages[value]
	return lang, ok
}
