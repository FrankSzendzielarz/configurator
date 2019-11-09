package cspec

// EIP represents the blockchain arguments at a
// given blockNumber. A list of these is a set
// of chainspec arguments, the ChainSpecArguments.
// The initial configuration is considered EIP0
type EIP struct {
	Enabled        bool //Enabled allows for explicit checks on presence
	Name           uint32
	BlockNumber    Number
	BlockReward    Number
	DifficultyBomb Number
	MaxCodeSize    Number

	// these are included because conceivably they are
	// blockchain arguments that can be altered in an
	// EIP.
	ChainID           Number
	NetworkID         Number
	AccountStartNonce Number
	GasLimit          Number
	MaxExtraDataSize  Number
}
