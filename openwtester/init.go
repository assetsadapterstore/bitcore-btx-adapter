package openwtester

import (
	"github.com/assetsadapterstore/bitcore-btx-adapter/bitcore_btx"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(bitcore_btx.Symbol, bitcore_btx.NewWalletManager())
}
