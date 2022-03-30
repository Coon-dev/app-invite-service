# app-invite-service

This service is deployed on heruku. You can curl test from your postman or other tools.

Document APIs in Swagger: https://coon-dev.github.io/app-invite-service/

Note: the Swagger cannot call service directly due to CORS error. You could copy the curl command to run on your terminal.

Service use MongoDB online for database

If you want to run code in your local just change this line

>port := os.Getenv("PORT")

to

>port := "8080" 


build: 
>go build

test: 
> go test -v ./test

run: 
>./app-invite-service
