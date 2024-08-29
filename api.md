# Blocks

Response Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#RpcBlock">RpcBlock</a>

Methods:

- <code title="get /blocks/{height}">client.Blocks.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#BlockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, height <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#RpcBlock">RpcBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreTransactionDocument">StoreTransactionDocument</a>

Methods:

- <code title="get /blocks/{height}/transactions">client.Blocks.Transactions.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#BlockTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, height <a href="https://pkg.go.dev/builtin#int64">int64</a>) ([][]<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk">satstream</a>.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#StoreTransactionDocument">StoreTransactionDocument</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Fees

Methods:

- <code title="get /fees">client.Fees.<a href="https://pkg.go.dev/github.com/satstream/satstream-go-sdk#FeeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/builtin#int64">int64</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
