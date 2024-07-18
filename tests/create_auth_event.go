package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func main() {
	sk := "f7023ed0821694ea3eeebd271aa7ee1fb5e1ace94ef88fc116f291276c68ebc3"
	ev := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      24242,
		Tags: nostr.Tags{
			{"t", "upload"},
			{"x", "b0153d7fa0023728e9b496fb82504c057eef3512a3704434e922363a370090a3"},
			{"expiration", fmt.Sprintf("%d", time.Now().Add(time.Hour*24).Unix())},
		},
	}
	ev.Sign(sk)

	bytes, err := json.Marshal(ev)
	if err != nil {
		log.Fatal(err)
	}

	b64 := base64.StdEncoding.EncodeToString(bytes)

	fmt.Println(b64)
}
