# NeetCode Roadmap — Linked List

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
## Основной блок: Linked List

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **2. Add Two Numbers** | Medium | Two Pointers + Carry | 2026-06-25 | 🟢 L4 |  | 10 |
| [x] | **19. Remove Nth Node From End of List** | Medium | Two Pointers (gap n) | 2026-06-21 | 🟢 L4 |  | 19 |
| [x] | **21. Merge Two Sorted Lists** | Easy | Two Pointers (merge) | 2026-04-12 | 🟢 L4 |  | 11 |
| [x] | **24. Swap Nodes in Pairs** | Medium | Two Pointers (pair swap) | 2026-06-21 | 🟢 L4 |  | 19 |
| [x] | **61. Rotate List** | Medium | Two Pointers (k mod n) | 2026-06-26 | 🟢 L4 |  | 14 |
| [x] | **83. Remove Duplicates from Sorted List** | Easy | Two Pointers (skip duplicates) | 2026-04-14 | 🟢 L4 |  | 16 |
| [x] | **92. Reverse Linked List II** | Medium | Reverse Sublist (left/right bounds) | 2026-06-25 | 🟢 L4 |  | 13 |
| [x] | **138. Copy List with Random Pointer** | Medium | HashMap / Interweave | 2026-06-25 | 🟢 L4 |  | 16 |
| [x] | **141. Linked List Cycle** | Easy | Floyd's Cycle Detection (fast/slow) | 2026-06-24 | 🟢 L4 |  | 11 |
| [x] | **143. Reorder List** | Medium | Find Middle + Reverse + Merge | 2026-06-25 | 🟢 L4 |  | 12 |
| [x] | **146. LRU Cache** | Medium | HashMap + Doubly Linked List | 2026-06-08 | 🟢 L4 | true | 24 |
| [x] | **160. Intersection of Two Linked Lists** | Medium | Two Pointers (swap lists on reaching end) | 2026-04-14 | 🟢 L4 |  | 9 |
| [x] | **203. Remove Linked List Elements** | Easy | Dummy Node / Two Pointers | 2026-04-04 | 🟢 L4 |  | 10 |
| [x] | **206. Reverse Linked List** | Easy | Iterative Reversal (prev/curr/next) | 2026-06-21 | 🟢 L4 |  | 8 |
| [x] | **234. Palindrome Linked List** | Easy | Two Pointers + Reverse (middle) | 2026-06-21 | 🟢 L4 |  | 11 |
| [x] | **237. Delete Node in a Linked List** | Medium | Node Copy Trick | 2026-06-26 | 🟢 L4 |  | 11 |
| [x] | **287. Find the Duplicate Number** | Medium | Floyd's Cycle Detection | 2026-06-25 | 🟢 L4 |  | 15 |
| [x] | **328. Odd Even Linked List** | Easy | linked_lists | 2026-07-22 | 🟢 L4 |  | 5 |
| [x] | **460. LFU Cache** | Hard | HashMap + Frequency Buckets | 2026-06-11 | 🟡 L3 |  | 17 |
| [x] | **707. Design Linked List** | Medium | Singly/Doubly Linked List Design | 2026-06-18 |  |  | 15 |
| [x] | **876. Middle of the Linked List** | Easy | Two Pointers (fast/slow) | 2026-06-18 | 🟢 L4 |  | 12 |
| [x] | **2095. Delete the Middle Node of a Linked List** | Medium | Two Pointers (fast/slow) | 2026-06-21 | 🟢 L4 |  | 16 |
| [x] | **2130. Maximum Twin Sum of a Linked List** | Medium | linked_lists | 2026-07-24 | 🟢 L4 |  | 5 |

## Дополнительная практика: Задачи с собеседований Яндекс

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **23. Merge k Sorted Lists** | Hard | Linked List + Divide and Conquer / Heap | 2026-06-26 |  | true | 4 |
| [x] | **622. Design Circular Queue** | Medium | Array-based Circular Queue | 2026-06-25 |  |  | 11 |

