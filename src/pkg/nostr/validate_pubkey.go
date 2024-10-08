package nostr

import goNostr "github.com/nbd-wtf/go-nostr"

func IsValidPubkey(pk string) bool {
	return goNostr.IsValidPublicKey(pk)
}
