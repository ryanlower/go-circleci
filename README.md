# go-circleci

go-circleci is a Go client for the [CircleCI REST API](https://circleci.com/docs/api)

[![Circle CI](https://circleci.com/gh/ryanlower/go-circleci.svg?style=svg&circle-token=ea68f54409368a07d8168c06cf8e7ee583488ced)](https://circleci.com/gh/ryanlower/go-circleci)

**This is WIP. Currently only the `GET /me` endpoint is supported.**

## Usage

Create a new client using your API token. If you don't have one already, you can create one from the [CircleCI dashboard](https://circleci.com/account/api).

```go
import circleci "github.com/ryanlower/go-circleci"

client := NewClient("your_api_token")
```

You can then call the supported API methods

```go
user := client.Me()
```

## Roadmap
* Support more API endpoints, starting with projects and their builds (to make this client actually useful!)
* Better error handling. Too often this package just logs errors :(
