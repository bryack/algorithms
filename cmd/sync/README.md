# Sync — CLI-утилита синхронизации прогресса алгоритмов

<p align="left">
  <img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go Version" />
  <img src="https://img.shields.io/badge/Tool-CLI-blueviolet?style=flat-square" alt="CLI Tool" />
  <img src="https://img.shields.io/badge/Tests-TDD%20Passed-success?style=flat-square" alt="Tests" />
</p>

Утилита автоматизирует ведение учебного процесса: связывает заметки по задачам в **Obsidian Vault** со структурированным каталогом **`catalog.yaml`** и генерирует наглядные таблицы планов **`NEETCODE_PLAN_*.md`**.

---

## 📐 Архитектура синхронизации

```
  ┌─────────────────────────────────┐
  │         Obsidian Vault          │  ◄── Динамический прогресс
  │  (заметки по задачам в Markdown)│      (даты, retention L1-L5, needRepeat)
  └────────────────┬────────────────┘
                   │  чтение frontmatter
                   ▼
  ┌─────────────────────────────────┐
  │        cmd/sync (Go CLI)        │
  └────────┬────────────────┬───────┘
           │                │
  чтение / запись      генерация Markdown
           ▼                ▼
┌──────────────────┐  ┌─────────────────────────────────┐
│   catalog.yaml   │  │       NEETCODE_PLAN_*.md        │ ◄── Витрина для Git
│ (реестр и схемы) │  │  (таблицы роадмапа по темам)    │
└──────────────────┘  └─────────────────────────────────┘
```

1. **Obsidian Vault** — единственный источник правды о **прогрессе**: сколько раз решена задача, когда, с какой оценкой качества (retention) и нужен ли повтор.
2. **`catalog.yaml`** — единственный источник правды о **структуре**: список блоков, соответствие папкам в Obsidian, выходные файлы планов, метаданные задач (сложность, паттерн, группа).
3. **`NEETCODE_PLAN_*.md`** — генерируемые файлы витрины для GitHub. **Не редактируются вручную**.

---

## 📝 Требования к заметкам в Obsidian

Чтобы утилита распознала файл как задачу по алгоритмам, должны выполняться условия:

### 1. Формат имени файла
Имя файла должно начинаться с номера задачи LeetCode с точкой и пробелом:
- `217. Contains Duplicate.md`
- `98. Validate Binary Search Tree.md`
- `1. Two Sum.md`

### 2. Расположение файла
Заметка должна находиться внутри папки соответствующего блока в хранилище (или в её подпапках):
- `Arrays_Hashing/` $\rightarrow$ блок `arrays_hashing`
- `Binary Tree/` или `Binary Tree/Binary Search Tree/` $\rightarrow$ блок `trees`
- `Two_pointers/` $\rightarrow$ блок `two_pointers`
- `Linked List/` $\rightarrow$ блок `linked_list`
- `Stack/` $\rightarrow$ блок `stack`
- `SQL/` $\rightarrow$ блок `sql`

### 3. YAML Frontmatter
В начале заметки обязательно должен присутствовать блок параметров между `---`:

```yaml
---
Type:
  - exercise
Project:
  - go
tags:
  - algorithms
  - leetcode
created: 2026-08-26
algorithms:
  - array
  - hash_table
solved_dates:
  - 2026-08-26
  - 2026-08-27
needRepeat: false
retention: "4"
difficulty: Easy
---
```

**Ключевые поля:**
- `Type: exercise` и `tags: [..., algorithms, ...]` — фильтр для отбора заметок-задач.
- `difficulty` — сложность (`Easy`, `Medium`, `Hard`). Используется при импорте.
- `algorithms` — список паттернов/тегов алгоритма (подставляется в колонку «Тема / Паттерн»).
- `solved_dates` — массив дат решений (определяет количество решений задачи и дату первого решения).
- `retention` — оценка качества усвоения от `1` до `5` (или `L1`–`L5`). Отображается эмодзи:
  - `🟢 L4` / `⚪ L5` — отлично / эксперт
  - `🟡 L3` — решено, но с шероховатостями
  - `🔴 L1` / `🟠 L2` — требовались подсказки
- `needRepeat` — булевый флаг (`true` / `false`), сигнализирующий о необходимости интервального повторения.

---

## ⚙️ Структура `catalog.yaml`

Файл каталога описывает конфигурацию блоков:

```yaml
version: 1
vault_root: /home/bryack/Documents/Obsidian/bryack/014_Go/algorithms
blocks:
  - id: trees
    title: Trees
    folder: Binary Tree          # Имя корневой папки в Obsidian
    plan_file: NEETCODE_PLAN_7.md # Файл, куда рендерится план
    sections:                    # Секции внутри плана
      - group: main
        title: Основной блок
    problems:
      - leetcode_id: 100
        title: Same Tree
        difficulty: Easy
        pattern: recursion
        group: main
```

---

## 🚀 Команды и флаги запуска

Все команды рекомендуется запускать из **корня репозитория** с префиксом `rtk `:

### 1. Обычная синхронизация прогресса
Считывает текущее состояние Obsidian, сопоставляет с `catalog.yaml` и обновляет все файлы `NEETCODE_PLAN_*.md`:
```bash
rtk go run ./cmd/sync
```

### 2. Импорт новых (Orphan) заметок в каталог
Если вы создали и решили новые заметки в Obsidian, которых ещё нет в `catalog.yaml`, флаг `-import-orphans` автоматически извлечет метаданные и добавит задачи в соответствующие блоки каталога в секцию `main`:
```bash
rtk go run ./cmd/sync -import-orphans
```

### 3. Предпросмотр изменений (Dry Run)
Показывает, какие файлы планов и каталога изменятся, **не записывая** ничего на диск:
```bash
rtk go run ./cmd/sync -dry-run
rtk go run ./cmd/sync -import-orphans -dry-run
```

### 4. Только диагностика (Check)
Выводит в терминал предупреждения (заметки без номеров, дубликаты, orphan-заметки) и статистику по блокам без перезаписи файлов:
```bash
rtk go run ./cmd/sync -check
```

### 5. Первичная инициализация каталога из существующих планов
Используется только при первоначальной настройке для обратного парсинга `NEETCODE_PLAN_*.md` в `catalog.yaml`:
```bash
rtk go run ./cmd/sync -init -plans-dir . -catalog catalog.yaml
```

---

## 📦 Сборка исполняемого файла

Если вы хотите запускать утилиту как скомпилированный бинарник без задержки компиляции `go run`:

```bash
# Сборка бинарника в корень репозитория:
rtk go build -o sync ./cmd/sync

# Быстрый запуск:
./sync
./sync -import-orphans
```

*(Файл `/sync` в корне репозитория уже прописан в `.gitignore` и не попадет в коммиты).*

---

## 🧪 Тестирование

Утилита разработана по методологии **TDD** и протестирована модульными и сквозными тестами:

```bash
# Запуск всех тестов утилиты:
rtk go test -v ./cmd/sync/...

# Проверка тестового покрытия:
rtk go test -cover ./cmd/sync/...
```
