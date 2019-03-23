# pixela-client-go

[![Build Status](https://travis-ci.org/noissefnoc/pixela-client-go.svg?branch=master)](https://travis-ci.org/noissefnoc/pixela-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/noissefnoc/pixela-client-go)](https://goreportcard.com/report/github.com/noissefnoc/pixela-client-go)
[![Coverage Status](https://coveralls.io/repos/github/noissefnoc/pixela-client-go/badge.svg?branch=master)](https://coveralls.io/github/noissefnoc/pixela-client-go?branch=master)

[pixela-client-go](https://github.com/noissefnoc/pixela-client-go) is unofficial [pixe.la](https://pixe.la) API client & CLI written by golang.

This program build and check with Go 1.11.


## Synopsis

### Create user (just one time)

First, create [pixe.la](https://pixe.la) user.

```
$ pixela user create USERNAME TOKEN
```

NOTE: pixe.la does not have user page nor user profile API. I recommend to take a note `USERNAME` and `TOKEN`.


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
    pixela [global options] command [command options] subcommand [subcommand options] [arguments...]

VERSION:
    0.0.3
    
AUTHOR:
    noissefnoc <noissefnoc@gmail.com>
    
COMMANDS:
    user    Create, Update token, Delete user
    graph   Create, Get definition, Get SVG data, Update definition, Delete graph, Get pixels date
    pixel   Create, Get, Increment, Decrement, Update, Delete pixel
    webhook Create, Get, Invoke, Delete webhook

SUBCOMMANDS:
    user
        create Create user
        update Update user token
        delete Delete user
    graph
        create Create graph
        def    Get graph definitions (all graphs you created)
        svg    Get graph SVG format
        update Update graph definitions
        delete Delete graph
        pixels Get pixel regestored dates in the graph
    pixel
        create Create pixel
        get    Get pixel's quantitiy and optional data
        inc    Increment pixel quantity
        dec    Decrement pixel quantity
        update Update pixel quantity and optionl data
        delete Delete pixel
    webhook
        create Create webhook
        get    Get webhook
        invoke Invoke webhook
        delete Delete webhook

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
