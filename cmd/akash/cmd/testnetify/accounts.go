package testnetify

import (
	"fmt"

	"github.com/theckman/yacspin"

	"github.com/cosmos/cosmos-sdk/codec"
)

func (ga *GenesisState) modifyAccounts(sp *yacspin.Spinner, cdc codec.Codec, cfg *AccountsConfig) error {
	for _, acc := range cfg.Add {
		if err := ga.AddNewAccount(cdc, acc.Address.AccAddress, acc.PubKey.PubKey); err != nil {
			return err
		}

		if err := ga.IncreaseBalances(cdc, acc.Address.AccAddress, acc.Coins.ToSDK()); err != nil {
			return err
		}
	}

	sp.Message(fmt.Sprintf("added new accounts"))

	return nil
}
