import pandas as pd
import json
import re

cols = [0, 1, 3]
column_names = ['Преподаватель', 'Группа', 'Предмет']

sip_data = pd.read_excel('./Сип.xlsx', usecols=cols, header=2, names=column_names)

teachers = {}
groups = {}
subjects = {}
cabinets = []
specializations = {}
specialization_id_counter = 1
subject_id_counter = 1
group_id_counter = 1
cabinet_id_counter = 1
group_to_id = {}
subject_to_id = {}
specialization_to_id = {}

def extract_group_info(group_name):
    match = re.match(r'(\d{1,3})-?([А-Я]{1,4})-?(\d{2})', group_name)
    if match:
        course = int(match.group(1)[0])  # Первый символ номера группы указывает на курс
        specialization_code = match.group(2)  # Код направления (например, ИС)
        return course, specialization_code
    return None, None

for _, row in sip_data.iterrows():
    teacher = row['Преподаватель']
    group = row['Группа']
    subject = row['Предмет']

    # Обработка предметов
    if subject not in subject_to_id:
        subject_id = subject_id_counter
        subjects[subject_id] = {"name": subject, "type": "Laboratory"}
        subject_to_id[subject] = subject_id_counter
        subject_id_counter += 1
    else:
        subject_id = subject_to_id[subject]

    # Извлечение информации о группе и специализации
    course, specialization_code = extract_group_info(group)
    if course is None or specialization_code is None:
        continue  # Пропустить, если группа не соответствует ожидаемому формату

    if specialization_code not in specialization_to_id:
        specialization_to_id[specialization_code] = specialization_id_counter
        specializations[specialization_id_counter] = {
            "name": specialization_code,
            "course": course,
            "short_plan": {}
        }
        specialization_id_counter += 1
    
    specialization_id = specialization_to_id[specialization_code]
    
    if group not in group_to_id:
        groups[group_id_counter] = {
            "specialization": {"id": specialization_id},
            "name": group,  # Здесь сохраняем полное название группы, например "303П-22"
            "max_pairs": 18
        }
        group_to_id[group] = group_id_counter
        group_id_counter += 1

    # Обработка преподавателей
    if teacher not in teachers:
        teachers[teacher] = {"name": teacher, "capacity": 18, "links": {}}
    
    group_id = group_to_id[group]
    if group_id not in teachers[teacher]["links"]:
        teachers[teacher]["links"][group_id] = []
    
    if subject_id not in teachers[teacher]["links"][group_id]:
        teachers[teacher]["links"][group_id].append(subject_id)

    # Обновление short_plan для специализаций
    if subject_id not in specializations[specialization_id]["short_plan"]:
        specializations[specialization_id]["short_plan"][str(subject_id)] = subject_id

# Генерация кабинетов
total_groups = len(groups)
total_cabinets = total_groups // 2  # Количество кабинетов меньше количества групп
for _ in range(total_cabinets):
    cabinets.append({
        "name": str(cabinet_id_counter),
        "type": "Normal"
    })
    cabinet_id_counter += 1

# Сохранение данных в JSON файлы
with open('subjects.json', 'w', encoding='utf-8') as f:
    json.dump(list(subjects.values()), f, ensure_ascii=False, indent=4)

with open('cabinets.json', 'w', encoding='utf-8') as f:
    json.dump(cabinets, f, ensure_ascii=False, indent=4)

with open('specializations.json', 'w', encoding='utf-8') as f:
    json.dump(list(specializations.values()), f, ensure_ascii=False, indent=4)

with open('groups.json', 'w', encoding='utf-8') as f:
    json.dump(list(groups.values()), f, ensure_ascii=False, indent=4)

with open('teachers.json', 'w', encoding='utf-8') as f:
    json.dump(list(teachers.values()), f, ensure_ascii=False, indent=4)
