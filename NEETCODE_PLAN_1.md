# NeetCode Roadmap — Arrays & Hashing

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
| [x] | **1. Two Sum** | Easy | HashMap (One-pass) | 2026-07-09 | 🟢 L4 |  | 8 |
| [x] | **14. Longest Common Prefix** | Easy | String (vertical/horizontal scanning) | 2026-07-04 | 🟢 L4 |  | 13 |
| [x] | **36. Valid Sudoku** | Medium | Matrix / HashSet | 2026-04-30 |  |  | 1 |
| [x] | **49. Group Anagrams** | Medium | Categorization / Map Keys | 2026-04-24 |  |  | 12 |
| [x] | **128. Longest Consecutive Sequence** | Medium | Intelligent HashSet Search | 2026-04-29 |  |  | 2 |
| [x] | **169. Majority Element** | Easy | Boyer-Moore Voting / HashMap | 2026-07-04 |  |  | 7 |
| [x] | **170. Two Sum III - Data structure design** | Easy | HashMap Design | 2026-07-10 | 🟢 L4 |  | 7 |
| [x] | **202. Happy Number** | Easy | HashSet / Floyd | 2026-04-30 |  |  | 5 |
| [x] | **205. Isomorphic Strings** | Easy | String Normalization / Mapping | 2026-04-25 | 🟢 L4 |  | 11 |
| [x] | **217. Contains Duplicate** | Easy | Hash Set (map[int]struct{}) | 2026-04-24 |  |  | 4 |
| [x] | **219. Contains Duplicate II** | Easy | Sliding Window / HashMap | 2026-04-25 | 🟢 L4 |  | 5 |
| [x] | **238. Product of Array Except Self** | Medium | Prefix & Suffix Products | 2026-04-27 | 🔴 L1 |  | 2 |
| [x] | **242. Valid Anagram** | Easy | Frequency Map / Array [26]int | 2026-04-24 | 🟢 L4 |  | 8 |
| [x] | **303. Range Sum Query - Immutable** | Easy | Prefix Sums | 2026-04-27 |  |  | 1 |
| [x] | **347. Top K Frequent Elements** | Medium | Bucket Sort / Heap | 2026-04-26 |  |  | 2 |
| [x] | **350. Intersection of Two Arrays II** | Easy | Frequency Counting / Intersection | 2026-04-25 | ⚪ L5 |  | 5 |
| [x] | **380. Insert Delete GetRandom O(1)** | Medium | Map + Dynamic Array | 2026-05-01 |  |  | 1 |
| [x] | **383. Ransom Note** | Easy | Frequency Counting / Sufficiency | 2026-04-26 |  |  | 2 |
| [x] | **387. First Unique Char** | Easy | Frequency Counting | 2026-04-24 |  |  | 3 |
| [x] | **451. Sort Characters By Frequency** | Medium | Bucket Sort / String Building | 2026-04-26 |  |  | 2 |
| [x] | **525. Contiguous Array** | Medium | Prefix Sums (0 as -1) | 2026-04-29 |  |  | 1 |
| [x] | **532. K-diff Pairs in an Array** | Medium | HashSet / Two Pointers | 2026-04-30 |  |  | 1 |
| [x] | **560. Subarray Sum Equals K** | Medium | Prefix Sums + HashMap | 2026-04-28 |  |  | 1 |
| [x] | **692. Top K Frequent Words** | Medium | Bucket Sort / Lexicographical Sort | 2026-04-27 |  |  | 1 |
| [x] | **706. Design HashMap** | Easy | HashMap Design (array of buckets) | 2026-07-04 | 🟢 L4 |  | 13 |
| [x] | **724. Find Pivot Index** | Easy | Prefix Sum | 2026-04-28 |  |  | 5 |
| [x] | **1002. Find Common Chars** | Easy | Frequency Intersection | 2026-04-04 | 🟢 L4 |  | 11 |
| [x] | **1189. Max Num of Balloons** | Easy | Frequency Counting / Limiting Factor | 2026-04-26 |  |  | 2 |
| [x] | **1480. Running Sum of 1d Array** | Easy | Prefix Sum | 2026-04-28 |  |  | 2 |
| [x] | **1636. Sort Array by Inc Frequency** | Easy | Bucket Sort / Custom Sort | 2026-04-27 |  |  | 2 |
| [x] | **1732. Find the Highest Altitude** | Easy | Prefix Sum | 2026-04-28 |  |  | 3 |
| [x] | **1929. Concatenation of Array** | Easy | Array | 2026-07-04 |  |  | 8 |
| [x] | **2351. First Letter Twice** | Easy | Frequency Counting | 2026-04-24 |  |  | 2 |

## Задачи с собеседований Яндекс

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **121. Best Time to Buy and Sell Stock** | Easy | Array / Prefix Min | 2026-06-18 | 🟢 L4 |  | 10 |
| [ ] | **136. Single Number** | Easy | Bit Manipulation / HashSet |  |  |  |  |
| [x] | **485. Max Consecutive Ones** | Easy | Array (linear scan) | 2026-04-11 |  |  | 1 |
| [x] | **1436. Destination City** | Easy | HashSet (single sink) | 2026-07-14 | 🟢 L4 |  | 5 |
| [ ] | **1496. Path Crossing** | Easy | HashSet (visited points) |  |  |  |  |
| [x] | **2215. Find the Difference of Two Arrays** | Easy | HashSet Difference | 2026-07-13 | 🟢 L4 |  | 5 |

## Вне каталога (Extra)

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **58. Length of Last Word** |  | array, strings | 2026-07-23 | 🟢 L4 |  | 4 |
| [x] | **151. Reverse Words in a String** |  | array, strings | 2026-07-26 | 🟠 L2 | true | 3 |
| [x] | **290. Word Pattern** |  | array, bijection | 2026-07-24 | 🟢 L4 |  | 6 |
| [x] | **334. Increasing Triplet Subsequence** |  | array, strings | 2026-07-29 | 🔴 L1 | true | 2 |
| [x] | **345. Reverse Vowels of a String** |  | strings | 2026-07-14 | 🟢 L4 |  | 4 |
| [x] | **605. Can Place Flowers** |  | array | 2026-07-16 | 🟢 L4 |  | 5 |
| [x] | **1071. Greatest Common Divisor of Strings** |  | Euclidean_algorithm, strings | 2026-04-05 | 🟢 L4 |  | 9 |
| [x] | **1207. Unique Number of Occurrences** |  | array | 2026-07-17 | 🟢 L4 |  | 12 |
| [x] | **1431. Kids With the Greatest Number of Candies** |  | hash | 2026-07-15 | 🟢 L4 |  | 4 |
| [x] | **1991. Find the Middle Index in Array** |  | prefix_sums | 2026-05-12 |  |  | 1 |

> 📌 Заметки выше решены, но отсутствуют в `catalog.yaml`. Если они часть роадмапа — добавь их в каталог.

