curl -i -k -X 'POST' \
  'https://app-invite-service.herokuapp.com/token_detail' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "token": "aq21V2das"
}'

# Response
# HTTP/1.1 200 OK
# Server: Cowboy
# Connection: keep-alive
# Content-Type: application/json; charset=utf-8
# Date: Wed, 30 Mar 2022 09:11:15 GMT
# Content-Length: 22
# Via: 1.1 vegur

# {"status":"not_found"}


curl -i -k -X 'POST' \
  'https://app-invite-service.herokuapp.com/token_detail' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "token": "jJkwzDk"
}'

# Response
# HTTP/1.1 200 OK
# Server: Cowboy
# Connection: keep-alive
# Content-Type: application/json; charset=utf-8
# Date: Wed, 30 Mar 2022 09:14:02 GMT
# Content-Length: 19
# Via: 1.1 vegur

# {"status":"active"}