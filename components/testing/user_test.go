package testing

import (
	"bot/components/models"
	"strings"
	"testing"
)

func TestNewUser_ValidInput(t *testing.T) {
	user, err := models.NewUser(1, "testuser", false, "test info", nil)

	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected ID 1, got %d", user.ID)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got '%s'", user.Username)
	}

	if user.IsPremium != false {
		t.Errorf("expected IsPremium false, got %v", user.IsPremium)
	}

	if user.Info != "test info" {
		t.Errorf("expected Info 'test info', got '%s'", user.Info)
	}
}

func TestNewUser_WithPaymentHistory(t *testing.T) {
	paymentHistory := map[string]string{
		"2024-01-01": "payment_123",
		"2024-01-02": "payment_456",
	}

	user, err := models.NewUser(1, "testuser", true, "test info", paymentHistory)

	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	if len(user.LastPayment) != 2 {
		t.Errorf("expected 2 payment history entries, got %d", len(user.LastPayment))
	}

	if user.LastPayment["2024-01-01"] != "payment_123" {
		t.Errorf("expected payment_123, got %s", user.LastPayment["2024-01-01"])
	}
}

func TestNewUser_TableDriven(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		username       string
		isPremium      bool
		info           string
		paymentHistory map[string]string
		wantErr        bool
		errorContains  string
	}{
		{
			name:           "valid user",
			id:             1,
			username:       "testuser",
			isPremium:      false,
			info:           "test info",
			paymentHistory: nil,
			wantErr:        false,
		},
		{
			name:           "invalid id zero",
			id:             0,
			username:       "testuser",
			isPremium:      false,
			info:           "test info",
			paymentHistory: nil,
			wantErr:        true,
			errorContains:  "id не может быть равен нулю",
		},
		{
			name:           "empty username",
			id:             1,
			username:       "",
			isPremium:      false,
			info:           "test info",
			paymentHistory: nil,
			wantErr:        true,
			errorContains:  "username не может быть пустым",
		},
		{
			name:           "username with spaces only",
			id:             1,
			username:       "   ",
			isPremium:      false,
			info:           "test info",
			paymentHistory: nil,
			wantErr:        true,
			errorContains:  "username не может быть пустым",
		},
		{
			name:           "username too long",
			id:             1,
			username:       "verylongusernameverylongusernameverylongusernameverylongusernameverylongusername",
			isPremium:      false,
			info:           "test info",
			paymentHistory: nil,
			wantErr:        true,
			errorContains:  "username не может быть длиннее 50 символов",
		},
		{
			name:           "info too long",
			id:             1,
			username:       "testuser",
			isPremium:      false,
			info:           string(make([]byte, 1001)), // 1001 символов
			paymentHistory: nil,
			wantErr:        true,
			errorContains:  "info не может быть длиннее 1000 символов",
		},
		{
			name:      "invalid payment history empty key",
			id:        1,
			username:  "testuser",
			isPremium: false,
			info:      "test info",
			paymentHistory: map[string]string{
				"": "payment_123",
			},
			wantErr:       true,
			errorContains: "ключи в paymentHistory не могут быть пустыми",
		},
		{
			name:      "invalid payment history empty value",
			id:        1,
			username:  "testuser",
			isPremium: false,
			info:      "test info",
			paymentHistory: map[string]string{
				"2024-01-01": "",
			},
			wantErr:       true,
			errorContains: "значения в paymentHistory не могут быть пустыми",
		},
		{
			name:      "valid payment history",
			id:        1,
			username:  "testuser",
			isPremium: true,
			info:      "test info",
			paymentHistory: map[string]string{
				"2024-01-01": "payment_123",
				"2024-01-02": "payment_456",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := models.NewUser(tt.id, tt.username, tt.isPremium, tt.info, tt.paymentHistory)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("error message '%s' does not contain '%s'", err.Error(), tt.errorContains)
				}
				return
			}

			// Проверяем, что пользователь создан корректно
			if user.ID != tt.id {
				t.Errorf("expected ID %d, got %d", tt.id, user.ID)
			}

			if user.Username != strings.TrimSpace(tt.username) {
				t.Errorf("expected username '%s', got '%s'", strings.TrimSpace(tt.username), user.Username)
			}

			if user.IsPremium != tt.isPremium {
				t.Errorf("expected IsPremium %v, got %v", tt.isPremium, user.IsPremium)
			}

			if user.Info != strings.TrimSpace(tt.info) {
				t.Errorf("expected Info '%s', got '%s'", strings.TrimSpace(tt.info), user.Info)
			}
		})
	}
}

func TestNewUser_Trimming(t *testing.T) {
	user, err := models.NewUser(1, "  testuser  ", false, "  test info  ", nil)

	if err != nil {
		t.Fatalf("NewUser() returned error: %v", err)
	}

	if user.Username != "testuser" {
		t.Errorf("expected trimmed username 'testuser', got '%s'", user.Username)
	}

	if user.Info != "test info" {
		t.Errorf("expected trimmed info 'test info', got '%s'", user.Info)
	}
}

func TestNewUser_NegativeID(t *testing.T) {
	// Тест на отрицательный ID (должен быть валидным)
	user, err := models.NewUser(-1, "testuser", false, "test info", nil)

	if err != nil {
		t.Fatalf("NewUser() returned error for negative ID: %v", err)
	}

	if user.ID != -1 {
		t.Errorf("expected ID -1, got %d", user.ID)
	}
}

// Вспомогательная функция для проверки содержимого строки
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
