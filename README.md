# yarpc [![GoDoc][doc-img]][doc] [![GitHub release][release-img]][release] [![Mit License][mit-img]][mit] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]

A message passing platform for Go that lets you:

* Write servers and clients with various encodings, including [JSON](http://www.json.org/), [Thrift](https://thrift.apache.org/), and [Protobuf](https://developers.google.com/protocol-buffers/).
* Expose servers over many transports simultaneously, including [HTTP/1.1](https://www.w3.org/Protocols/rfc2616/rfc2616.html), [gRPC](https://grpc.io/), and [TChannel](https://github.com/uber/tchannel).
* Migrate outbound calls between transports without any code changes using config.

Explore working code in the [examples](internal/examples) package, or read the following guides:

| Guide | Description |
| :---- | :------- |
| [Installation](.docs/installation.md) | Glide and SemVer |
| [Introduction](.docs/introduction.md) | Concepts and Vocabulary |
| [Getting Started](.docs/first-services.md) | Your First Services |
| [Adding Structure](.docs/json-encoding.md) | Using the JSON Encoding |
| [Understanding Errors](.docs/errors.md) | |
| [Middleware](.docs/middleware.md) | |
| [Binary Encodings](.docs/binary-encodings.md) | Thrift and Protobuf |
| [Configuring Transports](.docs/transports.md) | Configuring how messages get passed |

## Stability

This library is `v1` and follows [SemVer](http://semver.org/) strictly.

No breaking changes will be made to exported APIs before `v2.0.0` with the
**exception of experimental packages**.

Experimental packages reside within packages named `x`, and are *not stable*. This means their
APIs can break at any time. The intention here is to validate these APIs and iterate on them
by working closely with internal customers. Once stable, their contents will be moved out of
the containing `x` package and their APIs will be locked.

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://godoc.org/go.uber.org/yarpc

[release-img]: https://img.shields.io/github/release/yarpc/yarpc-go.svg
[release]: https://github.com/yarpc/yarpc-go/releases

[mit-img]: http://img.shields.io/badge/License-MIT-blue.svg
[mit]: https://github.com/yarpc/yarpc-go/blob/master/LICENSE

[ci-img]: https://img.shields.io/travis/yarpc/yarpc-go/master.svg
[ci]: https://travis-ci.org/yarpc/yarpc-go/branches

[cov-img]: https://codecov.io/gh/yarpc/yarpc-go/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/yarpc/yarpc-go/branch/master

[report-card-img]: https://goreportcard.com/badge/go.uber.org/yarpc
[report-card]: https://goreportcard.com/report/go.uber.org/yarpc
