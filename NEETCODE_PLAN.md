# NeetCode Roadmap — Arrays & Hashing

Цель: Закрепить фундамент работы с массивами и хэш-таблицами перед переходом к более сложным структурам.

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Ревью |
| :---: | :--- | :---: | :--- | :--- | :--- |
| [x] | **217. Contains Duplicate** | Easy | Hash Set (map[int]struct{}) | 2026-04-24 | 2026-06-28 (повтор, хорошо усвоено) |
| [x] | **242. Valid Anagram** | Easy | Frequency Map / Array [26]int | 2026-04-24 | 2026-05-03 (Efficiency) |
| [x] | **1. Two Sum** | Easy | HashMap (One-pass) | 2026-04-24 | |
| [x] | **387. First Unique Char** | Easy | Frequency Counting | 2026-04-24 | |
| [x] | **2351. First Letter Twice** | Easy | Frequency Counting | 2026-04-24 | |
| [x] | **49. Group Anagrams** | Medium | Categorization / Map Keys | 2026-04-24 | |
| [x] | **205. Isomorphic Strings** | Easy | String Normalization / Mapping | 2026-04-25 | 2026-05-04 (Algorithmic Logic) |
| [x] | **219. Contains Duplicate II** | Easy | Sliding Window / HashMap | 2026-04-25 | |
| [x] | **1002. Find Common Chars** | Easy | Frequency Intersection | 2026-04-25 | |
| [x] | **Practice: 350. Intersection of Two Arrays II** | Easy | Frequency Counting / Intersection | 2026-04-25 | |
| [x] | **Practice: 383. Ransom Note** | Easy | Frequency Counting / Sufficiency | 2026-04-26 |
| [x] | **Practice: 1189. Max Num of Balloons** | Easy | Frequency Counting / Limiting Factor | 2026-04-26 |
| [x] | **347. Top K Frequent Elements** | Medium | Bucket Sort / Heap | 2026-04-26 |
| [x] | **Practice: 451. Sort Characters By Frequency** | Medium | Bucket Sort / String Building | 2026-04-26 | |
| [x] | **Practice: 1636. Sort Array by Inc Frequency** | Easy | Bucket Sort / Custom Sort | 2026-04-27 | |
| [x] | **Practice: 692. Top K Frequent Words** | Medium | Bucket Sort / Lexicographical Sort | 2026-04-27 | |
| [x] | **238. Product of Array Except Self** | Medium | Prefix & Suffix Products | 2026-04-27 |
| [x] | **Practice: 303. Range Sum Query - Immutable** | Easy | Prefix Sums | 2026-04-27 | |
| [x] | **1480. Running Sum of 1d Array** | Easy | Prefix Sum | 2026-04-28 | |
| [x] | **1732. Find the Highest Altitude** | Easy | Prefix Sum | 2026-04-28 | |
| [x] | **724. Find Pivot Index** | Easy | Prefix Sum | 2026-04-28 | |
| [x] | **Practice: 560. Subarray Sum Equals K** | Medium | Prefix Sums + HashMap | 2026-04-28 | |
| [ ] | **Practice: 525. Contiguous Array** | Medium | Prefix Sums (0 as -1) | | |
| [ ] | **128. Longest Consecutive Sequence** | Medium | Intelligent HashSet Search | |

## Ежедневная разминка (Repetition)
0. **Warm-up**: Перед новой задачей решаем одну из уже выполненных ([x]) самостоятельно.
   - **Вербализация**: Сначала описываешь алгоритм и код словами (почему и как).
   - **Zero Hints**: Я не даю подсказок во время решения разминки, только проверяю итог.
   - **Без тестов**: На разминке тесты не пишем — моделируем ситуацию на собеседовании, где времени на TDD нет.

## Как работаем
1. **TDD**: Сначала тесты в `day_X/main_test.go`.
2. **Go Way**: Используем `make` с аллокацией, `struct{}` для сетов, `range` для итераций.
3. **Complexity**: Анализируем Time и Space Complexity для каждого решения.
