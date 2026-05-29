# Workflow: Управление сессиями

> **Последнее обновление: 2026-05-27**

## Файлы

- `.session_start` — timestamps начала/окончания сессий, перерывов и активностей
- `journal.csv` — журнал завершённых задач, лекций, карточек и других активностей

---

## Единый формат записей в `.session_start`

**ВСЕ записи в файле должны следовать строго одному формату:**

```
YYYY-MM-DD HH:MM:SS # MARKER: описание
```

### Правила формата

1. **Дата и время** — всегда в начале строки, формат `YYYY-MM-DD HH:MM:SS`
2. **Разделитель** — пробел + `#` + пробел
3. **MARKER** — заглавными буквами, без пробелов, с двоеточием в конце:
   - `SESSION_START:` — начало учебной сессии
   - `SESSION_END:` — окончание учебной сессии
   - `BREAK_START:` — начало перерыва
   - `BREAK_END:` — окончание перерыва
   - `ACTIVITY_START:` — начало конкретной активности (задача, лекция, карточки, видео)
   - `ACTIVITY_END:` — окончание конкретной активности
4. **Описание** — произвольный текст после двоеточия и пробела. Для `ACTIVITY_START/END` описание должно совпадать, чтобы скрипты могли найти пару

### Примеры правильных записей

```
2026-05-18 08:48:03 # SESSION_START
2026-05-18 11:25:00 # BREAK_START
2026-05-18 11:58:00 # BREAK_END
2026-05-18 16:07:37 # ACTIVITY_START: Nature of Go - xv6 SYS_va2pa
2026-05-18 16:22:49 # ACTIVITY_END: Nature of Go - xv6 SYS_va2pa
2026-05-18 17:03:32 # ACTIVITY_START: LeetCode 643: Maximum Average Subarray I
2026-05-18 17:28:34 # ACTIVITY_END: LeetCode 643: Maximum Average Subarray I
2026-05-18 20:30:26 # SESSION_END
```

### Критические запреты

- ❌ НЕ использовать альтернативные форматы (`ACTIVITY_START: 2026-05-18 16:07:37 | Описание`)
- ❌ НЕ добавлять пустые строки между записями
- ❌ НЕ использовать `>` (перезапись), только `>>` (добавление)
- ❌ НЕ создавать записи без даты-времени в начале строки

---

## Команды

### Начало сессии

**Триггер:** "Начинаем новую сессию", "Старт сессии", "Начинаем заниматься", "Сегодня [дата]"

```bash
echo "$(date '+%Y-%m-%d %H:%M:%S') # SESSION_START" >> .session_start
```

**Ответ:** "✅ Сессия зафиксирована: [время]"

### Окончание сессии

**Триггер:** "Заканчиваем на сегодня", "Всё, хватит"

```bash
echo "$(date '+%Y-%m-%d %H:%M:%S') # SESSION_END" >> .session_start
```

### Перерыв

**Начало перерыва** (триггер: "сейчас перерыв", "перерыв", "отдых"):
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S') # BREAK_START" >> .session_start
```

**Ответ:** "✅ Перерыв с [время]"

**Окончание перерыва** (триггер: "перерыв окончен", "вернулась", "продолжаем"):
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S') # BREAK_END" >> .session_start
```

**Ответ:** "✅ Перерыв окончен в [время]. Длительность: [N] минут."

### Активность

**Начало активности** (триггер: "Отметь начало активности", "Начинаю решать задачу", "Слушаю лекцию"):
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S') # ACTIVITY_START: [название активности]" >> .session_start
```

**Окончание активности** (триггер: "Завершила задачу", "Занеси прогресс", "Отметь окончание"):
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S') # ACTIVITY_END: [название активности]" >> .session_start
```

**Важно:** название активности в `ACTIVITY_START` и `ACTIVITY_END` должно совпадать символ в символ, иначе скрипт не найдёт пару.

### Запись в `journal.csv`

**Триггер:** "Запиши в журнал", "Я завершила задачу", "Занеси в журнал", "Обнови журнал"

**Формат записи:**
```csv
Дата,Направление,Тема,Что делала,Артефакты,Сложность,Время,Комментарий
```

**Колонки:**
- **Дата** — `YYYY-MM-DD`
- **Направление** — `Алгоритмы`, `Go`, `Карточки`, `Математика`, `Собеседование`, `SQL` и т.д.
- **Тема** — название задачи, лекции, темы карточек
- **Что делала** — краткое описание (1–2 предложения)
- **Артефакты** — файлы, ссылки (например, `day_42/main.go`, `Telegram-канал`)
- **Сложность** — 1–10 или Easy/Medium/Hard
- **Время** — в часах (например, `1.5ч`, `25мин`)
- **Комментарий** — дополнительные заметки (retention, баги, инсайты)

---

## Bash-скрипты для расчёта времени

### Скрипт 1: время конкретной активности

```bash
TASK_NAME="LeetCode 643: Maximum Average Subarray I"
SESSION_FILE=".session_start"

# Находим последнее начало активности по точному названию
start_time=$(grep "# ACTIVITY_START:.*${TASK_NAME}" "$SESSION_FILE" | tail -1 | awk '{print $1" "$2}')
# Находим последнее окончание активности по точному названию
end_time=$(grep "# ACTIVITY_END:.*${TASK_NAME}" "$SESSION_FILE" | tail -1 | awk '{print $1" "$2}')

# Если активность ещё не завершена, берём текущее время
[ -z "$end_time" ] && end_time=$(date '+%Y-%m-%d %H:%M:%S')

start_epoch=$(date -d "$start_time" +%s)
end_epoch=$(date -d "$end_time" +%s)

# Вычитаем перерывы, попавшие в интервал активности
total_break=0
in_break=false
break_start=0

while IFS= read -r line; do
    timestamp=$(echo "$line" | awk '{print $1" "$2}')
    epoch=$(date -d "$timestamp" +%s)
    [ "$epoch" -lt "$start_epoch" ] || [ "$epoch" -gt "$end_epoch" ] && continue
    
    if echo "$line" | grep -q "# BREAK_START"; then
        break_start=$epoch; in_break=true
    elif echo "$line" | grep -q "# BREAK_END" && [ "$in_break" = true ]; then
        total_break=$((total_break + epoch - break_start))
        in_break=false
    fi
done < "$SESSION_FILE"

[ "$in_break" = true ] && total_break=$((total_break + end_epoch - break_start))

pure_seconds=$((end_epoch - start_epoch - total_break))
pure_hours=$(awk "BEGIN {printf \"%.1f\", $pure_seconds / 3600}")

echo "Чистое время: ${pure_hours}ч"
echo "(Брутто: $(( (end_epoch - start_epoch) / 60 )) мин, Перерывы: $(( total_break / 60 )) мин)"
```

### Скрипт 2: время за весь день

```bash
SESSION_FILE=".session_start"

# Находим последнюю сессию
session_line=$(grep -n "# SESSION_START" "$SESSION_FILE" | tail -1 | cut -d: -f1)
session_start=$(sed -n "${session_line}p" "$SESSION_FILE" | awk '{print $1" "$2}')
session_epoch=$(date -d "$session_start" +%s)

# Конец — либо SESSION_END, либо последняя запись, либо текущее время
end_time=$(tail -1 "$SESSION_FILE" | awk '{print $1" "$2}')
if tail -1 "$SESSION_FILE" | grep -q "# BREAK_START"; then
    end_time=$(date '+%Y-%m-%d %H:%M:%S')
fi
end_epoch=$(date -d "$end_time" +%s)

# Вычитаем все перерывы внутри сессии
total_break=0
in_break=false
break_epoch=0

while IFS= read -r line; do
    if echo "$line" | grep -q "# BREAK_START"; then
        ts=$(echo "$line" | awk '{print $1" "$2}')
        break_epoch=$(date -d "$ts" +%s)
        in_break=true
    elif echo "$line" | grep -q "# BREAK_END" && [ "$in_break" = true ]; then
        ts=$(echo "$line" | awk '{print $1" "$2}')
        end_break_epoch=$(date -d "$ts" +%s)
        total_break=$((total_break + end_break_epoch - break_epoch))
        in_break=false
    fi
done < <(tail -n +"$session_line" "$SESSION_FILE")

[ "$in_break" = true ] && total_break=$((total_break + end_epoch - break_epoch))

total_seconds=$((end_epoch - session_epoch))
active_seconds=$((total_seconds - total_break))

total_hours=$(awk "BEGIN {printf \"%.1f\", $total_seconds / 3600}")
active_hours=$(awk "BEGIN {printf \"%.1f\", $active_seconds / 3600}")
break_hours=$(awk "BEGIN {printf \"%.1f\", $total_break / 3600}")

echo "Сессия: $session_start → $end_time"
echo "Общее время:    ${total_hours}ч ($((total_seconds / 60)) мин)"
echo "Перерывы:       ${break_hours}ч ($((total_break / 60)) мин)"
echo "Активное время: ${active_hours}ч ($((active_seconds / 60)) мин)"
```

---

## Пример диалога

```
Пользователь: Начинаем новую сессию
Я: ✅ Сессия зафиксирована: 2026-05-18 08:48:03

Пользователь: Сейчас перерыв
Я: ✅ Перерыв с 11:25:00

Пользователь: Перерыв окончен
Я: ✅ Перерыв окончен в 11:58:00. Длительность: 33 минуты.

Пользователь: Отметь начало активности - буду слушать лекцию Nature of Go
Я: ✅ Активность зафиксирована: 2026-05-18 16:07:37

Пользователь: Завершила задачу, занеси прогресс
Я: ✅ Активность завершена: 2026-05-18 16:22:49
   ✅ Запись добавлена в journal.csv
```

---

## Быстрый reference

| Ситуация | Команда |
|----------|---------|
| Начало сессии | `echo "$(date '+%Y-%m-%d %H:%M:%S') # SESSION_START" >> .session_start` |
| Конец сессии | `echo "$(date '+%Y-%m-%d %H:%M:%S') # SESSION_END" >> .session_start` |
| Перерыв | `echo "$(date '+%Y-%m-%d %H:%M:%S') # BREAK_START" >> .session_start` |
| Конец перерыва | `echo "$(date '+%Y-%m-%d %H:%M:%S') # BREAK_END" >> .session_start` |
| Начало активности | `echo "$(date '+%Y-%m-%d %H:%M:%S') # ACTIVITY_START: [название]" >> .session_start` |
| Окончание активности | `echo "$(date '+%Y-%m-%d %H:%M:%S') # ACTIVITY_END: [название]" >> .session_start` |
| Расчёт активности | скрипт из раздела "Bash-скрипты" выше |
| Расчёт дня | скрипт из раздела "Bash-скрипты" выше |
