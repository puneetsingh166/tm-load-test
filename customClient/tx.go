package main
import (
  banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
  "github.com/cosmos/cosmos-sdk/simapp"
  "fmt"
  "github.com/cosmos/cosmos-sdk/types/tx/signing"
  xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
  "github.com/cosmos/cosmos-sdk/client/tx"
  cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
  "github.com/cosmos/cosmos-sdk/testutil/testdata"
  sdk "github.com/cosmos/cosmos-sdk/types"
)


const (
  // Bech32PrefixAccAddr defines the Bech32 prefix of an account's address.
  Bech32PrefixAccAddr = "onomy"
  // Bech32PrefixAccPub defines the Bech32 prefix of an account's public key.
  Bech32PrefixAccPub = "onomypub"
  // Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address.
  Bech32PrefixValAddr = "onomyvaloper"
  // Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key.
  Bech32PrefixValPub = "onomyvaloperpub"
  // Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address.
  Bech32PrefixConsAddr = "onomyvalcons"
  // Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key.
  Bech32PrefixConsPub = "onomyvalconspub"
)

func init() { // nolint:gochecknoinits
  config := sdk.GetConfig()
  config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
  config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
  config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}

func GenTx() ([]byte){
  encCfg := simapp.MakeTestEncodingConfig()
  
  // Create a new TxBuilder
  txBuilder := encCfg.TxConfig.NewTxBuilder()
  priv1, _, addr1 := testdata.KeyTestPubAddr()
  _, _, addr2 := testdata.KeyTestPubAddr()
  msg1 := banktypes.NewMsgSend(addr1, addr2, sdk.NewCoins(sdk.NewInt64Coin("footoken", 1)))
  err := txBuilder.SetMsgs(msg1)
  if err != nil {
	  fmt.Println("error")
  }

  txBuilder.SetGasLimit(200000)
  txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewInt64Coin("footoken", 5)))
  txBuilder.SetMemo("footoken")
  txBuilder.SetTimeoutHeight(5)
  //------------------------------- ------Signing a Transaction---------------------------------//
  
  privs := []cryptotypes.PrivKey{priv1}
  var sigsV2 []signing.SignatureV2
   for _, priv := range privs {
	  sigV2 := signing.SignatureV2{
		  PubKey: priv.PubKey(),
		  Data: &signing.SingleSignatureData{
			  SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
			  Signature: nil,
		  },
		  
		  Sequence: 0,
	  }
	   sigsV2 = append(sigsV2, sigV2)
  }
  err = txBuilder.SetSignatures(sigsV2...)
  if err != nil {
	  fmt.Println("error - ",err)
  }
  sigsV2 = []signing.SignatureV2{}

  for _, priv := range privs {
	  signerData := xauthsigning.SignerData{
		  ChainID:       "onomy",
		  AccountNumber: 0,
		  Sequence:      0,
	  }
	  sigV2, err := tx.SignWithPrivKey(
		  encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		  txBuilder, priv, encCfg.TxConfig, 0)
		  
	if err != nil {
		  fmt.Println("error - ",err)
	  }
	  sigsV2 = append(sigsV2, sigV2)
  }
  
  err = txBuilder.SetSignatures(sigsV2...)
  if err != nil {
	  fmt.Println("error - ",err)
  }

  txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
  if err != nil {
    fmt.Println("error - ",err)
  }
  return txBytes
}
