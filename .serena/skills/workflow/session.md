# Workflow: Управление сессиями

## Файлы

- `.session_start` — timestamps начала сессий, перерывов и задач
- `journal.csv` — журнал решённых задач с временем

---

## Команды

### Начало сессии

**Триггер**: "Начинаем новую сессию", "Старт сессии", "Начинаем заниматься", "Сегодня [дата]"

```bash
date "+%Y-%m-%d %H:%M:%S  # SESSION_START" >> .session_start
```

**Ответ:** "✅ Сессия зафиксирована: [время]"

### Конец сессии

**Триггер**: "Заканчиваем на сегодня", "Всё, хватит"

```bash
echo "$(date '+%Y-%m-%d %H:%M:%S')  # SESSION_END" >> .session_start
```

### Перерыв

**Начало перерыва** (триггер: "сейчас перерыв", "перерыв", "отдых"):
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S')  # BREAK_START" >> .session_start
```

**Ответ:** "✅ Перерыв с [время]"

**Окончание перерыва** (триггер: "перерыв окончен", "вернулась", "продолжаем"):
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S')  # BREAK_END" >> .session_start
```

**Ответ:** "✅ Перерыв окончен в [время]. Длительность: [N] минут."

### Запись в journal.csv

**Триггер**: "Запиши в журнал", "Я завершила задачу", "Занеси в журнал", "Обнови журнал"

1. Найти время начала задачи в `.session_start`:
```bash
grep "JOURNAL:.*[название задачи].*(start)" .session_start | tail -1
```

2. Зафиксировать завершение:
```bash
echo "$(date '+%Y-%m-%d %H:%M:%S')  # JOURNAL: [название] (завершена)" >> .session_start
```

3. Формат записи в `journal.csv`:
```csv
[дата],[направление],[тема],[что делала],[артефакты],[сложность],[время],[комментарий]
```

---

## Правила форматирования `.session_start`

Чтобы скрипты автоматического расчёта работали корректно:

1. **Только `>>` (добавление), никогда `>` (перезапись)**  
   Перезапись стирает историю и ломает подсчёт времени.

2. **Единый формат записей**  
   ```
   YYYY-MM-DD HH:MM:SS  # MARKER: описание
   ```

3. **JOURNAL — всегда пара: `(start)` и `(завершена)`**  
   Иначе grep-скрипт не найдёт записи.

4. **BREAK_START и BREAK_END — строго парные**  
   Нельзя начать новый перерыв, не закончив предыдущий.

5. **Без пустых строк**  
   Пустые строки ломают `while read` циклы.

---

## Bash-скрипты для расчёта времени

### Скрипт 1: время конкретной задачи

```bash
TASK_NAME="Название задачи"
SESSION_FILE=".session_start"

start_time=$(grep "JOURNAL:.*${TASK_NAME}.*(start)" "$SESSION_FILE" | tail -1 | awk '{print $1" "$2}')
end_time=$(grep "JOURNAL:.*${TASK_NAME}.*(завершена)" "$SESSION_FILE" | tail -1 | awk '{print $1" "$2}')
[ -z "$end_time" ] && end_time=$(date '+%Y-%m-%d %H:%M:%S')

start_epoch=$(date -d "$start_time" +%s)
end_epoch=$(date -d "$end_time" +%s)

total_break=0
in_break=false
break_start=0

while IFS= read -r line; do
    timestamp=$(echo "$line" | awk '{print $1" "$2}')
    epoch=$(date -d "$timestamp" +%s)
    [ "$epoch" -lt "$start_epoch" ] || [ "$epoch" -gt "$end_epoch" ] && continue
    
    if echo "$line" | grep -q "BREAK_START"; then
        break_start=$epoch; in_break=true
    elif echo "$line" | grep -q "BREAK_END" && [ "$in_break" = true ]; then
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

session_line=$(grep -n "SESSION_START" "$SESSION_FILE" | tail -1 | cut -d: -f1)
session_start=$(sed -n "${session_line}p" "$SESSION_FILE" | awk '{print $1" "$2}')
session_epoch=$(date -d "$session_start" +%s)

end_time=$(tail -1 "$SESSION_FILE" | awk '{print $1" "$2}')
if tail -1 "$SESSION_FILE" | grep -q "BREAK_START"; then
    end_time=$(date '+%Y-%m-%d %H:%M:%S')
fi
end_epoch=$(date -d "$end_time" +%s)

total_break=0
in_break=false
break_epoch=0

while IFS= read -r line; do
    if echo "$line" | grep -q "BREAK_START"; then
        ts=$(echo "$line" | awk '{print $1" "$2}')
        break_epoch=$(date -d "$ts" +%s)
        in_break=true
    elif echo "$line" | grep -q "BREAK_END" && [ "$in_break" = true ]; then
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
Я: ✅ Сессия зафиксирована: 2026-04-04 15:41:35

Пользователь: Сейчас перерыв
Я: ✅ Перерыв с 16:30

Пользователь: Перерыв окончен
Я: ✅ Перерыв окончен в 16:45. Длительность: 15 минут.

Пользователь: Я завершила задачу
Я: ✅ Задача заняла ~2.5 часа (авторасчёт из .session_start)
   ✅ Запись добавлена в journal.csv
```

---

## Быстрый reference

| Ситуация | Команда |
|----------|---------|
| Начало сессии | `date "+%Y-%m-%d %H:%M:%S # SESSION_START" >> .session_start` |
| Конец сессии | `echo "$(date '+%Y-%m-%d %H:%M:%S') # SESSION_END" >> .session_start` |
| Перерыв | `echo "$(date '+%Y-%m-%d %H:%M:%S') # BREAK_START" >> .session_start` |
| Конец перерыва | `echo "$(date '+%Y-%m-%d %H:%M:%S') # BREAK_END" >> .session_start` |
| Журнал (start) | `echo "$(date '+%Y-%m-%d %H:%M:%S') # JOURNAL: [название] (start)" >> .session_start` |
| Журнал (end) | `echo "$(date '+%Y-%m-%d %H:%M:%S') # JOURNAL: [название] (завершена)" >> .session_start` |
| Расчёт задачи | `bash calc_task.sh` |
| Расчёт дня | `bash calc_day.sh` |
