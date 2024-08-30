package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/covenet module sentinel errors
var (
	ErrSample                        = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrInvalidSourceChainId          = errorsmod.Register(ModuleName, 1101, "submitted source chain-id is invalid, accepts only unsigned integers")
	ErrInvalidProofType              = errorsmod.Register(ModuleName, 1102, "submitted proof-type is invalid, accepts only 'specimen', 'result'")
	ErrInvalidBlockHeight            = errorsmod.Register(ModuleName, 1103, "submitted block-height is invalid, accepts only unsigned integers")
	ErrInvalidBlockHash              = errorsmod.Register(ModuleName, 1104, "submitted block-hash is invalid, accepts only hash string")
	ErrInvalidProofHash              = errorsmod.Register(ModuleName, 1105, "submitted proof-hash is invalid, accepts only string")
	ErrInvalidStoreAddress           = errorsmod.Register(ModuleName, 1106, "submitted store-address is invalid, accepts only string")
	ErrInvalidDeadline               = errorsmod.Register(ModuleName, 1107, "deadline cannot be parsed, accepts only utc unix string time")
	ErrInvalidProofCount             = errorsmod.Register(ModuleName, 1108, "submitted proof-count is invalid, accepts only unsigned integers")
	ErrInvalidProofSessionId         = errorsmod.Register(ModuleName, 1109, "submitted proof-session-id is invalid")
	ErrAlreadySubmitted              = errorsmod.Register(ModuleName, 1110, "submitted tx creator is already session member")
	ErrChainRoleExists               = errorsmod.Register(ModuleName, 1111, "submitted chain role exists for chain id")
	ErrEmptyAddress                  = errorsmod.Register(ModuleName, 1112, "submitted address cannot be empty")
	ErrInvalidOperatorAddress        = errorsmod.Register(ModuleName, 1113, "submitted operator address is an invalid Cosmos address")
	ErrInvalidSequencerAddress       = errorsmod.Register(ModuleName, 1114, "submitted Sequencer address is invalid")
	ErrInvalidValidatorAddress       = errorsmod.Register(ModuleName, 1115, "submitted Validator address is an invalid Ethereum address")
	ErrChainRoleDoesNotExist         = errorsmod.Register(ModuleName, 1116, "provided chain role does not exist")
	ErrProofSessionValidation        = errorsmod.Register(ModuleName, 1117, "proof session transaction validation failed")
	ErrCannotFindExpiredProofSession = errorsmod.Register(ModuleName, 1118, "expired proof session cannot be found")
	ErrProofOperatorNotFound         = errorsmod.Register(ModuleName, 1119, "proof operator not found in chain roles")
	ErrProofValidatorNotFound        = errorsmod.Register(ModuleName, 1120, "proof validator not found in chain roles")
	ErrProofSequencerNotFound        = errorsmod.Register(ModuleName, 1121, "proof sequencer not found in chain roles")
	ErrChainIdNotSupported           = errorsmod.Register(ModuleName, 1122, "submitted chain id not supported with chain roles")
	ErrChainConfigExists             = errorsmod.Register(ModuleName, 1123, "submitted chain config already exists for chain id")
	ErrChainConfigDoesNotExist       = errorsmod.Register(ModuleName, 1124, "provided chain config does not exist")
	ErrInvalidQuorumReq              = errorsmod.Register(ModuleName, 1125, "provided quorum is not usable")
	ErrInvalidSrcBlockTimeSec        = errorsmod.Register(ModuleName, 1126, "provided source block time is not usable")
	ErrInvalidSnkBlockTimeSec        = errorsmod.Register(ModuleName, 1127, "provided sink block time is not usable")
	ErrInvalidRewardAlloc            = errorsmod.Register(ModuleName, 1128, "provided reward allocation per session is not usable")
	ErrInvalidMinStakeReq            = errorsmod.Register(ModuleName, 1129, "provided min stake required is not usable")
	ErrInvalidMinSubmissionReq       = errorsmod.Register(ModuleName, 1130, "provided min proof submissions required is not usable")
	ErrInvalidSessionTimeSec         = errorsmod.Register(ModuleName, 1131, "provided session duration is not usable")
	ErrInvalidSecondTimeSuffix       = errorsmod.Register(ModuleName, 1132, "provided duration is not usable without `s` suffix")
	ErrBigIntString                  = errorsmod.Register(ModuleName, 1133, "provided big int string is not usable")
	ErrSubRoleCannotBeEmpty          = errorsmod.Register(ModuleName, 1134, "cannot remove all the sub-roles for this chain-id")
	ErrIncorrectChainConfig          = errorsmod.Register(ModuleName, 1135, "invalid chain configuration for proof creation")
	ErrInvalidBoolean                = errorsmod.Register(ModuleName, 1136, "provided value is an invalid boolean")
	ErrStructFieldIsEmpty            = errorsmod.Register(ModuleName, 1137, "provided value(s) to the struct cannot be empty")
	ErrProofOutOfBounds              = errorsmod.Register(ModuleName, 1138, "proof session submitted out of acceptable live bounds")
	ErrChainProofTypeNotSupported    = errorsmod.Register(ModuleName, 1139, "provided proof type value is not supported")
	ErrValidatorNotEnabled           = errorsmod.Register(ModuleName, 1140, "provided proof validator is not enabled")

	ErrInt32OverFlow = errorsmod.Register(ModuleName, 0001, "submitted int32 is invalid, should be between 0 and 2147483647")
	ErrInt64OverFlow = errorsmod.Register(ModuleName, 0002, "submitted int64 is invalid, should be between 0 and 18446744073709551615")
	ErrStringNotUrl  = errorsmod.Register(ModuleName, 0003, "submitted string is invalid url, should contain a hostname and path to resource")
)
