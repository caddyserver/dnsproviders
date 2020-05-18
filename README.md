DNS providers for Caddy v1 (obsolete)
=====================================

**⚠️ This repo is for Caddy v1, which has been obsoleted by Caddy 2.**

[Caddy 2 also supports the DNS challenge](https://caddyserver.com/docs/automatic-https#dns-challenge) in a similar way to v1, but [using backwards-incompatible APIs](https://github.com/caddy-dns) that are much more flexible and easier to use. This repository is no longer relevant or maintained.

Migrating to the new APIs is solely a community effort. [It is very easy to implement a provider](https://caddy.community/t/writing-new-dns-provider-modules-for-caddy/7786) if you know Go!

**Version 2 resources:**

- [How to use DNS providers in Caddy 2](https://caddy.community/t/how-to-use-dns-provider-modules-in-caddy-2/8148)
- [Implementing a new DNS provider for Caddy 2](https://caddy.community/t/writing-new-dns-provider-modules-for-caddy/7786)
- [libdns packages](https://github.com/libdns) (underlying provider implementations)
- [caddy-dns packages](https://github.com/caddy-dns) (small wrappers over libdns providers that convert them into Caddy modules)
- [Temporary module shim](https://github.com/caddy-dns/lego-deprecated) (supports all of lego's 75+ DNS providers)



Old readme (for v1)
===================

These providers can be used to help solve the ACME DNS challenge by plugging them into Caddy 0.9-1.x:

```go
import _ "github.com/caddyserver/dnsproviders/<provider>"
```

You can then use this in your Caddyfile with the `tls` directive like so:

```plain
tls {
	dns <provider>
}
```

Credentials for your DNS provider should be set in environment variables. This information is in the [Automatic HTTPS](https://caddyserver.com/docs/automatic-https#dns-challenge) page of the Caddy documentation. For more information about using your DNS provider, see [the docs for your provider](https://godoc.org/github.com/go-acme/lego/providers/dns) directly.

If you specify a DNS provider, the DNS challenge will be used exclusively; other challenge types will be disabled. Be aware that some DNS providers may be slow in applying changes.


## About these packages

Caddy 0.9 and newer supports solving the ACME DNS challenge. This challenge is unique because the server that is requesting a TLS certificate does not need to start a listener and be accessible from external networks. This quality is essential when behind load balancers or in other advanced networking scenarios.

The DNS challenge sets a DNS record and the ACME server verifies its correctness in order to issue the certificate. Caddy can do this for you automatically, but it needs credentials to your DNS provider to do so. Since every DNS provider is different, we have these adapters you can plug into Caddy in order to complete this challenge.

The underlying logic that actually solves the challenge is implemented in a different package not far away from here. Caddy uses [go-acme/lego](https://github.com/go-acme/lego), a library originally written for use in Caddy, to solve ACME challenges. If you wish to add a new provider, see the documentation for that library and write your own provider implementation. Then writing the adapter for Caddy is very easy: just copy+paste any of these existing ones, replace the names and tweak a few things, and submit a pull request. Done!
