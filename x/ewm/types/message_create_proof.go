package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateProof = "create_proof"

var _ sdk.Msg = &MsgCreateProof{}

func NewMsgCreateProof(creator string, chainId int32, proofType string, blockHeight uint64, blockHash string, proofHash string, storeAddress string) *MsgCreateProof {
	return &MsgCreateProof{
		Creator:      creator,
		ChainId:      chainId,
		ProofType:    proofType,
		BlockHeight:  blockHeight,
		BlockHash:    blockHash,
		ProofHash:    proofHash,
		StoreAddress: storeAddress,
	}
}

func (msg *MsgCreateProof) Type() string {
	return TypeMsgCreateProof
}

func (msg *MsgCreateProof) GetSigners() []sdk.AccAddress {
	creator, err := CovenetAccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProof) ValidateBasic() error {
	_, err := CovenetAccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator account address (%s)", err)
	}

	err = AssertInt32(msg.ChainId)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidSourceChainId, "invalid source chain-id (%s)", err)
	}

	err = AssertProofType(msg.ProofType)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidProofType, "invalid proof-type (%s)", err)
	}

	err = AssertInt64(msg.BlockHeight)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidBlockHeight, "invalid block-height (%s)", err)
	}

	err = AssertHashLength(msg.BlockHash)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidBlockHash, "invalid block-hash (%s)", err)
	}

	err = AssertHashLength(msg.ProofHash)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidProofHash, "invalid proof-hash (%s)", err)
	}

	err = AssertUrl(msg.StoreAddress)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidStoreAddress, "invalid store-address (%s)", err)
	}

	return nil
}
