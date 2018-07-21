## HTTP-DB

http-db is a simple in-memory key-value database. All operations can be performed by the http-methods GET, PUT and DELETE.

There is no build-in authentication and authorization mechanism.

### Examples

Start http-db as a docker container:

```sh
docker run --name=http-db -d -p 8080:8080 psanetra/http-db
```

Save plain text "Hello World" at absolute path /hello/world

```sh
curl -v -XPUT -H 'Content-Type: text/plain' --data "Hello World" http://localhost:8080/hello/world
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> PUT /hello/world HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Type: text/plain
> Content-Length: 11
> 
* upload completely sent off: 11 out of 11 bytes
< HTTP/1.1 201 Created
< Location: /hello/world
< Date: Sat, 21 Jul 2018 15:51:48 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
```

Get the resource

```sh
curl -v http://localhost:8080/hello/world
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hello/world HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: text/plain
< X-Content-Type-Options: nosniff
< Date: Sat, 21 Jul 2018 15:56:35 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
Hello World
```

Delete the resource 

```sh
curl -v -XDELETE http://localhost:8080/hello/world
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> DELETE /hello/world HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 204 No Content
< Date: Sat, 21 Jul 2018 15:58:53 GMT
< 
* Connection #0 to host localhost left intact
```
