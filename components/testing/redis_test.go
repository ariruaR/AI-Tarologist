package testing

import (
	"bot/components/redis"
	"context"
	"testing"
	"time"
)

func TestRedisClient_NewClient(t *testing.T) {
	client := redis.NewClient()

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
}

func TestRedisClient_SetterAndGetter(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Тест установки и получения значения
	key := "test_key"
	value := "test_value"
	expiration := time.Hour

	err := client.Setter(ctx, key, value, expiration)
	if err != nil {
		t.Fatalf("Setter() returned error: %v", err)
	}

	// Получаем значение
	result, err := client.Getter(ctx, key)
	if err != nil {
		t.Fatalf("Getter() returned error: %v", err)
	}

	if result != value {
		t.Errorf("expected value '%s', got '%s'", value, result)
	}
}

func TestRedisClient_SetNewUser(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Тест создания нового пользователя
	err := client.SetNewUser(ctx, 1, "testuser", false, "test info", nil, time.Hour)

	if err != nil {
		t.Fatalf("SetNewUser() returned error: %v", err)
	}
}

func TestRedisClient_SetNewUser_InvalidInput(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Тест с невалидными данными (должен вызвать panic)
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for invalid user data")
		}
	}()

	// Попытка создать пользователя с невалидными данными
	client.SetNewUser(ctx, 0, "", false, "", nil, time.Hour)
}

func TestRedisClient_UpdateFieldUser(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Сначала создаем пользователя
	err := client.SetNewUser(ctx, 1, "testuser", false, "test info", nil, time.Hour)
	if err != nil {
		t.Fatalf("SetNewUser() returned error: %v", err)
	}

	// Обновляем поле username
	err = client.UpdateFieldUser(ctx, 1, "username", "newusername", time.Hour)

	if err != nil {
		t.Fatalf("UpdateFieldUser() returned error: %v", err)
	}
}

func TestRedisClient_UpdateFieldUser_InvalidField(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Сначала создаем пользователя
	err := client.SetNewUser(ctx, 1, "testuser", false, "test info", nil, time.Hour)
	if err != nil {
		t.Fatalf("SetNewUser() returned error: %v", err)
	}

	// Тест обновления несуществующего поля
	err = client.UpdateFieldUser(ctx, 1, "nonexistent", "value", time.Hour)

	if err == nil {
		t.Error("expected error for invalid field")
	}
}

func TestRedisClient_UpdateFieldUser_InvalidType(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Сначала создаем пользователя
	err := client.SetNewUser(ctx, 1, "testuser", false, "test info", nil, time.Hour)
	if err != nil {
		t.Fatalf("SetNewUser() returned error: %v", err)
	}

	// Тест обновления поля с неправильным типом
	err = client.UpdateFieldUser(ctx, 1, "username", 123, time.Hour)

	if err == nil {
		t.Error("expected error for invalid type")
	}
}

func TestRedisClient_ReadUser(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Сначала создаем пользователя
	err := client.SetNewUser(ctx, 1, "testuser", false, "test info", nil, time.Hour)
	if err != nil {
		t.Fatalf("SetNewUser() returned error: %v", err)
	}

	// Читаем пользователя
	user, err := client.ReadUser(ctx, 1)
	if err != nil {
		t.Fatalf("ReadUser() returned error: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected ID 1, got %d", user.ID)
	}

	if user.Username != "testuser" {
		t.Errorf("expected username 'testuser', got '%s'", user.Username)
	}
}

func TestRedisClient_ReadUser_NotFound(t *testing.T) {
	client := redis.NewClient()
	ctx := context.Background()

	// Попытка прочитать несуществующего пользователя
	_, err := client.ReadUser(ctx, 999)

	if err == nil {
		t.Error("expected error for non-existent user")
	}
}

// Benchmark тесты
func BenchmarkRedisClient_Setter(b *testing.B) {
	client := redis.NewClient()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "benchmark_key_" + string(rune(i))
		value := "benchmark_value_" + string(rune(i))
		client.Setter(ctx, key, value, time.Hour)
	}
}

func BenchmarkRedisClient_Getter(b *testing.B) {
	client := redis.NewClient()
	ctx := context.Background()

	// Подготовка данных
	key := "benchmark_get_key"
	value := "benchmark_get_value"
	client.Setter(ctx, key, value, time.Hour)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.Getter(ctx, key)
	}
}

func BenchmarkRedisClient_SetNewUser(b *testing.B) {
	client := redis.NewClient()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.SetNewUser(ctx, i, "testuser", false, "test info", nil, time.Hour)
	}
}
