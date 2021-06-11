package multicall_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alethio/web3-go/ethrpc"
	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huahuayu/web3-multicall-go/multicall"
	"math/big"
	"testing"
	"time"
)

func TestExampleViwCall(t *testing.T) {
	eth, err := getETH("https://mainnet.infura.io/v3/17ed7fe26d014e5b9be7dfff5368c69d")
	vc1 := multicall.NewViewCall(
		"key1",
		"0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643",
		"totalReserves()(uint256)",
		[]interface{}{},
	)
	vc2 := multicall.NewViewCall(
		"key2",
		"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		"balanceOf(address) (uint)",
		[]interface{}{"0x2f0b23f53734252bda2277357e97e1517d6b042a"},
	)
	vc3 := multicall.NewViewCall(
		"key3",
		"0xffa98a091331df4600f87c9164cd27e8a5cd2405",
		"getReserves()(uint112, uint112, uint32)",
		[]interface{}{},
	)
	vc4 := multicall.NewViewCall(
		"key4",
		"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		"transfer(address,uint256)(bool)",
		[]interface{}{"0x000000000000000000000000000000000000dead", "10"},
	)

	vcs := multicall.ViewCalls{vc1, vc2, vc3, vc4}
	mc, _ := multicall.New(eth)
	block := "latest"
	res, err := mc.Call(vcs, block)
	if err != nil {
		panic(err)
	}
	result := res.Calls["key2"]
	marshalJSON, _ := result.Decoded[0].(*multicall.BigIntJSONString).MarshalJSON()
	fmt.Println("result: ", string(marshalJSON))

	resJson, _ := json.Marshal(res)
	fmt.Println(string(resJson))
}

func getETH(url string) (ethrpc.ETHInterface, error) {
	provider, err := httprpc.New(url)
	if err != nil {
		return nil, err
	}
	provider.SetHTTPTimeout(5 * time.Second)
	return ethrpc.New(provider)
}

func TestTransfer(t *testing.T) {
	ethClient, _ := ethclient.DialContext(context.Background(), "https://mainnet.infura.io/v3/17ed7fe26d014e5b9be7dfff5368c69d")
	vc := multicall.NewViewCall(
		"0",
		"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		"transfer(address,uint256)(bool)",
		[]interface{}{"0x000000000000000000000000000000000000dead", "10"},
	)
	callData, _ := vc.CallData()
	fmt.Println(hexutil.Encode(callData))
	to := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	msg := ethereum.CallMsg{From: common.HexToAddress("0xDC39546D5eB7b7Db48E04BdC98067603d6112081"), To: &to, Gas: 90000, GasPrice: big.NewInt(10000000000), Value: big.NewInt(0), Data: callData}
	bs0, err := ethClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		t.Fatal("call: ", err)
	}
	decode0, _ := vc.Decode(bs0)
	result := decode0[0].(bool)
	fmt.Println(result)
}
