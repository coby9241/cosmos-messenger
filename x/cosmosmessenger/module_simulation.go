package cosmosmessenger

import (
	"math/rand"

	"cosmos-messenger/testutil/sample"
	cosmosmessengersimulation "cosmos-messenger/x/cosmosmessenger/simulation"
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cosmosmessengersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateMessage = "op_weight_msg_create_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMessage int = 100

	opWeightMsgRegisterWalletKey = "op_weight_msg_register_wallet_key"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterWalletKey int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cosmosmessengerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cosmosmessengerGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMessage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMessage, &weightMsgCreateMessage, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMessage = defaultWeightMsgCreateMessage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMessage,
		cosmosmessengersimulation.SimulateMsgCreateMessage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterWalletKey int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterWalletKey, &weightMsgRegisterWalletKey, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterWalletKey = defaultWeightMsgRegisterWalletKey
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterWalletKey,
		cosmosmessengersimulation.SimulateMsgRegisterWalletKey(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
