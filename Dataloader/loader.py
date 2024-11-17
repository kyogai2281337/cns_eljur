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
        'auth': "<token>"
    }
    try:
        print(f"Sending request to: http://localhost/api/admin/private/create")
        print(f"Request payload: {json.dumps(data, indent=2)}")
        
        response = requests.post(
            'http://localhost/api/admin/private/create',
            headers=headers,
            cookies=cookies,
            json=data
        )
        
        print(f"Response status code: {response.status_code}")
        print(f"Response headers: {dict(response.headers)}")
        print(f"Response body: {response.text}")
        
        response.raise_for_status()
        result = response.json()
        print(f"Operation successful. Response data: {json.dumps(result, indent=2)}")
        return result
        
    except requests.exceptions.ConnectionError as e:
        print(f"Connection error occurred: {e}")
        print(f"Failed to connect to server at: http://localhost/api/admin/private/create")
        print(f"Request details: {e.__class__.__name__}")
        return None
        
    except requests.exceptions.Timeout as e:
        print(f"Request timed out: {e}")
        print(f"Timeout duration: {e.timeout if hasattr(e, 'timeout') else 'unknown'}")
        print(f"Request details: {e.__class__.__name__}")
        return None
        
    except requests.exceptions.HTTPError as e:
        print(f"HTTP error occurred: {e}")
        print(f"Response status code: {e.response.status_code}")
        print(f"Response headers: {dict(e.response.headers)}")
        print(f"Response text: {e.response.text}")
        print(f"Request URL: {e.response.url}")
        print(f"Request method: {e.response.request.method}")
        print(f"Request headers: {dict(e.response.request.headers)}")
        print(f"Request body: {e.response.request.body}")
        return None
        
    except requests.exceptions.RequestException as e:
        print(f"An error occurred while making the request: {e}")
        print(f"Error type: {e.__class__.__name__}")
        print(f"Error details: {str(e)}")
        if hasattr(e, 'response'):
            print(f"Response status code: {e.response.status_code if e.response else 'No response'}")
            print(f"Response headers: {dict(e.response.headers) if e.response else 'No response'}")
            print(f"Response body: {e.response.text if e.response else 'No response'}")
        return None
    
def read_json_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        data = json.load(file)
        return data

def main():
    cabinetsData = read_json_file("data/cabinets.json")
    total = len(cabinetsData)
    for i, cabinet in enumerate(cabinetsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "cabinets",
            "data": cabinet
        })
        time.sleep(0.1)

    subjectsData = read_json_file("data/subjects.json")
    total = len(subjectsData)
    for i, subject in enumerate(subjectsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "subjects",
            "data": subject
        })
        time.sleep(0.1)

    specializationsData = read_json_file("data/specializations.json")
    total = len(specializationsData)
    for i, specialization in enumerate(specializationsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "specializations",
            "data": specialization
        })
        time.sleep(0.1)

    groupsData = read_json_file("data/groups.json")
    total = len(groupsData)
    for i, group in enumerate(groupsData, 1):
        print(f"Processing {i}/{total}", end='\r')
        post_create_private({
            "tablename": "groups",
            "data": group
        })
        time.sleep(0.1)

    teachersData = read_json_file("data/teachers.json")
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