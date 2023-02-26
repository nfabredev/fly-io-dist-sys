## Challenge #1: Echo

https://fly.io/dist-sys/1/

### Dependencies

- Go
- [Maelstrom 0.2.2](https://github.com/jepsen-io/maelstrom/releases/tag/v0.2.2)

### Build

```sh
go install .
```

### Test

`./maelstrom test -w echo --bin $GOPATH/bin/maelstrom-echo --node-count 1 --time-limit 10`
