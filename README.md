# [@szabado](https://github.com/szabado)'s go tooling monorepo

[![codecov](https://codecov.io/gh/szabado/go-tools/graph/badge.svg?token=0P63IKE7ZG)](https://codecov.io/gh/szabado/go-tools)
[![goreportcard](https://goreportcard.com/badge/github.com/szabado/go-tools)](https://goreportcard.com/report/github.com/szabado/go-tools)

This contains various tooling, organized in a monorepo for ease of management.

Tools:
- `cache`: A tool to cache the execution of command line tools.
- `mssh`: A tool to SSH into multiple servers and execute a command.
- `zkcli`: A Zookeeper CLI, designed for very large Zookeeper instances.

Everything in this repo is licensed under the MIT license, EXCEPT for the `cmd/zkcli/` folder which is released under the Apache 2.0 License.