# Go client for Weatherbit API

Weatherbit.io provides an API that allows access to forecasts, current and historical data for locations across the world.

Weatherbit API documentation can be found [here](https://www.weatherbit.io/api).

The API endpoints currently supported by the Go API wrapper matches the endpoints supported by the [Python API wrapper](https://github.com/weatherbit/weatherbit-python).

The Go program has three requirements:

Go 1.10 or later is required as the program uses [strings builder](https://golang.org/src/strings/builder.go).

|                                             |           |
| ------                                      |:--------- |
| [API key](https://www.weatherbit.io/pricing)| Save in api_key.txt. [Requirements](#requirements) section provides more details   |
| Location                                    | modify wb.Parameters - see populaterequestparameters() in ./example/main.go        |
| Granularity                                 | modify wb.Parameters - see populaterequestparameters() in ./example/main.go        |

## Example

With a Golang installation and an API key saved in the **example** folder, open a terminal from the **example** directory and run:

```bash
go run main.go
```

WARNING: this iteration of the Weatherbit wrapper experiments with strings.Builder(), which is provided in Go 1.10 onwards.

## Design decisions

In accordance with [Mat Ryer - Writing Beautiful Packages in Go](https://youtu.be/cAWlv2SeQus?t=794), the weatherbit-go package is not asynchronous but can be used asynchronously should the user wish to do so using Golang primitives. A slice of type Parameter is recommended over a map, unless you *really* know what you're doing.

## Documentation

Available at [GoDoc](https://godoc.org/github.com/alljames/weatherbit-go/weatherbit)

## Installation

```bash
go get github.com/alljames/weatherbit-go/
```

To uninstall delete the directory in your $GOPATH/src.

## Requirements

Weatherbit [API key](https://www.weatherbit.io/pricing) - free tier option available.

Place your API Key into a file named api_key.txt (inside the same directory as the *.go* file where the weatherbit package is invoked). Don't forget to add api_key.txt to your .gitignore file if using a public repository. An example .gitignore file is provided in this repository.

## Acknowledgments

[Weatherbit.io](https://github.com/weatherbit)

## Contribution guide

Get involved! This package has a *lot* of room for improvement. Below are a list of Go features which could be used in this package to good effect.

* Tests! Testing has significant room for improvement.
* Handle spaces in City name more gracefully - i.e. if space, convert to %20. This is a fairly easy first issue.

Read the docs, write and try out some improvements and then make a pull request :smile: !

## Use cases

Coming soon.