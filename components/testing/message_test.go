package testing

import (
	"bot/components/message"
	"strings"
	"testing"
)

func TestMessageConstants_NotEmpty(t *testing.T) {
	// Проверяем, что константы сообщений не пустые
	if strings.TrimSpace(message.StartText) == "" {
		t.Error("StartText is empty")
	}

	if strings.TrimSpace(message.BuyText) == "" {
		t.Error("BuyText is empty")
	}

	if strings.TrimSpace(message.StarRequest) == "" {
		t.Error("StarRequest is empty")
	}

	if strings.TrimSpace(message.NotalMap) == "" {
		t.Error("NotalMap is empty")
	}

	if strings.TrimSpace(message.TaroAdvice) == "" {
		t.Error("TaroAdvice is empty")
	}
}

func TestMessageConstants_ContainsPlaceholders(t *testing.T) {
	// Проверяем, что сообщения содержат плейсхолдеры для подстановки данных
	if !strings.Contains(message.StarRequest, "%s") {
		t.Error("StarRequest should contain placeholder")
	}

	if !strings.Contains(message.NotalMap, "%s") {
		t.Error("NotalMap should contain placeholder")
	}

	if !strings.Contains(message.TaroAdvice, "%s") {
		t.Error("TaroAdvice should contain placeholder")
	}
}

func TestMessageConstants_Content(t *testing.T) {
	// Проверяем содержание сообщений
	if !strings.Contains(message.StartText, "ИИ-Таролог") {
		t.Error("StartText should mention ИИ-Таролог")
	}

	if !strings.Contains(message.BuyText, "стоимости") {
		t.Error("BuyText should mention стоимости")
	}

	if !strings.Contains(message.StarRequest, "таролог") {
		t.Error("StarRequest should mention таролог")
	}

	if !strings.Contains(message.NotalMap, "натальной карты") {
		t.Error("NotalMap should mention натальной карты")
	}

	if !strings.Contains(message.TaroAdvice, "Таро") {
		t.Error("TaroAdvice should mention Таро")
	}
}

func TestMessageConstants_Length(t *testing.T) {
	// Проверяем, что сообщения имеют разумную длину
	if len(message.StartText) < 50 {
		t.Error("StartText seems too short")
	}

	if len(message.BuyText) < 50 {
		t.Error("BuyText seems too short")
	}

	if len(message.StarRequest) < 100 {
		t.Error("StarRequest seems too short")
	}

	if len(message.NotalMap) < 100 {
		t.Error("NotalMap seems too short")
	}

	if len(message.TaroAdvice) < 100 {
		t.Error("TaroAdvice seems too short")
	}
}

func TestMessageConstants_Formatting(t *testing.T) {
	// Проверяем форматирование сообщений
	if !strings.Contains(message.StartText, "\n") {
		t.Error("StartText should contain line breaks")
	}

	if !strings.Contains(message.BuyText, "\n") {
		t.Error("BuyText should contain line breaks")
	}

	if !strings.Contains(message.StarRequest, "\n") {
		t.Error("StarRequest should contain line breaks")
	}

	if !strings.Contains(message.NotalMap, "\n") {
		t.Error("NotalMap should contain line breaks")
	}

	if !strings.Contains(message.TaroAdvice, "\n") {
		t.Error("TaroAdvice should contain line breaks")
	}
}
