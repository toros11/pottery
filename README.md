# Summary

Pottery is a simple GUI module works with [Loam](https://github.com/qb0C80aE/loam) on [Clay](https://github.com/qb0C80aE/clay).
It provides graphical network diagram views based on clay datastore.

### Sample UI - network design
![Network design](./images/sample1.png)

### Sample UI - physial diagram from the system model store
![Physical diagram](./images/sample2.png)

### Sample UI - logical diagram from the system model store
![Logical diagram](./images/sample3.png)

# How to build and run

```
$ # note: add "pottery.HookSubmodules()" in HookSubmodules()#submodules.go in clay.
$ # note: in current version, you need to create symlink to pottery ui directory in the same directory which clay is located.
$ cd /path/to/clay
$ go build
$ ./clay &
```

## Environmental variables

You can give the environmental variables to Clay.

|Key         |Description                                                                      |Options    |Default        |
|:-----------|:--------------------------------------------------------------------------------|:----------|:--------------|
|ENDPOINT    |The host:port formed API address what this UI access                             |-          |${HOST}:${PORT}|

# How to use

## Example diagram

The diagram resource return the physical and logical diagram for [inet-henge](https://github.com/codeout/inet-henge) based on Loam data models.

```
$ curl -X GET "localhost:8080/v1/diagrams/physical"
$ curl -X GET "localhost:8080/v1/diagrams/logical"
```

# API Server

## Endpoint list

### Diagram Resource

```
GET    /<version>/diagrams/physical
GET    /<version>/diagrams/logical
```

# Thanks

* Pottery uses https://github.com/codeout/inet-henge