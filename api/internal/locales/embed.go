package locales

import "embed"

// FS 保存 API 语言资源
//
//go:embed locale/*.json
var FS embed.FS
