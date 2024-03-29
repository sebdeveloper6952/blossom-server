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
	sk := nostr.GeneratePrivateKey()
	ev := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      24242,
		Tags: nostr.Tags{
			{"expiration", fmt.Sprintf("%d", time.Now().Add(time.Hour*24).Unix())},
			{"t", "upload"},
			{"size", "36194"},
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
