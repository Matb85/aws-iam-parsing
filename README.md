# Home assignment for the Remitly Summer Internship 2024 recruitment process

## Created by Mateusz Bis 10 IV 2024

### Features

This project impements an **verifyPolicyJSON** method that:

- takes in a single argument - an AWS IAM Policy string as []bytes
- verifies whether the []bytes string is a **valid** AWS::IAM::Role Policy
- returns a bool value: **true** if the Resource policy field contains exactly one asterisk character, otherwise it returns **false** (inluding when the policy is invalid)

The method can be executed by running main.go - it reads JSON from data.json (which should be located in the root working directory), passes it to **verifyPolicyJSON** and prints out the function's result.

### Requirements

I've tested the project on go1.22.2, I cannot guarantee it will work on older versions.

Link to download Golang version 1.22.2: https://go.dev/dl/

### Steps to run the project:

1. Clone the project

```bash
$ git clone https://github.com/Matb85/remitly-home-assignment.git
$ cd ./remitly-home-assignment
```

2. Install dependencies

```bash
$ go get ./...
# or
$ go mod download
```

3. Prepare data and insert it to the data.json file

```bash
$ echo YOURDATA >> data.json

```

4. Run the project

```bash
$ go run ./main.go
```

5. Alternatively, build & run the project

```bash
$ go build ./main.go
$ ./main
```

### Unit Tests

I have written several unit tests for **verifyPolicyJSON** and **UnmarshalJSON** methods in order to ensure that everything works as expected.

To run the test suite, execute the following command:

```bash
$ go test ./...
```
