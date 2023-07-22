# open-microservice-api

1. send mail via google mail

## Install
set environment
```
PORT=8080;MAIL_USER=;MAIL_PASSWORD=
```
MAIL_USER is your google acount,MAIL_PASSWORD is your google account's app password
## Get Google Account App Password
[https://myaccount.google.com/security](https://myaccount.google.com/security)
turn on 2-Step Verification
![](readme_files/4.jpg)
App password
![](readme_files/2.jpg)
geneate app password
![](readme_files/3.jpg)
![](readme_files/1.jpg)

## send mail
### curl
```
curl --location --request POST 'https://localhost:8080' \
--header 'User-Agent: apifox/1.0.0 (https://www.apifox.cn)' \
--header 'Content-Type: application/json' \
--data-raw '{
  "to": "to",
  "subject": "sbuject",
  "body": "body"
}'
```
### pytyon requests 
````
import requests
import json

url = "https://localhost:8080"

payload = json.dumps({
   "to": "litong@hawaii.edu",
   "subject": "王志安 更新",
   "body": "王志安有更新了信息,请访问下面的地址查看."
})
headers = {
   'User-Agent': 'apifox/1.0.0 (https://www.apifox.cn)',
   'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

```