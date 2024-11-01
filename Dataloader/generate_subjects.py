import pandas as pd
import json

# Чтение данных из Excel файла
file_path = 'Сип.xlsx'
data = pd.read_excel(file_path, header=None)

# Инициализация массива subjects
subjects = []
subject_id = 1

# Пропускаем первые две строки
for index, row in data.iterrows():
    if index < 2:  # Пропускаем первые две строки
        continue
    
    teacher_name = row[0]  # Столбец A - ФИО преподавателя
    group_name = row[1]     # Столбец B - Номер группы
    subject_name = row[3]   # Столбец D - Название предмета
    first_semester = row[4] # Столбец E - 1 полугодие (число пар)
    second_semester = row[18]  # Столбец S - 2 полугодие (число пар)

    # Проверяем, есть ли название предмета
    if pd.notna(subject_name) and subject_name not in [s['name'] for s in subjects]:
        subjects.append({
            "id": subject_id,
            "name": subject_name,
            "type": "Normal",  # Задаём тип кабинета по умолчанию
            "capacity_1": 0 if pd.isna(first_semester) else int(first_semester + 0.5),
            "capacity_2": 0 if pd.isna(second_semester) else int(second_semester + 0.5)
        })
        subject_id += 1

# Преобразование в JSON
subjects_json = json.dumps(subjects, ensure_ascii=False, indent=4)

# Запись результата в файл
with open('subjects.json', 'w', encoding='utf-8') as f:
    f.write(subjects_json)

print("Массив предметов успешно сохранён в subjects.json")
