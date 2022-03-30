curl -i -k -X 'POST' \
  'https://app-invite-service.herokuapp.com/token_generate' \
  -H 'accept: application/json' \
  -H 'Authorization: Basic a2mJIp6IOyZihYvw60WSwzprkB8AHGyOxtvmh0k1U4Lr0upv1LVpi4y' \
  -H 'Content-Type: application/json' \
  -d ''

# Response
# HTTP/1.1 200 OK
# Server: Cowboy
# Connection: keep-alive
# Content-Type: application/json; charset=utf-8
# Date: Wed, 30 Mar 2022 09:10:40 GMT
# Content-Length: 129
# Via: 1.1 vegur

# {"token":"jJkwzDk","status":"active","created_at":"2022-03-30T09:10:38.827672632Z","expired_at":"2022-04-06T09:10:38.827672632Z"}