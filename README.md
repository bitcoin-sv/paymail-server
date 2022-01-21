# Paymail Server

Paymail Server is a basic reference implementation of the Paymail Standard service discovery protocol.
This is written in go and integrates with a wallet running the Payment Protocol PayD Interface.

For the most part is does two things.

 - Generates two capabilities files for static hosting.
 - Responds to pki and p2paymail capability requests

The p2paymail standard is a halfway point between standard bitcoin address use, and proper SPV. 
Under the hood, this server will translate between p2paymail and invoice based payments to minimise creation of code we know will 
likely be decprecated after widespread SPV is in use. In the mean time this can act as a bridge between those worlds.

## Configuring Paymail

The server has a series of environment variables that allow you to configure the behaviours and integrations of the server.
Values can also be passed at build time to provide information such as build information, region, version etc.

### Server

| Key                    | Description                                                           | Default        |
|------------------------|-----------------------------------------------------------------------|----------------|
| DOMAIN_TLD             | Domain name and top level domain on which this paymail service runs   | nchain.com     |
| SERVER_PORT            | Port which this server should use                                     | :8446          |
| SERVER_HOST            | Host name under which this server is found                            | paymail        |
| SERVER_SWAGGER_ENABLED | If set to true we will expose an endpoint hosting the Swagger docs    | true           |
| SERVER_SWAGGER_HOST    | Sets the base url for swagger ui calls                                | localhost:8446 |

### Environment / Deployment Info

| Key                 | Description                                                                | Default          |
|---------------------|----------------------------------------------------------------------------|------------------|
| ENV_ENVIRONMENT     | What enviornment we are running in, for example 'production'               | dev              |
| ENV_REGION          | Region we are running in, for example 'eu-west-1'                          | local            |
| ENV_COMMIT          | Commit hash for the current build                                          | test             |
| ENV_VERSION         | Semver tag for the current build, for example v1.0.0                       | v0.0.0           |
| ENV_BUILDDATE       | Date the code was build                                                    | Current UTC time |
| ENV_BITCOIN_NETWORK | What bitcoin network we are connecting to (mainnet, testnet, stn, regtest) | regtest          |

### Logging

| Key       | Description                                                           | Default |
|-----------|-----------------------------------------------------------------------|---------|
| LOG_LEVEL | Level of logging we want within the server (debug, error, warn, info) | info    |

### PayD Wallet

| Key         | Description                                              | Default |
|-------------|----------------------------------------------------------|---------|
| PAYD_HOST   | Host for the wallet we are connecting to                 | payd    |
| PAYD_PORT   | Port the PayD wallet is listening on                     | :8443   |
| PAYD_SECURE | If true the P4 server will validate the wallet TLS certs | false   |
| PAYD_NOOP   | If true we will use a dummy data store in place of payd  | true    |

## Working with Paymail Server

There are a set of makefile commands listed under the [Makefile](Makefile) which give some useful shortcuts when working
with the repo.

Some of the more common commands are listed below:

`make pre-commit` - ensures dependencies are up to date and runs linter and unit tests.

`make build-image` - builds a local docker image, useful when testing paymail-server in docker.

`make run-compose` - runs Paymail Server in compose, a reference PayD wallet will be added to compose soon NOTE the above command will need ran first.
