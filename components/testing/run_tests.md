# Инструкции по запуску тестов

## Структура тестов

В директории `components/testing/` созданы следующие тестовые файлы:

- `user_test.go` - тесты для модели пользователя
- `redis_test.go` - тесты для Redis клиента
- `config_test.go` - тесты для конфигурации
- `keyboards_test.go` - тесты для клавиатур
- `message_test.go` - тесты для констант сообщений
- `integration_test.go` - интеграционные тесты

## Запуск тестов

### Запуск всех тестов в модуле testing
```bash
cd components/testing
go test
```

### Запуск с подробным выводом
```bash
go test -v
```

### Запуск конкретного теста
```bash
# Тесты пользователя
go test -run TestNewUser

# Тесты Redis
go test -run TestRedisClient

# Интеграционные тесты
go test -run TestUserCreationAndStorage
```

### Запуск benchmark тестов
```bash
go test -bench=.
```

### Запуск с покрытием кода
```bash
go test -cover
```

### Генерация отчета о покрытии
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Требования для тестов

### Redis
Для запуска тестов Redis необходимо:
1. Установить Redis сервер
2. Запустить Redis на localhost:6379
3. Или настроить переменные окружения в .env файле

### Переменные окружения
Создайте файл `.env` в корне проекта:
```env
REDIS-ADDR=localhost:6379
REDIS-PASSWORD=
APIURL=your_api_url
APIKEY=your_api_key
TELEGRAMBOTTOKEN=your_bot_token
```

## Типы тестов

### Unit тесты
- `user_test.go` - тестирует логику создания и валидации пользователей
- `config_test.go` - тестирует загрузку конфигурации
- `keyboards_test.go` - тестирует создание клавиатур
- `message_test.go` - тестирует константы сообщений

### Integration тесты
- `redis_test.go` - тестирует взаимодействие с Redis
- `integration_test.go` - тестирует взаимодействие между модулями

### Benchmark тесты
- Измеряют производительность функций
- Запускаются с флагом `-bench=.`

## Интерпретация результатов

### Успешные тесты
```
PASS
ok      bot/components/testing   0.123s
```

### Неудачные тесты
```
--- FAIL: TestNewUser_InvalidInput (0.001s)
    user_test.go:45: expected error for invalid input, got nil
FAIL
exit status 1
```

### Покрытие кода
```
PASS
coverage: 85.7% of statements
ok      bot/components/testing   0.234s
```

## Отладка тестов

### Запуск с выводом логов
```bash
go test -v -run TestName
```

### Пропуск медленных тестов
```bash
go test -short
```

### Параллельное выполнение
```bash
go test -parallel 4
```

## Лучшие практики

1. **Именование**: `TestFunctionName_Scenario`
2. **Один тест - одна проверка**
3. **Table driven tests** для множественных сценариев
4. **Тестирование граничных случаев**
5. **Тестирование ошибок**
6. **Использование моков для внешних зависимостей**

## Добавление новых тестов

1. Создайте новый файл `new_module_test.go`
2. Импортируйте необходимые пакеты
3. Создайте функции с префиксом `Test`
4. Используйте `t.Error()` для ошибок и `t.Fatal()` для критических ошибок
5. Добавьте комментарии к сложным тестам 