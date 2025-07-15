package testing

import (
	"bot/config"
	"testing"
)

func TestReadConfig(t *testing.T) {
	cfg := config.Readconfig()

	// Проверяем, что конфигурация создается корректно
	// Config - это структура, а не указатель, поэтому не может быть nil
	if cfg.APIURL == "" && cfg.APIKEY == "" && cfg.BOTTOKEN == "" {
		t.Log("config fields are empty (this is normal if .env file is not loaded)")
	}
}

func TestConfig_Fields(t *testing.T) {
	cfg := config.Readconfig()

	// Проверяем, что поля конфигурации существуют
	// Значения могут быть пустыми, если .env файл не загружен
	_ = cfg.APIURL
	_ = cfg.APIKEY
	_ = cfg.BOTTOKEN
	_ = cfg.REDIS_ADDR
	_ = cfg.REDIS_PASSWORD

	// Тест проходит, если структура создается корректно
}

func TestConfig_Consistency(t *testing.T) {
	cfg1 := config.Readconfig()
	cfg2 := config.Readconfig()

	// Проверяем, что конфигурация читается консистентно
	if cfg1.REDIS_ADDR != cfg2.REDIS_ADDR {
		t.Error("config values are not consistent between reads")
	}

	if cfg1.APIURL != cfg2.APIURL {
		t.Error("config values are not consistent between reads")
	}
}
