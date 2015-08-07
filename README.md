Shampoo
=======

Shampoo flaky tests. Runs the given command `n` times, failing if any iteration fails. Inspired by https://gist.github.com/JakeWharton/7fe7deb1f7f4a795c120.

Installation
============

`go get github.com/f2prateek/shampoo`


Usage
======

```
Shampoo.

Shampoo away your flaky tests.

Usage:
  shampoo [--iterations=<iterations>] [--parallel] <cmd> <args>...
  shampoo -h | --help
  shampoo --version

Options:
  -h --help                  Show this screen.
  --version                  Show version.
  --iterations=<iterations>  Number of iterations to run [default: 10].
  --parallel                 Run in parallel.
```