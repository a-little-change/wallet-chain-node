package fallback

import (
	"github.com/SavourDao/savour-hd/rpc/common"
	"github.com/SavourDao/savour-hd/rpc/wallet"
)

type WalletAdaptor struct{}

func (w *WalletAdaptor) GetSupportCoins(request *wallet.SupportCoinsRequest) (*wallet.SupportCoinsResponse, error) {
	panic("implement me")
}

func (w *WalletAdaptor) GetNonce(request *wallet.NonceRequest) (*wallet.NonceResponse, error) {
	return &wallet.NonceResponse{
		Code:  common.ReturnCode_ERROR,
		Nonce: "",
	}, nil
}

func (a *WalletAdaptor) GetGasPrice(req *wallet.GasPriceRequest) (*wallet.GasPriceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *WalletAdaptor) SendTx(req *wallet.SendTxRequest) (*wallet.SendTxResponse, error) {
	//TODO implement me
	panic("implement me")
}