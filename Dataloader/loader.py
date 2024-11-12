import requests
import json
import time
from dotenv import load_dotenv
import os

load_dotenv()

def post_create_private(data):
    headers = {
        'Content-Type': 'application/json'
    }
    cookies = {
        'auth': os.getenv('TOKEN')
    }
    try:
        response = requests.post(
            'http://localhost/api/admin/private/create',
            headers=headers,
            cookies=cookies,
            json=data
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.ConnectionError as e:
        print(f"Connection error occurred: {e}")
        print(f"Failed to connect to server at: http://localhost/api/admin/private/create")
        return None
    except requests.exceptions.Timeout as e:
        print(f"Request timed out: {e}")
        return None
    except requests.exceptions.HTTPError as e:
        print(f"HTTP error occurred: {e}")
        print(f"Response status code: {e.response.status_code}")
        print(f"Response text: {e.response.text}")
        print(f"Payload: {data}")
        return None
    except requests.exceptions.RequestException as e:
        print(f"An error occurred while making the request: {e}")
        print(f"Request details: {e.__class__.__name__}")
        return None

def read_json_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        data = json.load(file)
        return data

def main():
    cabinetsData = read_json_file("cabinets.json")
    total = len(cabinetsData)
    for i, cabinet in enumerate(cabinetsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "cabinets",
            "data": cabinet
        })
        time.sleep(0.1)

    subjectsData = read_json_file("subjects.json")
    total = len(subjectsData)
    for i, subject in enumerate(subjectsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "subjects",
            "data": subject
        })
        time.sleep(0.1)

    specializationsData = read_json_file("specializations.json")
    total = len(specializationsData)
    for i, specialization in enumerate(specializationsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "specializations",
            "data": specialization
        })
        time.sleep(0.1)

    groupsData = read_json_file("groups.json")
    total = len(groupsData)
    for i, group in enumerate(groupsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "groups",
            "data": group
        })
        time.sleep(0.1)

    teachersData = read_json_file("teachers.json")
    total = len(teachersData)
    for i, teacher in enumerate(teachersData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "teachers",
            "data": teacher
        })
        time.sleep(0.1)

if __name__ == "__main__":
    main()