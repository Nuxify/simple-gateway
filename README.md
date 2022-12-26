# Simple API Gateway
using Traefik

## Setup

1. Clone https://github.com/Nuxify/simple-gateway
2. Create `.env` (from .env.example) and provide correct values
3. Setup the services hosts and ports in .env
4. Make sure that the docker network is created before running this service (proxy by default)
4. Up the network with `make up`

## Services Configuration

- Make sure to appropriately name the `mapped` host to this standard: `*_API_URL=`
- The first argument of the variable `MUST` be equal to the service prefix in routes.

.env
```go
AUTH_API_URL=http://localhost:8000

is expected to re-route request from,

/v1/auth/{*}

AUTH must be auth. The gateway handles the word cases.

``` 
- Plural `services` are automatically pointed to its singular counterpart. Services ending with `es` or `ies` are pointed to their singular counterpart.
```go
Request to /v1/users will point automatically to USER_API_URL={some API}. The `users` service name is automatically casted to `USER`
```

Made with ❤️ at [Nuxify](https://nuxify.tech)
