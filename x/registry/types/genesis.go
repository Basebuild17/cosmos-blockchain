package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DataList: []Data{},
	}
}

// Validate performs basic validation of GenesisState fields
func (gs GenesisState) Validate() error {
	// validate data entries
	for _, data := range gs.DataList {
		if data.Key == "" {
			return ErrInvalidKey
		}
		if data.Value == "" {
			return ErrInvalidValue
		}
	}
	return nil
}
