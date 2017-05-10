# ARDB

A simple throw-away utility tool, used to run multiple ardb servers.

## Install

```
$ go install github.com/glendc/ardb
```

## Usage

After installing you can simply run as follows:

```
$ ardb
localhost:16380
```

Multiple servers can be run using the `-n` flag,
5 servers for example:

```
$ ardb -n 5
localhost:16380,localhost:16381,localhost:16382,localhost:16383,localhost:16384
```

**Note**: it is required that you have installed [redis](https://redis.io) in order to use this tool.

More Information:

```
$ ardb -h
Usage of ardb:
  -n int
    	amount of ardb servers to run (default 1)
  -port int
    	first of multiple consecutive ardb-used ports (default 16380)
  -v	print ardb output
```
