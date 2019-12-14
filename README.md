# management-tracking-system
A POC of a managment and tracking system for delivery using GO, ~~MongoDB~~, ~~Redis~~ and Docker

## [management-tracking-system/orders](https://github.com/israel-martinez/management-tracking-system/tree/develop/orders)

This is a REST microservices that receive /POST /GET /PATH /DELETE for rutes exposed using [github.com/gorilla/mux](github.com/gorilla/mux). This are the endpoints:

*  [**GET**] http://localhost:8080/				    : for get name of microservice
*  [**GET**] http://localhost:8080/orders		    : for get all orders
*  [**POST**] http://localhost:8080/orders		    : for create a new order
*  [**GET**] http://localhost:8080/orders/{id}      : for get one order by id
*  [**PATCH**] http://localhost:8080/orders/{id}    : for update one order by id
*  [**PATCH**] http://localhost:8080/orders/{id}    : for delete one order by id


See postman collections for test microservice in [management-tracking-system/postman-collections](https://github.com/israel-martinez/management-tracking-system/tree/develop/postman-collections)


**WIP** Store Registers in a MongoDB docker image, this is  feature mongo-db

### Deploy Docker Image of Microservice and Run
**Note:** Your need install Docker in your Desktop and start Docker Service, see [Docker Installation](https://docs.docker.com/v17.09/engine/installation/).

Clone the repo and go to the projet path /management-tracking-system/orders. 

Build Docker Image
```
$ docker build --no-cache -t papitajuan/orders:latest .
```

For See docker images. Your would see **papitajuan/orders** :
```
$ docker images
```
Next Run the created Docker Image with:
```
$ docker run --rm --name papitajuan-orders -d -p 8080:8080 -it papitajuan/orders:latest
```

You can see all Docker Containers with:

```
$ docker ps
```

**Note:** You too can start microservice only with GO in your local machine.

Go to path /management-tracking-system/orders and compile the GO script:
```
$ go build
```
Run GO microservices of orders:
```
$ ./orders
```


## [management-tracking-system/delivery]()

**TODO** Develop delivery management based on **management-tracking-system/orders** API functions and MongoDB in docker configurations for independent DB instances.

## [management-tracking-system/tracking]()

**TODO** Develop delivery management based on **management-tracking-system/orders** API functions and MongoDB in docker configurations for independent DB instances. Also add Redis for DB distribuited in memory for best performance in realtime tracking, and use WebSockets for bidirectional connections between FrontEnd Apps and Legacy Backend.
