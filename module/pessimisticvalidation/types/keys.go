package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "pessimisticvalidation"

	// StoreKey is the store key string for pessimisticvalidation
	StoreKey = ModuleName
)

var (
	// ParamsKey is the prefix for pessimisticvalidation parameters
	ParamsKey = collections.NewPrefix(0)
)
