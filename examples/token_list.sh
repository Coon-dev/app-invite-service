curl -i -k -X 'POST' \
  'https://app-invite-service.herokuapp.com/token_list' \
  -H 'accept: application/json' \
  -H 'Authorization: Basic a2mJIp6IOyZihYvw60WSwzprkB8AHGyOxtvmh0k1U4Lr0upv1LVpi4y' \
  -H 'Content-Type: application/json' \
  -d ''

# Response
# HTTP/1.1 200 OK
# Server: Cowboy
# Connection: keep-alive
# Content-Type: application/json; charset=utf-8
# Date: Wed, 30 Mar 2022 09:48:05 GMT
# Content-Length: 134
# Via: 1.1 vegur

# {"token_list":[{"token":"jJkwzDk","status":"active","created_at":"2022-03-30T09:10:38.827Z","expired_at":"2022-04-06T09:10:38.827Z"},{"token":"9jZ8uVbhV3vC","status":"inactive","created_at":"2022-03-30T09:48:38.267Z","expired_at":"2022-04-06T09:48:38.267Z"}]}