package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMessage = "create_message"

var _ sdk.Msg = &MsgCreateMessage{}

func NewMsgCreateMessage(creator string, receiver string, body string) *MsgCreateMessage {
	return &MsgCreateMessage{
		Creator:  creator,
		Receiver: receiver,
		Body:     body,
	}
}

func (msg *MsgCreateMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreateMessage) Type() string {
	return TypeMsgCreateMessage
}

func (msg *MsgCreateMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
