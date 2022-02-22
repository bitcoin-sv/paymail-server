# Paymail Server

This server runs a basic implementation of a Paymail standard service with cut down capabilities specifically built to 
integrate with a Direct Payments wallet. The main purpose is to allow incoming funds from existing paymail services.

Bitcoin wallets such as Handcash & MoneyButton implement a server to server paymail capability with brfcids 2a40af698840 and 5f1323cddf31 
which are a step away from standard bitcoin address use, towards proper SPV. This server will act as a bridge between the old and new,
by translating between that legacy protocol and the full SPV direct payment protocol.

## Configuring Paymail

The server has a series of environment variables that allow you to configure the behaviours and integrations of the server.

### Server

| Key                    | Description                                                            | Default          |
|------------------------|------------------------------------------------------------------------|------------------|
| DOMAIN_TLD             | Domain name and top level domain on which this paymail service runs    | carefulbear.com  |
| SERVER_PORT            | Port which this server should use                                      | :8446            |
| SERVER_HOST            | Host name under which this server is found                             | paymail          |
| PAYD_HOST              | Host for the wallet we are connecting to                               | payd             |
| PAYD_PORT              | Port the PayD wallet is listening on                                   | :8443            |

### DNS Records

In order to work with existing infrastructure, an SRV record is required pointing your desired domain at the domain you're running this server from.
```
# SRV Record
_bsvalias._tcp.yourdomain.com. 3600    10 10 443 where-ever-you-are-running-this-code.com.

# sometimes registrars will make you separate out these values, if so they are as follows:
service: bsvalias
protocol: tcp
TTL: 3600
Priority: 10
Weight: 10
Host: where-ever-you-are-running-this-code.com.

```

Needless to say the `where-ever-you-are-running-this-code.com.` must point to your host IP via an A record. If the domain is different then DNSSEC is required. Read more about the [Paymail Standard](https://tsc.bitcoinassociation.net/standards/paymail/) for details.


## Working with Paymail Server

There are a set of makefile commands listed under the [Makefile](Makefile) which give some useful shortcuts when working
with the repo.

Some of the more common commands are listed below:

`make pre-commit` - ensures dependencies are up to date and runs linter and unit tests.

`make build-image` - builds a local docker image, useful when testing paymail-server in docker.

`make run-compose` - runs Paymail Server in compose, a reference PayD wallet will be added to compose soon NOTE the above command will need ran first.
