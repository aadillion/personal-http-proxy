# personal-http-proxy
Personal Http Proxy:

CURL of the Request:
```
curl --location --request POST 'localhost:8090/proxy' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://google.com",
    "method": "GET",
    "headers": {
        "authorization": "token"
    }
}'
```
