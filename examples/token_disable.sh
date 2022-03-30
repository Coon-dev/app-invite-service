curl -i -k -X 'POST' \
  'https://app-invite-service.herokuapp.com/token_disable' \
  -H 'accept: application/json' \
  -H 'Authorization: Basic a2mJIp6IOyZihYvw60WSwzprkB8AHGyOxtvmh0k1U4Lr0upv1LVpi4y' \
  -H 'Content-Type: application/json' \
  -d '{
  "token": "9jZ8uVbhV3vC"
}'

# Response
# HTTP/1.1 200 OK
# Server: Cowboy
# Connection: keep-alive
# Date: Wed, 30 Mar 2022 09:49:25 GMT
# Content-Length: 0
# Via: 1.1 vegur