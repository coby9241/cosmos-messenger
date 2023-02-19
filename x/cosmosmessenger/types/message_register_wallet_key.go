package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterWalletKey = "register_wallet_key"

var _ sdk.Msg = &MsgRegisterWalletKey{}

func NewMsgRegisterWalletKey(creator string, pubkey string) *MsgRegisterWalletKey {
	return &MsgRegisterWalletKey{
		Creator: creator,
		Pubkey:  pubkey,
	}
}

func (msg *MsgRegisterWalletKey) Route() string {
	return RouterKey
}

func (msg *MsgRegisterWalletKey) Type() string {
	return TypeMsgRegisterWalletKey
}

func (msg *MsgRegisterWalletKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterWalletKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterWalletKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
