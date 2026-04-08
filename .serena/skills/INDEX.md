# Skills Index — Algorithms Learning Project

> Читать этот файл в начале каждого чата для восстановления контекста.

## Быстрые ссылки

| Что нужно | Куда идти |
|-----------|-----------|
| Методика обучения (Socratic, TDD, фазы) | `algorithms-mentor/SKILL.md` |
| Статус карточек (прогресс, следующая) | `cards/references/status.md` |
| Методика ответа на карточку | `cards/SKILL.md` + `cards/assets/response-template.md` |
| Шаблон выдачи задачи | `algorithms-mentor/assets/task-template.md` |
| Правила git | `workflow/git.md` |
| Алгоритмические паттерны (GCD и др.) | `algorithms-mentor/references/patterns.md` |
| Примеры из проектов | `projects/` |

## Текущий статус обучения

**Ученик**: Анна (Go backend-разработчик)  
**Цель**: Стажировка в российский бигтех  
**Срок**: 6 месяцев (жёсткий дедлайн)

### Фазы обучения
- Фаза 0: Математический фундамент (недели 1-2) ← **Текущая**
- Фаза 1: Массивы + Рекурсия (недели 3-7)
- Фаза 2: Линейные структуры (недели 8-12)
- Фаза 3: Деревья и Графы (недели 13-18)
- Фаза 4: ДП базовое (недели 19-20)
- Фаза 5: Полировка (недели 21-24)

### Карточки
- **Всего**: 54
- **Переработано**: 10 (Goroutines, Channels, Select, Mutex, WaitGroup, Context, Graceful Shutdown, errors.Is/errors.As, Dependency Injection, Generics)
- **Следующая**: #11 Clean Architecture

### Решённые задачи (неделя)
- [x] LeetCode 504: Base 7
- [x] LeetCode 405: Hexadecimal  
- [ ] LeetCode 12: Integer to Roman (запланировано на 2026-03-24)

## Критические правила (всегда помнить!)

1. **НИКОГДА не давать готовый код** — только наводящие вопросы
2. **TDD обязателен** — Red → Green → Refactor
3. **Early returns** — проверять граничные случаи первым делом
4. **Git** — я создаю сообщение, пользователь выполняет commit
5. **Карточки** — разговорный стиль, примеры из Task Manager/RAG, никакого кода в тексте

## Если пользователь просит...

| Запрос | Действие |
|--------|----------|
| "Создай комментарий для git" | `workflow/git.md` → создать сообщение, НЕ выполнять commit |
| "Проверь карточку" | `cards/SKILL.md` → проверить структуру ответа |
| "Помоги с карточкой" | `cards/SKILL.md` + наводящие вопросы, НЕ писать за неё |
| "Выдай задачу" | `algorithms-mentor/SKILL.md` + `assets/task-template.md` |
| "Сжатие / compress" | `workflow/compress.md` → создать сводку в `.session_summary` |
| "Прочитай сводку" | `workflow/compress.md` → прочитать `.session_summary` |
| "Запиши в journal" | `workflow/session.md` → обновить journal.csv + `.session_start` |
| "Начинаем сессию" | `workflow/session.md` → записать время в `.session_start` |
| "Перерыв" | `workflow/session.md` → BREAK_START в `.session_start` |
| "Продолжаем" | `workflow/session.md` → BREAK_END в `.session_start` |

## Gotchas (специфичные ошибки этого проекта)

### Алгоритмы
- ❌ Не использовать `math.Pow` для целых степеней двойки (использовать `1 << n`)
- ❌ Не забывать про two's complement для отрицательных чисел (uint32)
- ✅ Early returns для пустых входов, нулей, единичных случаев

### Go
- ❌ `len("привет")` = 12 байт, не 6 символов — использовать `[]rune`
- ❌ Map iteration order случайный — использовать slice для порядка
- ✅ Встроенные `min()`, `max()` доступны с Go 1.21+
- ✅ Nil checks в связных списках и деревьях

### Общение
- ❌ Никаких "ресторанных" аналогий
- ❌ Никакого кода в ответе на карточку
- ✅ Примеры из Task Manager / RAG-сервиса
- ✅ Разговорный стиль, связки: "Проще говоря...", "То есть..."

## Последнее обновление

2026-03-23 — создана новая структура skills
