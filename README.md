# Fizz-Buzz
Fizz-Buzz Rest server.
A pimp version of the FizzBuzz project: 

`
The goal is to implement a web server that will expose a REST API endpoint that: make a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
`

## Build and Run the project

To build the project inside a docker container with a postgres instance use the following command:

´docker-compose build && docker-compose up -d´

If you have a local postgres instance you can run the project using:

`POSTGRES_HOST=localhost go run ./cmd/fizzbuzzd/main.go`

Once you have your server running you can used the Postman collection in the pkg/ directory.
Use GetFizzBuzz with assossiate body (int1, int2, str1, str2 and limit).
Use GetFizzBuzzStats to find your most made request.

## Test the project

To run test over the poject without postgres instance you can locally use:

`make test`

With postgres instance:

`make docker-integration`

