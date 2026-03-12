# selectelLinter

Линтер для Go, совместимый с `golangci-lint`, который проверяет сообщения логов на стиль и безопасность. Реализован как `go/analysis` анализатор и может запускаться как:
1. отдельный бинарник (singlechecker),
2. плагин для `golangci-lint`.

## Поддерживаемые логгеры
- `log/slog`
- `go.uber.org/zap`

Проверяются вызовы:
- `slog.Debug/Info/Warn/Error`
- методы `Logger` из `log/slog`
- методы `Logger` и `SugaredLogger` из `go.uber.org/zap`

## Правила
1. Сообщение начинается со строчной буквы.
2. Сообщение только на английском (только латинские буквы).
3. Сообщение без спецсимволов/эмодзи (разрешены буквы, цифры, пробелы).
4. Сообщение не содержит чувствительных данных (по ключевым словам в именах переменных).

## Как работает анализ
Линтер анализирует первый аргумент лог-вызова и обрабатывает только «статические» сообщения:
- строковый литерал,
- идентификатор,
- конкатенация строк (`"text " + var + " more"`).

Если сообщение невозможно восстановить статически, проверка пропускается.

Для правила чувствительных данных используется поиск по именам переменных в конкатенации. Ключевые слова:
`password`, `passwd`, `pwd`, `key`, `api_key`, `apikey`, `token`, `secret`, `access_key`, `private_key`.

## Установка и сборка
Требования: Go 1.22+ (в `go.mod` указан `go 1.25.0`).

Сборка бинарника:
```bash
go build -o bin/selectelLinter ./cmd/selectelLinter
```

Запуск как singlechecker:
```bash
./bin/selectelLinter ./...
```

Тесты:
```bash
go test ./...
```

## Плагин для golangci-lint
*Сборка плагина:
```bash
go build -buildmode=plugin -o loglint.so ./plugin
```

*флаг `-buildmode=plugin` не поддерживается в Windows

Пример `.golangci.yml`:
```yaml
linters:
  enable:
    - loglint
  settings:
    custom:
      loglint:
        path: ./loglint.so
        description: Selectel log linter

run:
  skip-dirs-use-default: false

version: 2
```

Запуск:
```bash
golangci-lint run ./...
```

## Примеры
Неправильно:
```go
slog.Info("Starting server on port 8080")
slog.Info("запуск сервера")
slog.Info("server started! 🚀")

password := "123"
slog.Info("user password " + password)
```

Правильно:
```go
slog.Info("starting server on port 8080")
slog.Info("server started")
slog.Info("user authenticated successfully")
```

## Ограничения
- Плагин `-buildmode=plugin` не поддерживается на Windows.
- Правило «English only» допускает только латинские буквы; любые кириллические/другие буквы будут ошибкой.
- Правило «no special chars» запрещает любую пунктуацию.
- Чувствительные данные детектируются только по именам переменных (возможны ложные срабатывания, например `monkey` содержит `key`).

## Структура проекта
- `analyzer/` — основной анализатор
- `rules/` — правила
- `loggers/` — определение поддерживаемых логгеров
- `plugin/` — обертка для golangci-lint
- `cmd/selectelLinter/` — запуск в виде singlechecker
- `testdata/` — тестовые кейсы для `analysistest`
