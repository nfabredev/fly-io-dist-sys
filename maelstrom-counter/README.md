## Challenge #4: Grow-Only Counter

https://fly.io/dist-sys/4/

### Dependencies

- Go
- [Maelstrom 0.2.2](https://github.com/jepsen-io/maelstrom/releases/tag/v0.2.2)

### Build

```sh
go install .
```

### Test

`./maelstrom test -w g-counter --bin $GOPATH/bin/maelstrom-counter --node-count 3 --rate 100 --time-limit 20 --nemesis partition`
