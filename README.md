# pixela-client-go

[pixela-client-go](https://github.com/noissefnoc/pixela-client-go) is unofficial [pixe.la](https://pixe.la) API client & CLI written by golang.

This program build and check with Go 1.11.


## Synopsis

### Create user (just one time)

First, create pixe.la user.

```
$ pixela user create USERNAME TOKEN
```

NOTE: pixe.la does not have user page nor get user info API. I recommend to take a note `USERNAME` and `TOKEN`.


### Create graph (just one time)

Second, create graph.

```
$ pixela graph create GRAPH_ID GRAPH_NAME UNIT TYPE COLOR [TIMEZONE]
```

NOTE: some arguments are limited following values.

* `TYPE` : `int` or `float`
* `COLOR` : `shibafu`, `momiji`, `sora`, `ichou`, `ajisai` or `kuro`
* `TIMEZONE` : default is `UTC`


### Record quantity to graph

And last, this is daily work for recording quantity to graph.

You can also modify quantity same command. (because `pixel/update` API create pixel when pixel has not create yet.)

```
$ pixela pixel update GRAPH_ID DATE QUANTYTY --optionalData='{"key":"value", ...}'
```

NOTE:

* `DATE` format is `yyyyMMdd`
* `--optionalData` format is json up to 10KB


## Usage

```
NAME:
    pixela - pixe.la client

USAGE:
    pixela [global options] command [command options] [arguments...]

VERSION:
    0.0.3
    
AUTHOR:
    noissefnoc <noissefnoc@gmail.com>
    
COMMANDS:
    user    Create, Update token, Delete user
    graph   Create, Get definition, Get SVG data, Update definition, Delete graph, Get pixels date
    pixel   Create, Get, Increment, Decrement, Update, Delete pixel
    webhook Create, Get, Invoke, Delete webhook
    
GLOBAL OPTIONS:
    --help, -h  show help
```


## Installation

### From Github release resource

Get from [Github release page](https://github.com/noissefnoc/pixela-client-go/releases).

### From homebrew (for macOS user)

```
$ brew install noissefnoc/tap/pixela
```

### From source code

```
$ go get github.com/noissefnoc/pixela-client-go
```


## License

This program is distributed under the MIT License. see LICENSE for more information.


## Author

[noissefnoc](noissefnoc@gmail.com)