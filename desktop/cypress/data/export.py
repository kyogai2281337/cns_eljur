import json
import requests
import time
import os
import logging
from tqdm import tqdm
import urllib3

# Отключение предупреждений InsecureRequestWarning
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# Настройка логирования
if not os.path.exists('logs'):
    os.makedirs('logs')

logging.basicConfig(
    filename=os.path.join('logs', 'upload_errors.log'),
    level=logging.INFO,
    format='%(asctime)s %(levelname)s: %(message)s',
    datefmt='%Y-%m-%d %H:%M:%S'
)

# URL и куки для авторизации
url = "https://localhost/api/admin/private/create"
auth_cookie = {"auth": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZXhhbXBsZS5jb20iLCJleHAiOjE3MjU1NzM1NDQsImlkIjoxLCJyb2xlIjoic3VwZXJ1c2VyIn0.rjO1vKl1rHhHfgzLtQfdqUSmcvtIiXy85D3-hkFwLxM"}

# Функция для отправки данных
def post_data(tablename, data):
    payload = {
        "tablename": tablename,
        "data": data
    }
    try:
        logging.info(f"Sending {tablename} data: {json.dumps(payload, ensure_ascii=False)}")
        response = requests.post(url, json=payload, cookies=auth_cookie, verify=False)
        time.sleep(1)  # Задержка в 1 секунду между запросами
        logging.info(f"Received response for {tablename}: {response.status_code} - {response.text}")
        if response.status_code == 200:
            return response.json()
        else:
            logging.error(f"Error in {tablename}: {response.status_code} - {response.text}")
            return None
    except requests.exceptions.RequestException as e:
        logging.error(f"Request failed for {tablename}: {e}")
        return None

# Загрузка данных из файлов JSON
def load_json_file(filename):
    with open(filename, 'r', encoding='utf-8') as f:
        return json.load(f)

# Функция для загрузки и отправки данных с прогресс-баром
def upload_data(filename, tablename, process_function):
    data = load_json_file(filename)
    results = []
    for item in tqdm(data, desc=f"Uploading {tablename}"):
        result = process_function(item)
        if result:
            results.append(result)
    return results

# Загрузка и отправка subjects
def process_subjects(subject):
    resp = post_data("subjects", subject)
    if resp:
        print(f"Subject uploaded: {resp}")
        return resp  # Возвращаем ответ от сервера

subjects_responses = upload_data('subjects.json', "subjects", process_subjects)

# Загрузка и отправка cabinets
def process_cabinets(cabinet, cabinet_id):
    cabinet['id'] = cabinet_id  # добавляем ID для кабинета
    resp = post_data("cabinets", cabinet)
    if resp:
        print(f"Cabinet uploaded: {resp}")
        return resp  # Возвращаем ответ от сервера

def upload_cabinets():
    cabinets = load_json_file('cabinets.json')
    results = []
    for cabinet_id, cabinet in enumerate(cabinets, start=1):
        result = process_cabinets(cabinet, cabinet_id)
        if result:
            results.append(result)
    return results

cabinets_responses = upload_cabinets()

# Загрузка и отправка specializations
def process_specializations(specialization, specialization_id):
    specialization['id'] = specialization_id  # добавляем ID для специализации
    resp = post_data("specializations", specialization)
    if resp:
        print(f"Specialization uploaded: {resp}")
        return resp  # Возвращаем ответ от сервера

def upload_specializations():
    specializations = load_json_file('specializations.json')
    results = []
    for specialization_id, specialization in enumerate(specializations, start=1):
        result = process_specializations(specialization, specialization_id)
        if result:
            results.append(result)
    return results

specializations_responses = upload_specializations()

# Загрузка и отправка groups
def process_groups(group, specialization_id):
    group['specialization']['id'] = specialization_id  # привязка к правильному ID специализации
    resp = post_data("groups", group)
    if resp:
        print(f"Group uploaded: {resp}")
        return resp  # Возвращаем ответ от сервера

def upload_groups():
    groups = load_json_file('groups.json')
    results = []
    for group_id, group in enumerate(groups, start=1):
        result = process_groups(group, group['specialization']['id'])
        if result:
            results.append(result)
    return results

groups_responses = upload_groups()

# Загрузка и отправка teachers
def process_teachers(teacher):
    resp = post_data("teachers", teacher)
    if resp:
        print(f"Teacher uploaded: {resp}")
        return resp  # Возвращаем ответ от сервера

teachers_responses = upload_data('teachers.json', "teachers", process_teachers)
