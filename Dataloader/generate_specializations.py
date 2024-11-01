import pandas as pd
import json
import re

# Чтение данных из Excel файла
file_path = 'Сип.xlsx'
data = pd.read_excel(file_path, header=None)

# Инициализация массива специализаций
specializations = []
specialization_id = 1

# Пример массива предметов, предполагая что у нас есть массив subjects с id и названиями
# Для демонстрации возьмем фиксированные предметы
subjects = [
    {"id": 1, "name": "Математика"},
    {"id": 2, "name": "Информационные технологии в профессиональной деятельности"},
    {"id": 3, "name": "Физика"}
]

# Создаем маппинг предметов для быстрого доступа к ID по имени
subject_mapping = {subject["name"]: subject["id"] for subject in subjects}

# Функция для определения курса из номера группы
def get_course(group_number):
    return group_number // 100

# Обработка данных
for index, row in data.iterrows():
    if index < 2:  # Пропускаем первые две строки
        continue
    
    group_name = row[1]  # Столбец B - Номер группы
    subject_name = row[3]  # Столбец D - Название предмета
    first_semester = row[4]  # Столбец E - 1 полугодие
    second_semester = row[18]  # Столбец S - 2 полугодие
    
    # Извлечение факультета, курса и номера группы из формата
    match = re.match(r'(\d+)-?([А-Я]+)-?(\d+)', group_name)
    if match:
        group_number = int(match.group(1))
        faculty = match.group(2)
        year = int(match.group(3))

        course = get_course(group_number)
        
        # Проверяем, есть ли такая специализация
        specialization = next((spec for spec in specializations if spec["name"] == faculty and spec["course"] == course), None)
        
        if not specialization:
            specialization = {
                "id": specialization_id,
                "name": faculty,
                "course": course,
                "short_plan": {}
            }
            specializations.append(specialization)
            specialization_id += 1
        
        # Добавляем предметы в план по полугодиям
        semester_plan = specialization["short_plan"]
        
        if pd.notna(subject_name):
            subject_id = subject_mapping.get(subject_name)
            if pd.notna(first_semester):
                semester_plan["1"] = semester_plan.get("1", []) + [subject_id]
            if pd.notna(second_semester):
                semester_plan["2"] = semester_plan.get("2", []) + [subject_id]

# Преобразование в JSON
specializations_json = json.dumps(specializations, ensure_ascii=False, indent=4)

# Запись результата в файл
with open('specializations.json', 'w', encoding='utf-8') as f:
    f.write(specializations_json)

print("Массив специализаций успешно сохранён в specializations.json")
