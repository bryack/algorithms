# NeetCode Roadmap — Stack

> **Сгенерировано командой `rtk go run ./cmd/sync`.**
> Не редактируй вручную — правь catalog.yaml (статика) и заметки Obsidian (прогресс).

---

## Легенда оценки (Retention)

| Метка | Значение | Когда ставить |
|-------|----------|---------------|
| 🟢 **L4** | Решила сама, быстро, чистый код | Без подсказок, за разумное время, сразу верный Big O |
| 🟡 **L3** | Решила сама, но медленно или с багами | Сама, но долго думала или баги в edge cases |
| 🟠 **L2** | Нужны были подсказки по алгоритму | Я направил к паттерну, но код написала сама |
| 🔴 **L1** | Существенная помощь, нужно повторить | Помогал с алгоритмом и/или кодом |
| ⚪ **L5** | Эксперт — могу объяснить другому | Решила, оптимизировала, знаю все edge cases |

- **Статус**: `[x]` — решена (есть заметка в Obsidian), `[ ]` — ещё не решена.
- **needRepeat**: `true` — в Obsidian стоит флаг «нужно повторить».
- **Решений**: количество дат в `solved_dates` (сколько раз решалась).

---

## Правило перехода

> **Не начинать новую задачу, пока текущая не доведена до L4.**
>
> Исключение: разминка (warm-up) перед новой задачей — повторение уже решённой задачи.

---

## Ежедневная разминка (Repetition)

0. **Warm-up**: Перед новой задачей решаем одну из уже выполненных самостоятельно.
   - **Вербализация**: Сначала описываешь алгоритм и код словами (почему и как).
   - **Zero Hints**: Без подсказок во время решения разминки.
   - **Без тестов**: Моделируем ситуацию на собеседовании.

---

## Как работаем

1. **TDD**: Сначала тесты в `day_X/main_test.go`.
2. **Go Way**: Используем `make` с аллокацией, `struct{}` для сетов, `range` для итераций.
3. **Complexity**: Анализируем Time и Space Complexity для каждого решения.

---
## Основной блок

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **20. Valid Parentheses** | Easy | Stack (bracket matching) | 2026-04-26 | 🟡 L3 | true | 11 |
| [x] | **22. Generate Parentheses** | Medium | Stack / Backtracking | 2026-07-19 | 🔴 L1 |  | 3 |
| [x] | **71. Simplify Path** | Medium | Stack (path components) | 2026-07-18 | 🟢 L4 |  | 6 |
| [ ] | **84. Largest Rectangle In Histogram** | Hard | Monotonic Stack |  |  |  |  |
| [x] | **150. Evaluate Reverse Polish Notation** | Medium | Stack (postfix evaluation) | 2026-07-23 | 🟢 L4 |  | 2 |
| [x] | **155. Min Stack** | Medium | Stack with Min Tracking | 2026-07-01 | 🟡 L3 |  | 13 |
| [x] | **224. Basic Calculator** | Hard | Stack (expression evaluation with parentheses, sign handling) | 2026-07-23 | 🟡 L3 | true | 5 |
| [x] | **225. Implement Stack Using Queues** | Easy | Queue-based Stack Design | 2026-07-25 | ⚪ L5 |  | 1 |
| [x] | **232. Implement Queue using Stacks** | Easy | Stack-based Queue Design | 2026-07-25 | 🟢 L4 | true | 3 |
| [x] | **394. Decode String** | Medium | Stack (nested decoding) | 2026-07-26 | 🟠 L2 | true | 3 |
| [x] | **682. Baseball Game** | Easy | Stack (score ops) | 2026-07-18 | 🟢 L4 |  | 5 |
| [x] | **735. Asteroid Collision** | Medium | Stack (collision simulation) | 2026-07-19 | 🟡 L3 | true | 8 |
| [x] | **739. Daily Temperatures** | Medium | Monotonic Stack | 2026-07-26 | 🔴 L1 | true | 2 |
| [ ] | **853. Car Fleet** | Medium | Stack / Sorting |  |  |  |  |
| [ ] | **895. Maximum Frequency Stack** | Hard | Stack + Frequency Map |  |  |  |  |
| [ ] | **901. Online Stock Span** | Medium | Monotonic Stack |  |  |  |  |
| [x] | **1047. Remove All Adjacent Duplicates In String** | Easy | Stack (char filtering) | 2026-07-18 | 🟢 L4 |  | 5 |
| [x] | **2390. Removing Stars From a String** | Medium | Stack (char filtering) | 2026-07-18 | 🟢 L4 |  | 5 |

