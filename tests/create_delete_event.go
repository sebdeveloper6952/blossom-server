package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nbd-wtf/go-nostr"
	"log"
	"time"
)

func main() {
	sk := "f7023ed0821694ea3eeebd271aa7ee1fb5e1ace94ef88fc116f291276c68ebc3"
	ev := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      24242,
		Tags: nostr.Tags{
			{"expiration", fmt.Sprintf("%d", time.Now().Add(time.Hour*24).Unix())},
			{"t", "delete"},
			{"x", "c402f0974e2f6ebe96efee967c64c3ebfd4366e2f284f4d8650371af7787fdb0"},
		},
	}
	ev.Sign(sk)

	bytes, err := json.Marshal(ev)
	if err != nil {
		log.Fatal(err)
	}

	b64 := base64.RawURLEncoding.EncodeToString(bytes)

	fmt.Println(b64)
}
