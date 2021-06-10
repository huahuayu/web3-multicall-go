package multicall_test

import (
	"encoding/json"
	"fmt"
	"github.com/alethio/web3-go/ethrpc"
	"github.com/alethio/web3-go/ethrpc/provider/httprpc"
	"github.com/huahuayu/web3-multicall-go/multicall"
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
		"balanceOf(address)(uint256)",
		[]interface{}{"0x2f0b23f53734252bda2277357e97e1517d6b042a"},
	)
	vc3 := multicall.NewViewCall(
		"key3",
		"0xffa98a091331df4600f87c9164cd27e8a5cd2405",
		"getReserves()(uint112, uint112, uint32)",
		[]interface{}{},
	)

	vcs := multicall.ViewCalls{vc1, vc2, vc3}
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
	fmt.Println(res)
	fmt.Println(err)
}

func getETH(url string) (ethrpc.ETHInterface, error) {
	provider, err := httprpc.New(url)
	if err != nil {
		return nil, err
	}
	provider.SetHTTPTimeout(5 * time.Second)
	return ethrpc.New(provider)
}
