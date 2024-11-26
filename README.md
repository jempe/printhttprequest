# PrintHTTPRequest

Print Header and Body of HTTP requests.

## Description

This repository contains a simple Go application that prints the headers and body of incoming HTTP requests. It is useful for debugging and inspecting HTTP requests.

## Features

- Prints HTTP request headers
- Prints HTTP request body

## Installation

To install and run the application, you need to have Go installed on your system. You can download and install Go from the [official website](https://golang.org/dl/).

Clone the repository:
```sh
git clone https://github.com/jempe/printhttprequest.git
cd printhttprequest
```

## Usage

Run the application with a specified port:

```sh
go run main.go -port=<PORT>
```

Replace <PORT> with the port number you want the web service to run on.

Example:
```sh
go run main.go -port=8080
```

## Example Output

When you send an HTTP request to the server, the application will print the headers and body of the request.

Example:
```sh
Header:
map[Content-Length:[18] Content-Type:[application/json]]

Body:
{"key": "value"}
```

## License

This project is licensed under the Apache 2.0 License. See the LICENSE file for details.





