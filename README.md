
<div align="center" style="display:flex;flex-direction:column;">
    <img width="300" src="https://imgdb.net/storage/uploads/495cc30ad5b741033ede8604cb0ef566cb48b5685a252f34de460850dabb82f6.png" alt="Probien logo"/>
  <h2>Microservice Core Written in Go - Services For a Pawn Shop Web Application</h2>
  <p>
    <a target="_blank" href="https://crowdin.com/project/excalidraw">
      <img src="https://img.shields.io/badge/License-GPL%20v3-yellow.svg">
    </a>
        <a target="_blank" href="https://crowdin.com/project/excalidraw">
      <img src="https://img.shields.io/github/last-commit/ThePandaDevs/Probien-Backend">
    </a>
      <h4>Phase 1/3, Im still working on this when I have time</h4>
  </p>
</div>

## Contributing

Contributions are always welcome!

This project is small, if you want to contribute improving code or add a new feature send a PR, make sure to add a description.

## Bussiness Logic

- Administrator
  - Can see all occurred logs into application (sessions, movements and payments).
  - Can see detailed reports by branch office (money recorded, total products, employee statistics, etc).
  - Can manage user access, create new branches, create new employees and create new category products.
  
- Employee
  - Can manage the category of products
  - Can do pawn orders to customers as well as do endorsements.
  - Can manage the status of pawn orders (in course, overdue, paid or lost)

- Customer
  - Can do multiple pawns, their stuff are evaluated and classified by a category product when a quote is realized.
  - Can do endorsements depending of payment modality (weekly or monthly).
  - The Customer has an extension date of payment depending on the modality
    - Weekly: 1 extra day to make the payment
    - Monthly: 3 extra days to make the payment 

- Extras
  - When a customer doesn't make any endorsement after the deadline, all the client's things become property of probien.

## Design - Following DDD (Domain Driven Design)
```
config/
├─ database.go | Connection to postgresql using ORM
├─ redis.go | Connection to redis cloud
|
├─ migrations/
│  ├─ stored procedures | Raw sql/
│  ├─ models | Entity struct that represent mapping to data model/
core/
├─ application | Write business logic/
├─ domain | Entity struct that represent I/O JSON format/
│  ├─ repository | Repository interface for infrastructure/
|
├─ infrastructure/
│  ├─ auth | Middleware and security filters/
│  ├─ persistance | Implements repository interface with database /
|
├─ interface | Expose http endpoints/
router | Routing for endpoints/
server.go
```

## Database Model (subject to change)

<img src="https://user-images.githubusercontent.com/67834146/181162633-8c323f57-3a70-4cc1-a5ae-7e2fc399464c.png" alt="database" border="0">

###

# Getting started - Manually

Please make sure to download and configure everything below to avoid problems.
## Dependencies needed
Gorm (ORM): 
``` go get -u gorm.io/gorm ```

Gin (Framework): ``` go get -u github.com/gin-gonic/gin ```

Go-Cron (Cron jobs): ``` go get -u github.com/go-co-op/gocron ```

Go-Redis: ``` go get -u github.com/go-redis/redis/v8 ```

Go-Jwt: ``` go get -u github.com/golang-jwt/jwt/v4 ```

Go-UUID: ``` go get -u github.com/satori/go.uuid ```

Godotenv: ``` go get -u github.com/joho/godotenv ``` 

## Environment vars
You must create a .env file called vars.env into the root project, then add the following environment vars:

```
export PRIVATE_KEY="your private key to sign tokens"

export DATABASE_URI_DEV="your database connection url for development"
export DATABASE_URI_PDN="your database connection url for production"

export REDIS_URI="your redis connection url"
export REDIS_PASSWORD="your redis password"
```


If it is the first time you are running the application, you must add the flag  obligately, add it after command:

| Flag | Type | Description|
| :---: | :---:  | :---: | 
| -migrate| boolean | Migrate datamodel structs and stored procedures to database

```
go run ./server.go -migrate=true
```
This will run the server, if you configured everything good, you will see the endpoints display on the console, at this point you can stop the server.


After configure the env vars and migrated the models, run project usually with following command:
```
go run ./server.go
```

# Getting started - Docker
Only execute the following command
```
docker build -t probien-backend .
docker run -it --rm --name probien-core probien-backend
```
