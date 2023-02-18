package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmosmessenger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmosmessenger"

	// SenderKey uniquely defines messages where one is a sender
	senderKey = "/Messages/Sender/"

	// SenderKey uniquely defines messages where one is a receiver
	receiverKey = "/Messages/Receiver/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeySenderPrefix(senderName string) string {
	return senderKey + senderName
}

func KeyReceiverPrefix(receiverName string) string {
	return receiverKey + receiverName
}
