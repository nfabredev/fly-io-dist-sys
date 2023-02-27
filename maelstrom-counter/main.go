package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	kv := maelstrom.NewSeqKV(n)
	ctx := context.TODO()

	n.Handle("add", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		rawDelta := body["delta"].(string)
		delta, err := strconv.Atoi(rawDelta)

		if err != nil {
			log.Fatal(err)
		}

		body["type"] = "add_ok"
		delete(body, "delta")

		counter, err := kv.ReadInt(ctx, "counter")
		if err != nil {
			log.Fatal(err)
		}
		kv.Write(ctx, "counter", counter+delta)

		// messages = append(messages, rawMessage)
		return n.Reply(msg, body)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		counter, err := kv.Read(ctx, "counter")
		if err != nil {
			log.Fatal(err)
		}
		body["type"] = "read_ok"
		body["value"] = counter

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
