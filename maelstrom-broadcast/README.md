## Challenge #3a: Single-Node Broadcast

https://fly.io/dist-sys/3/

### Dependencies

- Go
- [Maelstrom 0.2.2](https://github.com/jepsen-io/maelstrom/releases/tag/v0.2.2)

### Build

```sh
go install .
```

### Test

`./maelstrom test -w broadcast --bin $GOPATH/bin/maelstrom-broadcast --node-count 1 --time-limit 20 --rate 10`
