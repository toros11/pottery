# Summary

Pottery is a simple GUI module works on clay.
It provides graphical network diagram views based on clay datastore.

### Sample UI - network design
![Network design](./images/sample1.png)

### Sample UI - physial diagram from the system model store
![Physical diagram](./images/sample2.png)

### Sample UI - logical diagram from the system model store
![Logical diagram](./images/sample3.png)

# How to use

```
$ # note: add "pottery.HookSubmodules()" in HookSubmodules()#submodules.go in clay.
$ # note: in current version, you need to create symlink to pottery ui directory in the same directory which clay is located.
$ cd /path/to/clay
$ go build
$ ./clay &
```

UI runs at http://localhost:8080/ui/ unless set HOST and PORT environmental variables.

# Example diagram

```
$ curl -X GET "localhost:8080/v1/diagrams/physical"
$ curl -X GET "localhost:8080/v1/diagrams/logical"
```

## Endpoint list

### Diagram Resource

```
GET    /<version>/diagrams/physical
GET    /<version>/diagrams/logical
```

# Thanks

* Pottery uses https://github.com/codeout/inet-henge