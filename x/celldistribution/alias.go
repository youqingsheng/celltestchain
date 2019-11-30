// overwrite the distribution
package celldistribution

import (
	cellkeeper "github.com/cell-network/cellchain/x/celldistribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const (
	DefaultParamspace = keeper.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	QueryParams       = types.QueryParams
)

var (
	NewKeeper = cellkeeper.NewKeeper

	BeginBlocker  = distribution.BeginBlocker
	InitGenesis   = distribution.InitGenesis
	ExportGenesis = distribution.ExportGenesis

	// functions aliases
	RegisterInvariants  = keeper.RegisterInvariants
	ParamKeyTable       = keeper.ParamKeyTable
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc              = types.ModuleCdc
	AttributeValueCategory = types.AttributeValueCategory
)

type (
	Hooks                                  = keeper.Hooks
	Keeper                                 = cellkeeper.Keeper
	DelegatorStartingInfo                  = types.DelegatorStartingInfo
	CodeType                               = types.CodeType
	FeePool                                = types.FeePool
	DelegatorWithdrawInfo                  = types.DelegatorWithdrawInfo
	ValidatorOutstandingRewardsRecord      = types.ValidatorOutstandingRewardsRecord
	ValidatorAccumulatedCommissionRecord   = types.ValidatorAccumulatedCommissionRecord
	ValidatorHistoricalRewardsRecord       = types.ValidatorHistoricalRewardsRecord
	ValidatorCurrentRewardsRecord          = types.ValidatorCurrentRewardsRecord
	DelegatorStartingInfoRecord            = types.DelegatorStartingInfoRecord
	ValidatorSlashEventRecord              = types.ValidatorSlashEventRecord
	GenesisState                           = types.GenesisState
	MsgSetWithdrawAddress                  = types.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward             = types.MsgWithdrawDelegatorReward
	MsgWithdrawValidatorCommission         = types.MsgWithdrawValidatorCommission
	CommunityPoolSpendProposal             = types.CommunityPoolSpendProposal
	QueryValidatorOutstandingRewardsParams = types.QueryValidatorOutstandingRewardsParams
	QueryValidatorCommissionParams         = types.QueryValidatorCommissionParams
	QueryValidatorSlashesParams            = types.QueryValidatorSlashesParams
	QueryDelegationRewardsParams           = types.QueryDelegationRewardsParams
	QueryDelegatorParams                   = types.QueryDelegatorParams
	QueryDelegatorWithdrawAddrParams       = types.QueryDelegatorWithdrawAddrParams
	QueryDelegatorTotalRewardsResponse     = types.QueryDelegatorTotalRewardsResponse
	DelegationDelegatorReward              = types.DelegationDelegatorReward
	ValidatorHistoricalRewards             = types.ValidatorHistoricalRewards
	ValidatorCurrentRewards                = types.ValidatorCurrentRewards
	ValidatorAccumulatedCommission         = types.ValidatorAccumulatedCommission
	ValidatorSlashEvent                    = types.ValidatorSlashEvent
	ValidatorSlashEvents                   = types.ValidatorSlashEvents
	ValidatorOutstandingRewards            = types.ValidatorOutstandingRewards
)
