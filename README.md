# go-grpc-server

---

Golang example project represent gRPC server.

[Golang Client](https://github.com/zh1gr/go-grpc-client)<br>
[Golang Domain](https://github.com/zh1gr/go-grpc-domain)

### install

---

```shell
go run server.go
```

### commands

---

`make go_mod` - update dependencies and store it in **vendor** directory in project

`make run` - Run server.<br>


### project files

---

 - **server.go** - Contains the implementation of the server-side logic for the gRPC service.
 - **Makefile** - Script containing commands for automating common development tasks such as building, testing, and running the project.
 - **go.mod** - File defining the module and dependencies for the Go project.
 - *vendor* - Directory containing third-party dependencies managed by a dependency management tool like Go modules or dep.


### Credits

---

Based by tutorial from _Maximilien Andile_ [video](https://youtu.be/gbrPMv_GuQY?)