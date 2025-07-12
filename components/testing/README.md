# Руководство по тестированию в Go

## Основы тестирования в Go

### 1. Структура тестовых файлов

В Go тестовые файлы должны:
- Называться с суффиксом `_test.go`
- Находиться в том же пакете, что и тестируемый код
- Импортировать пакет `testing`

```go
// user_test.go
package models

import (
    "testing"
)
```

### 2. Функции тестирования

#### Базовые тесты
```go
func TestNewUser(t *testing.T) {
    // Тест создается здесь
}
```

#### Тесты с подтестами
```go
func TestNewUser_ValidInput(t *testing.T) {
    t.Run("valid user creation", func(t *testing.T) {
        // подтест
    })
    
    t.Run("invalid user creation", func(t *testing.T) {
        // подтест
    })
}
```

### 3. Основные методы тестирования

#### t.Error() и t.Errorf()
- Не останавливает выполнение теста
- Показывает ошибку и продолжает

#### t.Fatal() и t.Fatalf()
- Останавливает выполнение теста
- Используется для критических ошибок

#### t.Log() и t.Logf()
- Выводит информацию во время теста
- Полезно для отладки

### 4. Assertions (проверки)

#### Проверка равенства
```go
if got != want {
    t.Errorf("got %v, want %v", got, want)
}
```

#### Проверка ошибок
```go
if err == nil {
    t.Error("expected error, got nil")
}
```

#### Проверка на nil
```go
if result == nil {
    t.Error("expected non-nil result")
}
```

### 5. Table Driven Tests

```go
func TestNewUser_TableDriven(t *testing.T) {
    tests := []struct {
        name    string
        id      int
        username string
        isPremium bool
        info    string
        wantErr bool
    }{
        {
            name: "valid user",
            id: 1,
            username: "testuser",
            isPremium: false,
            info: "test info",
            wantErr: false,
        },
        {
            name: "invalid id",
            id: 0,
            username: "testuser",
            isPremium: false,
            info: "test info",
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := NewUser(tt.id, tt.username, tt.isPremium, tt.info, nil)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### 6. Benchmark тесты

```go
func BenchmarkNewUser(b *testing.B) {
    for i := 0; i < b.N; i++ {
        NewUser(1, "testuser", false, "test info", nil)
    }
}
```

### 7. Примеры тестов

#### Тест функции NewUser
```go
func TestNewUser_ValidInput(t *testing.T) {
    user, err := NewUser(1, "testuser", false, "test info", nil)
    
    if err != nil {
        t.Fatalf("NewUser() returned error: %v", err)
    }
    
    if user.ID != 1 {
        t.Errorf("expected ID 1, got %d", user.ID)
    }
    
    if user.Username != "testuser" {
        t.Errorf("expected username 'testuser', got '%s'", user.Username)
    }
}
```

#### Тест валидации
```go
func TestNewUser_InvalidInput(t *testing.T) {
    _, err := NewUser(0, "", false, "", nil)
    
    if err == nil {
        t.Error("expected error for invalid input, got nil")
    }
}
```

### 8. Запуск тестов

```bash
# Запуск всех тестов в пакете
go test

# Запуск с подробным выводом
go test -v

# Запуск конкретного теста
go test -run TestNewUser

# Запуск benchmark тестов
go test -bench=.

# Запуск с покрытием кода
go test -cover

# Генерация отчета о покрытии
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### 9. Лучшие практики

1. **Именование тестов**: `TestFunctionName_Scenario`
2. **Один тест - одна проверка**: Каждый тест должен проверять одну вещь
3. **Используйте table driven tests**: Для множественных сценариев
4. **Тестируйте граничные случаи**: Пустые значения, максимальные значения
5. **Тестируйте ошибки**: Убедитесь, что функции возвращают ошибки когда нужно
6. **Используйте моки**: Для внешних зависимостей

### 10. Моки и интерфейсы

```go
type UserRepository interface {
    Save(user User) error
    Get(id int) (User, error)
}

type MockUserRepository struct {
    users map[int]User
}

func (m *MockUserRepository) Save(user User) error {
    m.users[user.ID] = user
    return nil
}

func (m *MockUserRepository) Get(id int) (User, error) {
    user, exists := m.users[id]
    if !exists {
        return User{}, errors.New("user not found")
    }
    return user, nil
}
```

### 11. Тестирование с контекстом

```go
func TestRedisClient_UpdateUser(t *testing.T) {
    ctx := context.Background()
    client := NewClient()
    
    // Setup
    user, _ := NewUser(1, "testuser", false, "test info", nil)
    
    // Test
    err := client.UpdateUser(ctx, 1, "newuser", true, "new info", nil, time.Hour)
    
    if err != nil {
        t.Errorf("UpdateUser() error = %v", err)
    }
}
```

Это базовое руководство по тестированию в Go. Практикуйтесь и изучайте больше примеров! 