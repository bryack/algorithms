# NeetCode Roadmap — Sliding Window

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
| [x] | **3. Longest Substring Without Repeating Characters** | Medium | Dynamic Sliding Window (distinct chars) | 2026-05-21 | 🟢 L4 |  | 6 |
| [x] | **76. Minimum Window Substring** | Hard | Dynamic Sliding Window (minimal valid) | 2026-06-16 | 🟠 L2 | true | 10 |
| [x] | **209. Minimum Size Subarray Sum** | Medium | Dynamic Sliding Window (sum >= target) | 2026-05-21 | 🟢 L4 |  | 10 |
| [x] | **239. Sliding Window Maximum** | Hard | Monotonic Deque | 2026-06-18 | 🔴 L1 | true | 10 |
| [x] | **424. Longest Repeating Character Replacement** | Medium | Dynamic Sliding Window (max freq) | 2026-06-03 | 🟢 L4 |  | 16 |
| [x] | **438. Find All Anagrams in a String** | Medium | Fixed-size Sliding Window + Frequency | 2026-06-03 |  |  | 7 |
| [x] | **567. Permutation in String** | Medium | Fixed-size Sliding Window + Frequency | 2026-06-03 | 🟢 L4 |  | 11 |
| [x] | **643. Maximum Average Subarray I** | Easy | Fixed-size Sliding Window | 2026-05-18 | 🟢 L4 |  | 10 |
| [x] | **904. Fruit Into Baskets** | Medium | Dynamic Sliding Window (at most 2 distinct) | 2026-05-21 | 🟢 L4 |  | 7 |
| [x] | **1004. Max Consecutive Ones III** | Medium | Dynamic Sliding Window (at most K flips) | 2026-05-21 | 🟢 L4 |  | 8 |
| [x] | **1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold** | Medium | Fixed-size Sliding Window (count valid) | 2026-05-18 | 🟢 L4 |  | 8 |
| [x] | **1423. Maximum Points You Can Obtain from Cards** | Medium | Fixed-size Sliding Window (edges) | 2026-05-20 | 🟢 L4 |  | 10 |
| [x] | **1456. Maximum Number of Vowels in a Substring of Given Length** | Medium | Fixed-size Sliding Window (string counter) | 2026-05-19 | 🟢 L4 |  | 8 |
| [x] | **1493. Longest Subarray of 1's After Deleting One Element** | Medium | Dynamic Sliding Window (at most 1 zero) | 2026-05-21 | 🟢 L4 |  | 6 |
| [x] | **1876. Substrings of Size Three with Distinct Characters** | Easy | Fixed-size Sliding Window (size 3) | 2026-05-18 | 🟢 L4 |  | 7 |
| [x] | **2090. K Radius Subarray Averages** | Medium | Fixed-size Sliding Window (radius average) | 2026-05-18 | 🟢 L4 |  | 10 |
| [x] | **2379. Minimum Recolors to Get K Consecutive Black Blocks** | Easy | Fixed-size Sliding Window (min ops) | 2026-05-19 | 🟢 L4 |  | 7 |
| [x] | **2461. Maximum Sum of Distinct Subarrays With Length K** | Medium | Fixed-size Sliding Window + distinct | 2026-05-19 | 🟢 L4 |  | 10 |
| [x] | **2653. Sliding Subarray Beauty** | Medium | Fixed-size Sliding Window (order stat) | 2026-05-19 | 🟢 L4 |  | 16 |

## Дополнительная практика: Задачи с собеседований Яндекс

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **159. Longest Substring with At Most Two Distinct Characters** | Medium | Dynamic Sliding Window (at most 2 distinct) | 2026-06-12 |  |  | 5 |
| [x] | **220. Contains Duplicate III** | Hard | Sliding Window (bucket/ordered set) | 2026-06-14 | 🟢 L4 |  | 14 |
| [x] | **340. Longest Substring with At Most K Distinct Characters** | Medium | Sliding Window | 2026-06-13 | ⚪ L5 |  | 5 |
| [x] | **395. Longest Substring with At Least K Repeating Characters** | Medium | Sliding Window / Divide and Conquer | 2026-06-16 | 🟢 L4 | true | 13 |
| [x] | **487. Max Consecutive Ones II** | Medium | Dynamic Sliding Window (at most 1 zero flip) | 2026-06-12 | 🟢 L4 |  | 3 |
| [x] | **658. Find K Closest Elements** | Medium | Two Pointers / Sliding Window | 2026-06-13 | 🟢 L4 |  | 12 |
| [ ] | **713. Subarray Product Less Than K** | Medium | Sliding Window (incremental product update) |  |  |  |  |
| [ ] | **862. Shortest Subarray with Sum at Least K** | Hard | Sliding Window + Monotonic Deque |  |  |  |  |
| [ ] | **1151. Minimum Swaps to Group All 1's Together** | Medium | Fixed-size Sliding Window (min zeros in window) |  |  |  |  |
| [ ] | **1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit** | Hard | Sliding Window + Monotonic Deque (lazy min/max) |  |  |  |  |
| [ ] | **2024. Maximize the Confusion of an Exam** | Medium | Non-shrinking Window (lazy max, аналог LC 1004/424) |  |  |  |  |

