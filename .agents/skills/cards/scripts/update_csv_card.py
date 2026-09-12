#!/usr/bin/env python3
"""
Обновляет карточку в CSV-файле "Программирование Карточки.csv".
Ищет строку по точному совпадению Front (первой колонки) и заменяет Back.
Используется для добавления раздела "Вопросы для собеседования" или полной перезаписи карточки.

Примеры использования:

1. Передать Back аргументом (однострочный):
    python3 update_csv_card.py "Название" "Новое содержимое"

2. Передать Back через stdin (многострочный, рекомендуется):
    python3 update_csv_card.py "Название" - << 'EOF'
    Первая строка
    Вторая строка
    EOF

3. Прочитать Back из файла:
    python3 update_csv_card.py "Название" --file back.txt

Если карточка не найдена — выводит "Card not found" и код возврата 1.
"""

import csv
import os
import sys


def resolve_csv_path() -> str:
    """Вычисляет абсолютный путь к CSV относительно расположения скрипта."""
    script_dir = os.path.dirname(os.path.abspath(__file__))
    # script_dir = .../.serena/skills/cards/scripts
    # поднимаемся на 4 уровня вверх до корня проекта
    project_root = os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(script_dir))))
    return os.path.join(project_root, "Программирование Карточки.csv")


def read_back_from_args_or_stdin(argv: list) -> str:
    if len(argv) < 3:
        raise SystemExit("Usage: python3 update_csv_card.py '<Front>' '<Back>'\n"
                         "       python3 update_csv_card.py '<Front>' -")

    back_source = argv[2]

    if back_source == '-':
        return sys.stdin.read()

    if back_source == '--file':
        if len(argv) < 4:
            raise SystemExit("Usage: python3 update_csv_card.py '<Front>' --file <path>")
        with open(argv[3], encoding='utf-8') as f:
            return f.read()

    return back_source


def update_card(csv_path: str, front: str, new_back: str) -> bool:
    if not os.path.exists(csv_path):
        print(f"CSV file not found: {csv_path}")
        return False

    temp_path = csv_path + ".tmp"
    found = False

    with open(csv_path, newline='', encoding='utf-8') as infile, \
         open(temp_path, 'w', newline='', encoding='utf-8') as outfile:
        reader = csv.reader(infile)
        writer = csv.writer(outfile, quoting=csv.QUOTE_MINIMAL)
        for row in reader:
            if len(row) >= 2 and row[0] == front:
                row[1] = new_back
                found = True
            writer.writerow(row)

    if found:
        os.replace(temp_path, csv_path)
        print("Updated successfully")
        return True
    else:
        os.remove(temp_path)
        print("Card not found")
        return False


if __name__ == "__main__":
    try:
        front_arg = sys.argv[1]
        back_arg = read_back_from_args_or_stdin(sys.argv)
    except IndexError:
        raise SystemExit("Usage: python3 update_csv_card.py '<Front>' '<Back>'\n"
                         "       python3 update_csv_card.py '<Front>' -")

    csv_path = resolve_csv_path()
    success = update_card(csv_path, front_arg, back_arg)
    sys.exit(0 if success else 1)
