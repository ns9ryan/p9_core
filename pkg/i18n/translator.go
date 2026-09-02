package i18n

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Translator 翻译器
type Translator struct {
	bundle          *goi18n.Bundle
	defaultLanguage string
}

// New 创建翻译器
func New(config Config, localeFS fs.FS) (*Translator, error) {
	defaultTag, err := language.Parse(config.DefaultLanguage)
	if err != nil {
		return nil, fmt.Errorf("无效的默认语言 %q: %w", config.DefaultLanguage, err)
	}

	bundle := goi18n.NewBundle(defaultTag)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	err = fs.WalkDir(localeFS, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || !strings.EqualFold(filepath.Ext(path), ".json") {
			return nil
		}

		if _, err := bundle.LoadMessageFileFS(localeFS, path); err != nil {
			return fmt.Errorf("加载语言文件 %q 失败: %w", path, err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Translator{
		bundle:          bundle,
		defaultLanguage: config.DefaultLanguage,
	}, nil
}

// Translate 翻译消息
func (t *Translator) Translate(language, key string) string {
	if language == "" {
		language = t.defaultLanguage
	}

	localizer := goi18n.NewLocalizer(
		t.bundle,
		language,
		t.defaultLanguage,
	)

	message, err := localizer.Localize(&goi18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil || message == "" {
		return key
	}

	return message
}

// Trans 根据上下文语言翻译消息
func (t *Translator) Trans(ctx context.Context, key string) string {
	return t.Translate(LanguageFromContext(ctx), key)
}
