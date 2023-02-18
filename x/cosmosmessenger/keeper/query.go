package keeper

import (
	"cosmos-messenger/x/cosmosmessenger/types"
)

var _ types.QueryServer = Keeper{}
