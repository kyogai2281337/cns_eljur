import json

# Инициализация массива кабинетов
cabinets = []

io = 1
# Генерация кабинетов
for i in range(20, 641):
    if 20 <= i <= 60:
        cabinet_type = "Computered"
    elif 101 <= i <= 140:
        cabinet_type = "Normal" if i % 2 == 0 else "Laboratory"
    elif 201 <= i <= 240:
        cabinet_type = "Normal" if i % 2 == 0 else "Laboratory"
    else:
        cabinet_type = "Laboratory" if i % 2 == 0 else "Normal"
    
    cabinets.append({
        "id": io,
        "name": str(i),
        "type": cabinet_type,
        "capacity": 1
    })

    io += 1


cabinets.append({
    "id": io,
    "name": "sport",
    "type": "Sport",
    "capacity": 0
})
# Преобразование в JSON
cabinets_json = json.dumps(cabinets, ensure_ascii=False, indent=4)


# Запись результата в файл
with open('cabinets.json', 'w', encoding='utf-8') as f:
    f.write(cabinets_json)

print("Массив кабинетов успешно сохранён в cabinets.json")
