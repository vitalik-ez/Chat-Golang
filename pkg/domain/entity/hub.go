package entity

type subscription struct {
	conn *connection
	room Room
}

type hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	send chan Message

	// Register requests from the connections.
	join chan subscription

	// Unregister requests from connections.
	leave chan subscription
}
