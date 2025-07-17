package testing

import (
	"bot/components/models"
	"bot/components/redis"
	"bot/config"
	"context"
	"testing"
	"time"
)

// TestUserCreationAndStorage тестирует полный цикл создания и сохранения пользователя
func TestUserCreationAndStorage(t *testing.T) {
	// Создаем пользователя
	user, err := models.NewUser(1, "testuser", false, "test info", nil)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Проверяем, что пользователь создан корректно
	if user.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", user.ID)
	}

	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}

	// Создаем Redis клиент
	client := redis.NewClient()
	if client == nil {
		t.Fatal("Failed to create Redis client")
	}

	// Сохраняем пользователя в Redis
	ctx := context.Background()
	err = client.SetNewUser(ctx, user.ID, user.Username, user.IsPremium, user.Info, user.LastPayment, time.Hour)
	if err != nil {
		t.Fatalf("Failed to save user to Redis: %v", err)
	}

	// Читаем пользователя из Redis
	savedUser, err := client.ReadUser(ctx, user.ID)
	if err != nil {
		t.Fatalf("Failed to read user from Redis: %v", err)
	}

	// Проверяем, что данные совпадают
	if savedUser.ID != user.ID {
		t.Errorf("Saved user ID mismatch: expected %d, got %d", user.ID, savedUser.ID)
	}

	if savedUser.Username != user.Username {
		t.Errorf("Saved user username mismatch: expected '%s', got '%s'", user.Username, savedUser.Username)
	}

	if savedUser.IsPremium != user.IsPremium {
		t.Errorf("Saved user IsPremium mismatch: expected %v, got %v", user.IsPremium, savedUser.IsPremium)
	}
}

// TestUserUpdateFlow тестирует процесс обновления пользователя
func TestUserUpdateFlow(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Создаем и сохраняем пользователя
	err := client.SetNewUser(ctx, 2, "olduser", false, "old info", nil, time.Hour)
	if err != nil {
		t.Fatalf("Failed to create initial user: %v", err)
	}

	// Обновляем поле пользователя
	err = client.UpdateFieldUser(ctx, 2, "username", "newuser", time.Hour)
	if err != nil {
		t.Fatalf("Failed to update user field: %v", err)
	}

	// Читаем обновленного пользователя
	updatedUser, err := client.ReadUser(ctx, 2)
	if err != nil {
		t.Fatalf("Failed to read updated user: %v", err)
	}

	// Проверяем, что поле обновилось
	if updatedUser.Username != "newuser" {
		t.Errorf("Expected updated username 'newuser', got '%s'", updatedUser.Username)
	}
}

// TestConfigurationAndRedis тестирует взаимодействие конфигурации и Redis
func TestConfigurationAndRedis(t *testing.T) {
	// Читаем конфигурацию
	cfg := config.Readconfig()

	// Проверяем, что конфигурация загружена
	if cfg.REDIS_ADDR == "" {
		t.Log("REDIS_ADDR is empty (this is normal if .env file is not loaded)")
	}

	// Создаем Redis клиент (использует конфигурацию)
	client := redis.NewClient()
	if client == nil {
		t.Fatal("Failed to create Redis client with configuration")
	}

	// Простой тест записи/чтения
	ctx := context.Background()
	testKey := "integration_test_key"
	testValue := "integration_test_value"

	err := client.Setter(ctx, testKey, testValue, time.Minute)
	if err != nil {
		t.Fatalf("Failed to set test value: %v", err)
	}

	result, err := client.Getter(ctx, testKey)
	if err != nil {
		t.Fatalf("Failed to get test value: %v", err)
	}

	if result != testValue {
		t.Errorf("Expected value '%s', got '%s'", testValue, result)
	}
}

// TestErrorHandling тестирует обработку ошибок
func TestErrorHandling(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Тест чтения несуществующего пользователя
	_, err := client.ReadUser(ctx, 99999)
	if err == nil {
		t.Error("Expected error when reading non-existent user")
	}

	// Тест обновления несуществующего пользователя
	err = client.UpdateFieldUser(ctx, 99999, "username", "newuser", time.Hour)
	if err == nil {
		t.Error("Expected error when updating non-existent user")
	}
}

// TestDataValidation тестирует валидацию данных через весь стек
func TestDataValidation(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Тест с невалидными данными (должен вызвать panic в SetNewUser)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid user data in SetNewUser")
		}
	}()

	// Попытка создать пользователя с невалидными данными
	client.SetNewUser(ctx, 0, "", false, "", nil, time.Hour)
}
