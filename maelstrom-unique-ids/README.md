## Challenge #2: Unique ID Generation

https://fly.io/dist-sys/2/

### Dependencies

- Go
- [Maelstrom 0.2.2](https://github.com/jepsen-io/maelstrom/releases/tag/v0.2.2)

### Build

```sh
go install .
```

### Test

`./maelstrom test -w unique-ids --bin $GOPATH/bin/maelstrom-unique-ids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition`
