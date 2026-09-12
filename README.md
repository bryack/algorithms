# Algorithms & Data Structures in Go

<p align="left">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go Version" />
  <img src="https://img.shields.io/badge/Practice-78%2B%20Days-orange?style=flat-square&logo=target" alt="78+ Days of Practice" />
  <img src="https://img.shields.io/badge/Tests-Unit%20%26%20TDD-success?style=flat-square&logo=testinglibrary&logoColor=white" alt="Tests" />
  <img src="https://img.shields.io/badge/Target-NeetCode%20150%20%7C%20Yandex-blue?style=flat-square" alt="Target" />
  <img src="https://img.shields.io/badge/License-MIT-lightgrey?style=flat-square" alt="License" />
</p>

Практический репозиторий для глубокого изучения структур данных, классических алгоритмов и подготовки к алгоритмическим секциям технических собеседований (Yandex, BigTech, Tier-1 компании).

Проект построен на принципах **инженерного подхода к алгоритмам**: регулярная ежедневная практика, покрытие тестами (TDD), строгий анализ вычислительной сложности ($O(\text{time})$ и $O(\text{space})$) и фокус на идиоматичном и производительном Go с учетом работы рантайма.

---

## 👩‍💻 Автор

**Анна Нургалеева** — Backend Go Developer  
Инженерный бэкграунд в IT более 10 лет: системный анализ, проектирование распределенных систем, переход в разработку на Go с акцентом на архитектуру, надежность и тестирование.

- **Telegram**: [@bryacka](https://t.me/bryacka)
- **GitHub**: [github.com/bryack](https://github.com/bryack)
- **LinkedIn**: [linkedin.com/in/anna-nurgaleeva](https://www.linkedin.com/in/anna-nurgaleeva-ba9a6338/)
- **Email**: [bryacka@gmail.com](mailto:bryacka@gmail.com)

---

## 🎯 Инженерная концепция и принципы

В отличие от простого «прорешивания задач», здесь алгоритмы рассматриваются как **проектирование вычислительных процессов**:

1. **Test-Driven Development (TDD):**
   Каждая задача сопровождается изолированным набором модульных тестов (`main_test.go`). Тест-кейсы покрывают как базовые сценарии, так и граничные условия: пустые входные данные, один элемент, дубликаты, переполнение диапазонов, крайние значения.
2. **Анализ сложности:**
   В коде решений явно фиксируется асимптотическая оценка временной и пространственной сложности:
   `// Time: O(n log n), Space: O(1)`.
3. **Осознанность в отношении Go Runtime & Memory:**
   - **Слайсы vs Массивы:** учет переаллокаций базового массива при росте `capacity`, использование pre-allocation (`make([]T, 0, n)`).
   - **Строки и Unicode:** корректная дифференциация байт (`byte`) и рун (`rune`), защита от некорректного слайсинга многобайтовых UTF-8 символов.
   - **Множества (Set):** реализация через `map[T]struct{}` (пустая структура не аллоцирует память в куче).
   - **Дженерики (Type Parameters):** создание переиспользуемых обобщенных структур данных (`generics/`).
   - **Низкоуровневые оптимизации:** профилирование, бенчмаркинг и эксперименты с Go Assembly (`archive/day_50/asm`).

---

## 📂 Структура репозитория

```text
.
├── day_72/ ... day_78/     # Актуальные ежедневные задачи с тестами
├── archive/                # Архив решений (Day 1 – Day 71: свыше 70 дней практики)
│   ├── day_10/ ... day_14/ # Задачи с контестов (VK, олимпиады)
│   ├── day_16/             # Стек, очередь, дек (теория + реализация)
│   └── day_50/asm/         # Go Assembly (AMD64) + бенчмарки
├── generics/               # Обобщенные структуры данных и утилиты
│   ├── Generic_Cache/      # Потокобезопасный generic-кэш
│   ├── Filter/             # Функциональный фильтр слайсов
│   ├── RemoveDuplicates/   # Дедупликация срезов на дженериках
│   └── SumMapValues/       # Агрегация числовых мап
├── nog/                    # Оптимизация работы с памятью и строками
├── cmd/sync/               # CLI-утилита синхронизации Obsidian <-> catalog.yaml <-> планы
│   └── README.md           # Документация по утилите синхронизации
├── .agents/                # Материалы менторинга, карточки задач и разборы
│   └── skills/
│       ├── algorithms-mentor/  # Дорожная карта и паттерны
│       └── yandex-interview/   # Шаблоны и подборка задач собеседований Яндекса
├── NEETCODE_PLAN_1..7.md   # Детальные трекеры тем NeetCode 150
├── Plan. Introduction to Algorithms.md # Базовый 6-месячный план подготовки
└── catalog.yaml            # Структурированный каталог всех решенных задач
```

---

## 🧠 Изучаемые паттерны и темы

| Категория | Паттерны и структуры данных | Примеры задач |
| :--- | :--- | :--- |
| **Arrays & Hashing** | Хэш-таблицы, Prefix Sums, Frequency Count | Two Sum, Group Anagrams, Top K Frequent |
| **Two Pointers** | Встречные указатели, быстрый/медленный указатель | 3Sum, Container With Most Water, Trapping Rain Water |
| **Sliding Window** | Фиксированное и динамическое окно | Longest Substring Without Repeating Characters |
| **Stack & Queue** | Монотонический стек, MinStack, очереди на слайсах | Valid Parentheses, Daily Temperatures |
| **Binary Search** | Поиск в отсортированном массиве, поиск по ответу | Binary Search, Search in Rotated Sorted Array |
| **Linked Lists** | Разворот, слияние, обнаружение циклов (Флойд) | Reverse Linked List, Merge Two Sorted Lists |
| **Trees & Graphs** | DFS, BFS, свойства BST, топологическая сортировка | Invert Tree, Validate BST, Number of Islands |
| **Advanced Go** | Generics, Go Assembly, Memory Layout, Runes | Generic Cache, SIMD/ASM Sum, UTF-8 Indexing |

---

## 🛠 Запуск тестов и бенчмарков

Проект не требует внешних зависимостей — используется стандартный инструментарий Go toolchain.

### Запуск всех тестов проекта
```bash
go test ./...
```

### Запуск тестов конкретного дня с подробным выводом
```bash
go test -v ./day_78/...
```

### Запуск тестов с проверкой покрытия кода (Coverage)
```bash
go test -cover ./...
```

### Запуск бенчмарков
```bash
go test -bench=. -benchmem ./generics/...
```

---

## 🔄 Автоматизация и синхронизация

В проект встроена CLI-утилита **`cmd/sync`**, которая синхронизирует решенные задачи из **Obsidian Vault** с каталогом [catalog.yaml](file:///home/bryack/Documents/algorithms/catalog.yaml) и автоматически генерирует дорожные карты `NEETCODE_PLAN_*.md`:

```bash
# Синхронизация текущего прогресса:
rtk go run ./cmd/sync

# Автоматический импорт новых заметок из Obsidian в каталог:
rtk go run ./cmd/sync -import-orphans
```

Подробное руководство по конфигурации, структуре заметок и флагам утилиты смотрите в [cmd/sync/README.md](file:///home/bryack/Documents/algorithms/cmd/sync/README.md).

---

## 📈 Планы и трекеры подготовки

В репозитории зафиксированы подробные трекеры и чек-листы:
- [Plan. Introduction to Algorithms.md](file:///home/bryack/Documents/algorithms/Plan.%20Introduction%20to%20Algorithms.md) — 6-месячная стратегия и математический фундамент.
- [NEETCODE_PLAN_1.md](file:///home/bryack/Documents/algorithms/NEETCODE_PLAN_1.md) – [NEETCODE_PLAN_7.md](file:///home/bryack/Documents/algorithms/NEETCODE_PLAN_7.md) — дорожные карты по классическим алгоритмическим темам NeetCode 150.
- [catalog.yaml](file:///home/bryack/Documents/algorithms/catalog.yaml) — единый реестр решенных задач с метаданными.
