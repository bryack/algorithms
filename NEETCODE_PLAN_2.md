# NeetCode Roadmap — Two Pointers

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
| [x] | **11. Container With Most Water** | Medium | Two Pointers (maximization) | 2026-05-06 | ⚪ L5 |  | 11 |
| [x] | **15. 3Sum** | Medium | Two Pointers + Sorting | 2026-05-04 | ⚪ L5 |  | 7 |
| [x] | **18. 4Sum** | Medium | Two Pointers + Sorting | 2026-05-05 | ⚪ L5 |  | 7 |
| [x] | **26. Remove Duplicates from Sorted Array** | Easy | Two Pointers (fast & slow, overwrite) | 2026-05-04 | ⚪ L5 |  | 12 |
| [x] | **42. Trapping Rain Water** | Hard | Two Pointers / Prefix Max | 2026-05-07 | ⚪ L5 |  | 10 |
| [x] | **88. Merge Sorted Array** | Easy | Two Pointers (fill from end) |  |  |  |  |
| [x] | **125. Valid Palindrome** | Easy | Two Pointers (left/right) | 2026-05-01 | ⚪ L5 |  | 10 |
| [x] | **167. Two Sum II - Input Array Is Sorted** | Medium | Two Pointers (sorted array, converge) | 2026-05-02 | ⚪ L5 |  | 10 |
| [x] | **189. Rotate Array** | Medium | Two Pointers / Reverse | 2026-05-06 | ⚪ L5 |  | 7 |
| [x] | **344. Reverse String** | Easy | Two Pointers (swap in-place) | 2026-05-01 | ⚪ L5 |  | 1 |
| [x] | **392. Is Subsequence** | Easy | Two Pointers (converging check) | 2026-05-04 | ⚪ L5 |  | 13 |
| [x] | **680. Valid Palindrome II** | Easy | Two Pointers (left/right, skip check) | 2026-05-03 | ⚪ L5 |  | 6 |
| [x] | **881. Boats to Save People** | Medium | Two Pointers (greedy pairing) | 2026-05-06 | ⚪ L5 |  | 7 |
| [x] | **1768. Merge Strings Alternately** | Easy | Two Pointers (interleave) | 2026-05-03 | ⚪ L5 |  | 6 |

## Bonus (NeetCode 150, уже решены)

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **283. Move Zeroes** | Easy | Two Pointers (fast & slow) | 2026-05-01 | ⚪ L5 |  | 9 |
| [x] | **977. Squares of a Sorted Array** | Easy | Two Pointers (fill from end) | 2026-05-02 | ⚪ L5 |  | 4 |

## Дополнительная практика: Палиндромы

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **5. Longest Palindromic Substring** | Medium | Two Pointers (expand around center) | 2026-05-09 | ⚪ L5 |  | 16 |
| [x] | **27. Remove Element** | Easy | Two Pointers | 2026-06-04 | ⚪ L5 |  | 10 |
| [x] | **28. Find the Index of the First Occurrence in a String** | Easy | Two Pointers | 2026-06-05 | 🟢 L4 | true | 19 |
| [x] | **80. Remove Duplicates from Sorted Array II** | Medium | Two Pointers | 2026-06-04 | ⚪ L5 |  | 11 |
| [x] | **161. One Edit Distance** | Medium | Two Pointers (compare strings with one allowed edit) | 2026-04-11 | ⚪ L5 |  | 10 |
| [x] | **228. Summary Ranges** | Easy | Two Pointers (consecutive range tracking) | 2026-05-24 | ⚪ L5 |  | 12 |
| [x] | **234. Palindrome Linked List** | Easy | Two Pointers | 2026-06-21 |  |  | 8 |
| [x] | **253. Meeting Rooms II** | Medium | Two Pointers (start/end) | 2026-06-05 | 🟢 L4 |  | 7 |
| [x] | **349. Intersection of Two Arrays** | Easy | Two Pointers / Sorting | 2026-06-04 | ⚪ L5 |  | 8 |
| [x] | **350. Intersection of Two Arrays II** | Easy | Two Pointers / Sorting | 2026-04-25 | ⚪ L5 |  | 5 |
| [x] | **415. Add Strings** | Easy | Two Pointers (add digit strings from end with carry) | 2026-05-31 | ⚪ L5 |  | 10 |
| [x] | **443. String Compression** | Medium | Two Pointers (fast & slow, in-place overwrite) | 2026-05-24 | ⚪ L5 |  | 20 |
| [x] | **557. Reverse Words in a String III** | Easy | Two Pointers (reverse each word in-place) | 2026-05-31 | ⚪ L5 |  | 9 |
| [x] | **647. Palindromic Substrings** | Medium | Two Pointers (expand around center, count) | 2026-05-31 | ⚪ L5 |  | 10 |
| [x] | **917. Reverse Only Letters** | Easy | Two Pointers (swap letters, skip non-letters) | 2026-05-08 | ⚪ L5 |  | 8 |
| [x] | **986. Interval List Intersections** | Medium | Two Pointers (two sorted interval lists) | 2026-05-31 | ⚪ L5 |  | 8 |

## Вне каталога (Extra)

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **121. Best Time to Buy and Sell Stock** |  | two_pointers | 2026-06-18 | 🟢 L4 |  | 10 |
| [x] | **844. Backspace String Compare** |  | two_pointers | 2026-05-13 | ⚪ L5 |  | 16 |
| [x] | **1679. Max Number of K-Sum Pairs** |  | two_pointers | 2026-07-11 | 🟢 L4 |  | 5 |

> 📌 Заметки выше решены, но отсутствуют в `catalog.yaml`. Если они часть роадмапа — добавь их в каталог.

