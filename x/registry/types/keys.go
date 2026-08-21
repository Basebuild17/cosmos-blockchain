package types

const (
	// ModuleName defines the module name
	ModuleName = "registry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_" + ModuleName

	// EventTypeSetData defines the event type for setting data
	EventTypeSetData = "set_data"

	// EventTypeGetData defines the event type for getting data
	EventTypeGetData = "get_data"

	// AttributeKeyCreator defines the creator attribute
	AttributeKeyCreator = "creator"

	// AttributeKeyKey defines the key attribute
	AttributeKeyKey = "key"

	// AttributeKeyValue defines the value attribute
	AttributeKeyValue = "value"
)
