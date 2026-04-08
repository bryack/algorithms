# Проект: RAG-сервис (HTTP-драйвер для LLM)

## Описание

Сервис для работы с LLM (Large Language Model) через HTTP API. Позволяет отправлять запросы к AI-моделям и получать структурированные ответы.

## Архитектура

- **Язык**: Go
- **Протокол**: HTTP/REST
- **Формат**: JSON
- **Особенности**: Generics для типизации ответов

## Примеры использования в карточках

### Generics
> "В RAG-сервисе использовал generics для HTTP-драйвера. Функция `post[T](url string, body any) (T, error)` принимает любой тип ответа и десериализует JSON прямо в него. Для метрик — `post[Metrics]`, для статуса — `post[Status]`. Это позволило не дублировать код для каждого типа."

### Context (timeout)
> "Для таймаута запросов к LLM использовал context.WithTimeout. Если модель не отвечала за 30 секунд — контекст отменялся, и клиент получал ошибку, а не ждал вечно."

### errors.Is / errors.As
> "При обработке ошибок HTTP использовал errors.As для извлечения конкретного типа ошибки. Например, отдельная логика для сетевых ошибок и для ошибок API (4xx, 5xx)."

### HTTP Client
> "Использовал http.Client с таймаутом и retry-логикой. Для rate limiting от LLM-провайдера добавил exponential backoff."

## Технические детали (для примеров)

- Универсальный HTTP-драйвер с generics
- Таймауты через context
- Retry с exponential backoff
- Обработка ошибок API (4xx/5xx)
- JSON marshaling/unmarshaling
- Структурированное логирование

## Функция post[T] (для справки)

```go
// Пример структуры, не для включения в ответ
func post[T any](url string, body any) (T, error) {
    var result T
    // ... HTTP запрос
    json.Unmarshal(respBody, &result)
    return result, nil
}
```

**Использование:**
- `post[Metrics]("/api/metrics", req)` → возвращает `Metrics`
- `post[Status]("/api/status", req)` → возвращает `Status`
