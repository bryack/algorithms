# NeetCode Roadmap — SQL

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
## Select

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **584. Find Customer Referee** | Easy | Select (WHERE, NULL handling) | 2026-07-24 | 🟢 L4 |  | 3 |
| [x] | **595. Big Countries** | Easy | Select (WHERE, multiple conditions) | 2026-07-24 | 🟢 L4 |  | 4 |
| [x] | **1148. Article Views I** | Easy | Select (DISTINCT, WHERE) | 2026-07-24 | 🟢 L4 |  | 3 |
| [x] | **1241. Number of Comments per Post** | Easy | Aggregate (COUNT, GROUP BY) | 2026-07-25 | 🟡 L3 | true | 4 |
| [x] | **1683. Invalid Tweets** | Easy | Select (WHERE, string length) | 2026-07-24 | 🟢 L4 |  | 3 |
| [x] | **1757. Recyclable and Low Fat Products** | Easy | Select (WHERE) | 2026-07-24 | 🟢 L4 |  | 3 |
| [x] | **2504. Concatenate the Name and the Profession** | Easy | Select (CONCAT) | 2026-07-24 | 🟢 L4 |  | 6 |
| [x] | **3436. Find Valid Emails** | Easy | Select (WHERE, REGEXP) | 2026-07-24 | 🟢 L4 |  | 3 |

## Basic Joins

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **197. Rising Temperature** | Easy | JOIN (self join, date diff) | 2026-07-26 | 🟡 L3 | true | 5 |
| [x] | **570. Managers with at Least 5 Direct Reports** | Medium | JOIN (self join, GROUP BY HAVING) | 2026-07-30 | 🔴 L1 | true | 2 |
| [x] | **577. Employee Bonus** | Easy | JOIN (left join, filter NULL) | 2026-07-28 | 🟢 L4 |  | 4 |
| [x] | **1068. Product Sales Analysis I** | Easy | JOIN (inner join) | 2026-07-25 | 🟡 L3 |  | 4 |
| [x] | **1280. Students and Examinations** | Hard | JOIN (cross join + left join, COUNT) | 2026-07-29 | 🟠 L2 | true | 3 |
| [x] | **1378. Replace Employee ID With The Unique Identifier** | Easy | JOIN (left join) | 2026-07-25 | 🟢 L4 |  | 4 |
| [x] | **1581. Customer Who Visited but Did Not Make Any Transactions** | Easy | JOIN (left join, filter NULL) | 2026-07-26 | 🟢 L4 | true | 5 |
| [x] | **1661. Average Time of Process per Machine** | Easy | JOIN (self join, AVG) | 2026-07-27 | 🟠 L2 | true | 4 |
| [x] | **1934. Confirmation Rate** | Medium | JOIN (left join, AVG, IF) | 2026-07-31 | 🔴 L1 | true | 1 |

## Basic Aggregate Functions

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [ ] | **550. Game Play Analysis IV** | Medium | Aggregate (DATEDIFF, subquery) |  |  |  |  |
| [ ] | **620. Not Boring Movies** | Easy | Aggregate (WHERE + ORDER BY) |  |  |  |  |
| [ ] | **1075. Project Employees I** | Easy | Aggregate (AVG, JOIN) |  |  |  |  |
| [ ] | **1174. Immediate Food Delivery II** | Medium | Aggregate (AVG, subquery) |  |  |  |  |
| [ ] | **1193. Monthly Transactions I** | Medium | Aggregate (COUNT, SUM, GROUP BY) |  |  |  |  |
| [ ] | **1211. Queries Quality and Percentage** | Medium | Aggregate (AVG, ROUND) |  |  |  |  |
| [ ] | **1251. Average Selling Price** | Medium | Aggregate (AVG, JOIN, BETWEEN) |  |  |  |  |
| [ ] | **1633. Percentage of Users Attended a Contest** | Medium | Aggregate (COUNT, GROUP BY, ROUND) |  |  |  |  |

## Sorting and Grouping

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [ ] | **596. Classes With at Least 5 Students** | Easy | Grouping (GROUP BY HAVING) |  |  |  |  |
| [ ] | **619. Biggest Single Number** | Hard | Subquery (MAX of singletons) |  |  |  |  |
| [ ] | **1045. Customers Who Bought All Products** | Medium | Grouping (HAVING COUNT = total) |  |  |  |  |
| [ ] | **1070. Product Sales Analysis III** | Medium | Subquery (MIN, GROUP BY) |  |  |  |  |
| [ ] | **1141. User Activity for the Past 30 Days I** | Easy | Sorting (WHERE date range) |  |  |  |  |
| [ ] | **1729. Find Followers Count** | Easy | Aggregate (COUNT, GROUP BY) |  |  |  |  |
| [ ] | **2356. Number of Unique Subjects Taught by Each Teacher** | Easy | Aggregate (COUNT DISTINCT, GROUP BY) |  |  |  |  |

## Advanced Select and Joins

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [ ] | **180. Consecutive Numbers** | Medium | JOIN (self join x3) |  |  |  |  |
| [ ] | **610. Triangle Judgement** | Easy | Select (WHERE, CASE) |  |  |  |  |
| [ ] | **1164. Product Price at a Given Date** | Medium | Subquery (window/filtered) |  |  |  |  |
| [ ] | **1204. Last Person to Fit in the Bus** | Hard | Window function (running sum) |  |  |  |  |
| [ ] | **1731. The Number of Employees Which Report to Each Employee** | Easy | JOIN (self join, aggregate) |  |  |  |  |
| [ ] | **1789. Primary Department for Each Employee** | Medium | Subquery (UNION + filter) |  |  |  |  |
| [ ] | **1907. Count Salary Categories** | Medium | CASE + aggregate |  |  |  |  |

## Subqueries

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [ ] | **185. Department Top Three Salaries** | Hard | Window function (DENSE_RANK) |  |  |  |  |
| [ ] | **585. Investments in 2016** | Medium | Subquery (COUNT, filter) |  |  |  |  |
| [ ] | **602. Friend Requests II: Who Has the Most Friends** | Medium | Aggregate (UNION ALL, COUNT) |  |  |  |  |
| [ ] | **626. Exchange Seats** | Medium | Subquery (CASE, MOD) |  |  |  |  |
| [ ] | **1321. Restaurant Growth** | Medium | Window function (running pct) |  |  |  |  |
| [ ] | **1341. Movie Rating** | Medium | Subquery (ORDER BY LIMIT) |  |  |  |  |
| [ ] | **1978. Employees Whose Manager Left the Company** | Easy | Subquery (NOT IN) |  |  |  |  |

## Advanced String Functions / Regex / Clause

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [ ] | **176. Second Highest Salary** | Medium | Subquery (MAX, IFNULL) |  |  |  |  |
| [ ] | **196. Delete Duplicate Emails** | Easy | DELETE + self join |  |  |  |  |
| [ ] | **1327. List the Products Ordered in a Period** | Easy | JOIN + HAVING |  |  |  |  |
| [ ] | **1484. Group Sold Products By The Date** | Easy | Aggregate (GROUP_CONCAT) |  |  |  |  |
| [ ] | **1517. Find Users With Valid E-Mails** | Medium | Regex (REGEXP_LIKE) |  |  |  |  |
| [ ] | **1527. Patients With a Condition** | Medium | String (LIKE, pattern) |  |  |  |  |
| [ ] | **1667. Fix Names in a Table** | Easy | String (UPPER, LOWER, CONCAT) |  |  |  |  |

## Вне каталога (Extra)

| Статус | Задача | Сложность | Тема / Паттерн | Дата | Retention | needRepeat | Решений |
| :---: | :--- | :---: | :--- | :--- | :--- | :---: | :---: |
| [x] | **3586. Find COVID Recovery Patients** |  | select | 2026-07-29 | 🟠 L2 | true | 3 |

> 📌 Заметки выше решены, но отсутствуют в `catalog.yaml`. Если они часть роадмапа — добавь их в каталог.

