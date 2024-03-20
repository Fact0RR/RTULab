
import requests
import time
import shutil

photo = str(int(time.time()))

json1 = {
    "camera_type":"camerus1",
    "data":
    {"transport_chars":"ABS",
	"transport_numbers":"123",
	"transport_region":"77",
	"camera_id": "1adf43fe-1a25-4f0d-b93788f8adcf5d88a",
	"violation_id":"f938b151-30d5-495f-92b3-d222ade609e1",
	"violation_value":"435km/h",
	"skill_value":1,
	"datetime":"2024-01-01T00:00:00+00:00"},
    "photo": photo
}

json2 = {
    "camera_type":"camerus2",
    "data":
    {
	"transport":{
		"chars":"ABS",
		"numbers":"123",
		"region":"77"
	},

	"camera": {
		"id":"1adf43fe-1a25-4f0d-b93788f8adcf5d88b"
	},

	"violation":{
		"id": "f938b151-30d5-495f-92b3-d222ade609e1",
		"value": "13km/h"
	},

	"skill":{
		"value": 1
	},

	"datetime":{
		"year": 2024,
		"month":1,
		"day": 1,
		"hour":0,
		"minute":0,
		"seconds":0,
		"utc_offset":"+00:00"
	}
},
    "photo": photo
}

json3 = {
    "camera_type":"camerus3",
    "data":
    {
	"transport":"ABC12377",
	"camera": {
		"id":"1adf43fe-1a25-4f0d-b93788f8adcf5d88c"
	},
	"violation":{
		"id": "f938b151-30d5-495f-92b3-d222ade609e1",
		"value": "13km/h"
	},
	"skill":1,
	"datetime":1704067200
},
    "photo": photo
}

urlPhotoLocal = 'http://localhost:8181/upload'
urlrtu = "https://recruit.rtuitlab.dev/serialize"
ulrMainLocal = "http://localhost:8080/send"

#response = requests.post(url, files=files)

#response = requests.post(url, json=json)

#shutil.copy2('photo/12223.png', photo +'.png')

#копируем фото для отправки
shutil.copy2('photo/111.png', photo +'.png')
#записываем фото в 
files = {'myFile': open(photo+'.png', 'rb')}
#получаем []byte вместо json
responseMain = requests.post(urlrtu, json=json1)
#отправляем фото на специальный сервер
responsePhoto = requests.post(urlPhotoLocal, files=files)
#отправляем данные на основной сервер
response = requests.post(ulrMainLocal,responseMain)
