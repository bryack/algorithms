# NeetCode Roadmap — Sliding Window

Цель: Освоить паттерн Sliding Window для оптимизации алгоритмов с подмассивами и подстроками.

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Оценка / Retention | Ревью |
| :---: | :--- | :---: | :--- | :--- | :--- | :--- |
| [x] | **643. Maximum Average Subarray I** | Easy | Fixed-size Sliding Window | 2026-05-18 | 🟢 L4 | |
| [x] | **1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold** | Medium | Fixed-size Sliding Window (count valid) | 2026-05-18 | 🟢 L4 | |
| [x] | **2090. K Radius Subarray Averages** | Medium | Fixed-size Sliding Window (radius average) | 2026-05-18 | 🟢 L4 | |
| [x] | **2461. Maximum Sum of Distinct Subarrays With Length K** | Medium | Fixed-size Sliding Window + distinct | 2026-05-19 | 🟢 L4 | |
| [x] | **1456. Maximum Number of Vowels in a Substring of Given Length** | Medium | Fixed-size Sliding Window (string counter) | 2026-05-19 | 🟢 L4 | |
| [x] | **2379. Minimum Recolors to Get K Consecutive Black Blocks** | Easy | Fixed-size Sliding Window (min ops) | 2026-05-19 | 🟢 L4 | |
| [x] | **1876. Substrings of Size Three with Distinct Characters** | Easy | Fixed-size Sliding Window (size 3) | 2026-05-18 | 🟢 L4 | |
| [x] | **2653. Sliding Subarray Beauty** | Medium | Fixed-size Sliding Window (order stat) | 2026-05-19 | 🟢 L4 | |
| [x] | **1423. Maximum Points You Can Obtain from Cards** | Medium | Fixed-size Sliding Window (edges) | 2026-06-01 | 🟢 L4 | |
| [x] | **3. Longest Substring Without Repeating Characters** | Medium | Dynamic Sliding Window (distinct chars) | 2026-05-21 | 🟢 L4 | |
| [x] | **209. Minimum Size Subarray Sum** | Medium | Dynamic Sliding Window (sum >= target) | 2026-06-01 | 🟢 L4 | |
| [x] | **1004. Max Consecutive Ones III** | Medium | Dynamic Sliding Window (at most K flips) | 2026-05-21 | 🟢 L4 | |
| [x] | **1493. Longest Subarray of 1's After Deleting One Element** | Medium | Dynamic Sliding Window (at most 1 zero) | 2026-05-21 | 🟢 L4 | |
| [x] | **904. Fruit Into Baskets** | Medium | Dynamic Sliding Window (at most 2 distinct) | 2026-05-21 | 🟢 L4 | |
| [x] | **424. Longest Repeating Character Replacement** | Medium | Dynamic Sliding Window (max freq) | 2026-06-03 | 🟢 L4 | |
| [x] | **567. Permutation in String** | Medium | Fixed-size Sliding Window + Frequency | 2026-06-03 | 🟢 L4 | |
| [x] | **438. Find All Anagrams in a String** | Medium | Fixed-size Sliding Window + Frequency | 2026-06-03 | 🟢 L4 | |
| [ ] | **76. Minimum Window Substring** | Hard | Dynamic Sliding Window (minimal valid) | | | |
| [ ] | **239. Sliding Window Maximum** | Hard | Monotonic Deque | | | |

## Дополнительная практика: Задачи с собеседований Яндекс

Задачи из списка Яндекса, решаемые паттерном Sliding Window и отсутствующие в основном блоке.

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Оценка / Retention | Ревью |
| :---: | :--- | :---: | :--- | :--- | :--- | :--- |
| [ ] | **487. Max Consecutive Ones II** | Medium | Dynamic Sliding Window (at most 1 zero flip) | | | |
| [ ] | **159. Longest Substring with At Most Two Distinct Characters** | Medium | Dynamic Sliding Window (at most 2 distinct) | | | |
| [ ] | **658. Find K Closest Elements** | Medium | Two Pointers / Sliding Window | | | |
| [ ] | **340. Longest Substring with At Most K Distinct Characters** | Medium | Sliding Window | | | |
| [ ] | **220. Contains Duplicate III** | Hard | Sliding Window (bucket/ordered set) | | | |
| [ ] | **395. Longest Substring with At Least K Repeating Characters** | Medium | Sliding Window / Divide and Conquer | | | |

## Легенда оценки (Retention)

| Метка | Значение | Когда ставить |
|-------|----------|---------------|
| 🟢 **L4** | Решила сама, быстро, чистый код | Без подсказок, за разумное время, сразу верный Big O |
| 🟡 **L3** | Решила сама, но медленно или с багами | Сама, но долго думала или баги в edge cases |
| 🟠 **L2** | Нужны были подсказки по алгоритму | Я направил к паттерну, но код написала сама |
| 🔴 **L1** | Существенная помощь, нужно повторить | Помогал с алгоритмом и/или кодом |
| ⚪ **L5** | Эксперт — могу объяснить другому | Решила, оптимизировала, знаю все edge cases |

### Критерии оценки

1. **Самостоятельность** — пришла ли к решению сама или нужны были подсказки?
2. **Скорость** — сколько времени заняло решение (относительно сложности)?
3. **Понимание** — можешь ли объяснить алгоритм своими словами и воспроизвести позже?
4. **Качество кода** — обработаны ли edge cases, чистый ли код, верный ли Big O анализ?
5. **Нужно ли повторение** — ставить дату ревью, если оценка L1-L3

---

## Правило перехода

> **Не начинать новую задачу, пока текущая не доведена до L4.**
>
> Исключение: разминка (warm-up) перед новой задачей — повторение уже решённой ([x]) задачи.

---

## Ежедневная разминка (Repetition)

0. **Warm-up**: Перед новой задачей решаем одну из уже выполненных ([x]) самостоятельно.
   - **Вербализация**: Сначала описываешь алгоритм и код словами (почему и как).
   - **Zero Hints**: Я не даю подсказок во время решения разминки, только проверяю итог.
   - **Без тестов**: На разминке тесты не пишем — моделируем ситуацию на собеседовании, где времени на TDD нет.

---

## Как работаем

1. **TDD**: Сначала тесты в `day_X/main_test.go`.
2. **Go Way**: Используем `make` с аллокацией, `struct{}` для сетов, `range` для итераций.
3. **Complexity**: Анализируем Time и Space Complexity для каждого решения.
