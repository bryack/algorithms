# NeetCode Roadmap — Trees

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
| [x] | **50. Pow(x, n)** | Medium | recursion, math | 2026-08-02 | 🟢 L4 |  | 9 |
| [x] | **98. Validate Binary Search Tree** | Medium | recursion | 2026-08-29 | 🟢 L4 |  | 7 |
| [x] | **100. Same Tree** | Easy | recursion | 2026-08-12 | 🟢 L4 |  | 5 |
| [x] | **101. Symmetric Tree** | Easy | recursion | 2026-08-12 | 🟢 L4 |  | 6 |
| [x] | **104. Maximum Depth of Binary Tree** | Easy | recursion | 2026-08-09 | 🟢 L4 |  | 7 |
| [x] | **105. Construct Binary Tree from Preorder and Inorder Traversal** | Easy | recursion | 2026-08-13 | 🔴 L1 | true | 4 |
| [x] | **110. Balanced Binary Tree** | Easy | recursion | 2026-09-07 | 🟡 L3 | true | 4 |
| [x] | **112. Path Sum** | Easy | recursion | 2026-08-13 | 🟢 L4 |  | 6 |
| [x] | **226. Invert Binary Tree** | Easy | recursion | 2026-08-09 | 🟢 L4 |  | 8 |
| [x] | **236. Lowest Common Ancestor of a Binary Tree** | Medium | recursion | 2026-08-20 | 🟢 L4 |  | 10 |
| [x] | **437. Path Sum III** | Medium | recursion | 2026-08-18 | 🟢 L4 |  | 17 |
| [x] | **450. Delete Node in a BST** | Medium | recursion | 2026-09-08 | 🟠 L2 | true | 3 |
| [x] | **509. Fibonacci Number** | Easy | recursion | 2026-08-02 | 🟢 L4 |  | 13 |
| [x] | **530. Minimum Absolute Difference in BST** | Easy | recursion | 2026-09-10 | 🔴 L1 | true | 1 |
| [x] | **700. Search in a Binary Search Tree** | Easy | recursion | 2026-08-27 | 🟢 L4 |  | 9 |
| [x] | **701. Insert into a Binary Search Tree** | Medium | recursion | 2026-08-27 | 🟢 L4 |  | 9 |
| [x] | **872. Leaf-Similar Trees** | Easy | recursion | 2026-08-13 | 🟢 L4 |  | 6 |
| [x] | **1372. Longest ZigZag Path in a Binary Tree** | Medium | recursion | 2026-08-19 | 🟢 L4 |  | 7 |
| [x] | **1448. Count Good Nodes in Binary Tree** | Easy | recursion | 2026-08-16 | 🟢 L4 |  | 5 |

