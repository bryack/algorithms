# Алгоритмические паттерны

## GCD (Алгоритм Евклида / НОД)

### Когда использовать
> GCD нужен, когда в задаче есть две или более величины, которые нужно "разделить поровну" на одинаковые части, и мы ищем максимально возможный размер этой части.

### Ключевые слова в условии
- "сократить", "simplified", "reduced", "несократимые дроби"
- "общий делитель", "greatest common divisor", "GCD"
- "разделить на группы одинакового размера"
- "повторяющийся паттерн" в строках
- "наибольший общий..."
- "взаимно простые" (GCD == 1)

### Примеры задач

| Задача | Что делим | На что делим | GCD даёт |
|--------|-----------|--------------|----------|
| Калькулятор дробей | Числитель и знаменатель | Общий делитель | Коэффициент сокращения |
| LeetCode 914 (карты) | Частоты карт (значения map) | Размер группы X | Максимальный размер группы |
| LeetCode 1071 (строки) | Длины строк | Длина подстроки | Длина GCD-строки |
| LeetCode 1447 (дроби) | Числитель и знаменатель | Проверка GCD == 1 | Отбор несократимых дробей |
| Прямоугольник → квадраты | Стороны прямоугольника | Сторону квадрата | Макс. сторону квадрата |

### Чек-лист для распознавания

**Вопросы для себя при чтении условия:**
1. Есть ли здесь несколько величин, которые должны делиться на одно и то же число?
2. Ищу ли я максимально возможное такое число?

**Если ответ "да" на оба — это GCD.**

### Важное замечание про строки (1071)

Строки ведут себя как числа — их можно "делить" на подстроки без остатка.
- Строка "ABCABC" длины 6 "делится" на "ABC" (длина 3) ровно 2 раза
- Поэтому GCD работает и для длин строк

### Связь с другими концепциями

- **GCD == 1** → числа взаимно простые (coprime), дробь несократима
- **GCD >= 2** → можно разделить на группы размера >= 2
- **LCM (НОК)** → связан через формулу: LCM(a,b) = (a × b) / GCD(a,b)

---

## Two Pointers (Два указателя)

### Когда использовать
- Отсортированный массив и нужно найти пару/тройку
- Палиндром проверка
- Слияние двух отсортированных массивов

### Паттерн
```
left = 0, right = len(arr) - 1
while left < right:
    if condition(arr[left], arr[right]):
        move both
    else:
        move one
```

### Классические задачи
- Valid Palindrome (125)
- Two Sum II (167) — отсортирован
- 3Sum (15)
- Container With Most Water (11)

---

## Sliding Window (Скользящее окно)

### Когда использовать
- Подстроки/подмассивы с определённым свойством
- Максимум/минимум в окне фиксированного размера
- "Самый длинный/короткий подмассив, где..."

### Паттерн
```
left = 0
for right in range(n):
    add arr[right] to window
    while window invalid:
        remove arr[left]
        left += 1
    update answer
```

### Классические задачи
- Longest Substring Without Repeating Characters (3)
- Minimum Window Substring (76)
- Max Consecutive Ones III (1004)

---

## Hash Map / Frequency Counter

### Когда использовать
- Поиск пар/групп по значению
- Подсчёт частот
- O(1) lookup нужен

### Паттерн
```
map := make(map[key]value)
for _, item := range items:
    map[item]++
    // или: if complement, ok := map[target - item]; ok { ... }
```

### Классические задачи
- Two Sum (1)
- Group Anagrams (49)
- Top K Frequent Elements (347)

---

## Fast & Slow Pointers (Floyd's Cycle)

### Когда использовать
- Обнаружение цикла в связном списке
- Нахождение середины списка
- Нахождение начала цикла

### Паттерн
```
slow = head
fast = head
for fast != nil && fast.Next != nil:
    slow = slow.Next      // 1 шаг
    fast = fast.Next.Next // 2 шага
    if slow == fast:
        // cycle detected
```

### Классические задачи
- Linked List Cycle (141)
- Linked List Cycle II (142)
- Find the Duplicate Number (287)

---

## Merge Intervals

### Когда использовать
- Пересечение/объединение интервалов
- Расписания, встречи
- "Свободное время"

### Паттерн
```
sort by start time
current = intervals[0]
for each interval:
    if overlaps(current, interval):
        current.end = max(current.end, interval.end)
    else:
        add current to result
        current = interval
```

### Классические задачи
- Merge Intervals (56)
- Insert Interval (57)
- Meeting Rooms II (253)

---

## BFS (Breadth-First Search)

### Когда использовать
- Кратчайший путь в невзвешенном графе
- Уровневая обработка (деревья)
- "Волновой" алгоритм

### Паттерн
```
queue := []start
visited := set{start}
for len(queue) > 0:
    node := queue[0]
    queue = queue[1:]
    for each neighbor:
        if not visited:
            visited.add(neighbor)
            queue = append(queue, neighbor)
```

### Классические задачи
- Number of Islands (200)
- Word Ladder (127)
- Rotting Oranges (994)

---

## DFS (Depth-First Search)

### Когда использовать
- Обход дерева/графа
- Поиск всех путей
- Backtracking
- "Уйди как можно глубже"

### Паттерн (рекурсивный)
```
func dfs(node, visited):
    if node in visited: return
    visited.add(node)
    for each neighbor:
        dfs(neighbor, visited)
```

### Классические задачи
- Clone Graph (133)
- Course Schedule (207)
- Pacific Atlantic Water Flow (417)

---

## Dynamic Programming (Базовое)

### Когда использовать
- Overlapping subproblems
- Optimal substructure
- "Сколько способов...", "Минимум/максимум..."

### Паттерны

**1. Fibonacci-style (1D)**
```
dp[i] = dp[i-1] + dp[i-2]
// Climbing Stairs, House Robber
```

**2. Grid (2D)**
```
dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + cost[i][j]
// Unique Paths, Minimum Path Sum
```

### Классические задачи
- Climbing Stairs (70)
- House Robber (198)
- Maximum Subarray (53)

---

## Binary Search

### Когда использовать
- Отсортированный массив
- Поиск в "монотонной" функции
- O(log n) обязательно

### Паттерн
```
left, right := 0, len(arr)-1
for left <= right:
    mid := left + (right - left) / 2
    if arr[mid] == target: return mid
    if arr[mid] < target: left = mid + 1
    else: right = mid - 1
```

### Классические задачи
- Binary Search (704)
- Search in Rotated Sorted Array (33)
- Find Minimum in Rotated Sorted Array (153)
- Koko Eating Bananas (875)
