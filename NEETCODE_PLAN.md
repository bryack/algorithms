# NeetCode Roadmap — Arrays & Hashing

Цель: Закрепить фундамент работы с массивами и хэш-таблицами перед переходом к более сложным структурам.

| Статус | Задача | Сложность | Тема / Паттерн | Дата |
| :---: | :--- | :---: | :--- | :--- |
| [x] | **217. Contains Duplicate** | Easy | Hash Set (map[int]struct{}) | 2026-04-24 |
| [x] | **242. Valid Anagram** | Easy | Frequency Map / Array [26]int | 2026-04-24 |
| [x] | **1. Two Sum** | Easy | HashMap (One-pass) | 2026-04-24 |
| [x] | **387. First Unique Char** | Easy | Frequency Counting | 2026-04-24 |
| [x] | **2351. First Letter Twice** | Easy | Frequency Counting | 2026-04-24 |
| [x] | **49. Group Anagrams** | Medium | Categorization / Map Keys | 2026-04-24 |
| [x] | **205. Isomorphic Strings** | Easy | String Normalization / Mapping | 2026-04-25 |
| [x] | **219. Contains Duplicate II** | Easy | Sliding Window / HashMap | 2026-04-25 |
| [x] | **1002. Find Common Chars** | Easy | Frequency Intersection | 2026-04-25 |
| [x] | **Practice: 350. Intersection of Two Arrays II** | Easy | Frequency Counting / Intersection | 2026-04-25 |
| [ ] | **Practice: 383. Ransom Note** | Easy | Frequency Counting / Sufficiency | |
| [ ] | **Practice: 1189. Max Num of Balloons** | Easy | Frequency Counting / Limiting Factor | |
| [ ] | **347. Top K Frequent Elements** | Medium | Bucket Sort / Heap | |
| [ ] | **238. Product of Array Except Self** | Medium | Prefix & Suffix Products | |
| [ ] | **128. Longest Consecutive Sequence** | Medium | Intelligent HashSet Search | |

## Как работаем
1. **TDD**: Сначала тесты в `day_X/main_test.go`.
2. **Go Way**: Используем `make` с аллокацией, `struct{}` для сетов, `range` для итераций.
3. **Complexity**: Анализируем Time и Space Complexity для каждого решения.
