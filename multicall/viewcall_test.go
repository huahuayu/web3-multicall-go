package multicall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewCall(t *testing.T) {
	vc := ViewCall{
		Id:        "key",
		Target:    "0x0",
		Method:    "balanceOf(address, uint64)(int256)",
		Arguments: []interface{}{"0x1234", uint64(12)},
	}
	expectedArgTypes := []string{"address", "uint64"}
	expectedCallData := []byte{
		0x29, 0x5e, 0xaa, 0xdf, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x12, 0x34, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0xc}
	assert.Equal(t, expectedArgTypes, vc.ArgumentTypes())
	callData, err := vc.CallData()
	assert.Nil(t, err)
	assert.Equal(t, expectedCallData, callData)
}

func TestCatchPanicOnInterfaceIssue(t *testing.T) {
	vc := ViewCall{
		Id:        "key",
		Target:    "0x0",
		Method:    "balanceOf(address)(int256)",
		Arguments: []interface{}{1234},
	}

	err := vc.Validate()
	assert.NotNil(t, err)
	assert.Error(t, err, "expected address argument to be a string")
}

func TestEncodeNumericArgument(t *testing.T) {
	vc1 := ViewCall{
		Id:        "key",
		Target:    "0x0",
		Method:    "balanceOf(uint256)(int256)",
		Arguments: []interface{}{"12312312312313"},
	}
	vc2 := ViewCall{
		Id:        "key",
		Target:    "0x0",
		Method:    "balanceOf(uint256)(int256)",
		Arguments: []interface{}{12312312312313},
	}

	data1, err1 := vc1.ArgsCallData()
	data2, err2 := vc2.ArgsCallData()
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, data1, data2)
}

func TestEncodeBytes32Argument(t *testing.T) {
	var bytes32Array = [32]uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	vc1 := ViewCall{
		Id:        "key",
		Target:    "0x0",
		Method:    "balanceOfPartition(bytes32, uint256)(int256)",
		Arguments: []interface{}{bytes32Array, "12312312312313"},
	}

	_, err1 := vc1.ArgsCallData()
	assert.Nil(t, err1)
}
