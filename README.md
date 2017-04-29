# Pottery

[![Build Status](https://travis-ci.org/qb0C80aE/pottery.svg?branch=develop)](https://travis-ci.org/qb0C80aE/pottery)

Pottery is a simple GUI module works with [Loam](https://github.com/qb0C80aE/loam) on [Clay](https://github.com/qb0C80aE/clay).
It provides graphical network diagram views based on clay datastore.

### UI - network design
![Network design](./images/sample1.png)

### UI - physial diagram from the system model store
![Physical diagram](./images/sample2.png)

### UI - logical diagram from the system model store
![Logical diagram](./images/sample3.png)

# How to build and run

```
$ # note: add the '_ "github.com/qb0C80aE/{loam, pottery}" // Install {Loam, Pottery} sub module by importing' lines into the import section of submodules.go in Clay.
$ # note: in current version, you need to create symlink to pottery ui directory in the same directory which Clay is located.
$ cd /path/to/clay
$ go build
$ ./clay &
```

# How to use

## Example diagram

The diagram resource return the physical and logical diagram for [inet-henge](https://github.com/codeout/inet-henge) based on Loam data models.

```
$ curl -X GET "localhost:8080/diagrams/physical"
$ curl -X GET "localhost:8080/diagrams/logical"
```

# API Server

## Endpoint list

### Diagram Resource

```
GET    /diagrams/physical
GET    /diagrams/logical
```

# Thanks

* Pottery uses https://github.com/codeout/inet-henge
* Pottery uses http://www.quackit.com/html/templates/download/bootstrap/portal-1/
