# Трекер задач — Algorithms Practice

> Отслеживание текущих задач для набивания руки на базовых паттернах.

---

## 🎯 Цель

Закрепить базовые паттерны до автоматизма:
- Two Pointers (merge, left/right)
- Linear Scan (min/max, consecutive)
- Early returns для edge cases

---

## Задачи на закрепление (Easy)

### Паттерн: Merge двух отсортированных массивов

| # | Статус | Задача | Сложность | Паттерн |
|---|--------|--------|-----------|---------|
| 1 | ✅ | **VK Education: Merge Two Sorted Arrays** | Easy | Two pointers (forward) |
| 2 | ✅ | **LeetCode 88**: Merge Sorted Array | Easy | Two pointers (backward, in-place) |
| 3 | ⬜ | **LeetCode 21**: Merge Two Sorted Lists | Easy | Merge linked lists |
| 4 | ⬜ | **LeetCode 1470**: Shuffle the Array | Easy | Array interleaving |

### Паттерн: Поиск минимума/максимума

| # | Статус | Задача | Сложность | Паттерн |
|---|--------|--------|-----------|---------|
| 1 | ✅ | **VK Education: Rock'n'Run High Score** | Easy | Linear max scan |
| 2 | ⬜ | **LeetCode 1464**: Maximum Product of Two Elements | Easy | Find 2 largest |
| 3 | ⬜ | **LeetCode 414**: Third Maximum Number | Easy | Find 3rd largest (edge cases!) |
| 4 | ⬜ | **LeetCode 485**: Max Consecutive Ones | Easy | Consecutive series |

### Комбинированные

| # | Статус | Задача | Сложность | Паттерн |
|---|--------|--------|-----------|---------|
| 1 | ⬜ | **LeetCode 2206**: Divide Array Into Equal Pairs | Easy | Frequency counter |
| 2 | ⬜ | **LeetCode 2293**: Min Max Game | Easy | Recursive min/max |

---

## Порядок решения

Рекомендуемый порядок (от простого к сложному):

1. ~~VK Education: Merge~~ ✅
2. ~~VK Education: Max~~ ✅
3. LeetCode 88: Merge Sorted Array
4. LeetCode 1464: Maximum Product
5. LeetCode 485: Max Consecutive Ones
6. LeetCode 1470: Shuffle the Array
7. LeetCode 21: Merge Two Sorted Lists
8. LeetCode 414: Third Maximum Number
9. LeetCode 2206: Divide Array Into Equal Pairs
10. LeetCode 2293: Min Max Game

---

## Критерии успеха

- ✅ Решение принято с первого раза
- ✅ TDD: Red → Green → Refactor
- ✅ Ранние выходы для edge cases
- ✅ Время на задачу ≤ 30 мин (включая тесты)
- ✅ Перевод на C++ и сдача в VK All Cups (где применимо)

---

*Создано: 2026-04-11*
