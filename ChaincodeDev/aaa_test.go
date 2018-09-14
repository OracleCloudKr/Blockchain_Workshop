package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestQuery2(t *testing.T) {
	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleAsset))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}
	args := [][]byte{[]byte("set"), []byte("A"), []byte("100")}
	invokeResult := stub.MockInvoke("1", args)

	args = [][]byte{[]byte("get"), []byte("A")}
	invokeResult = stub.MockInvoke("1", args)
	fmt.Printf("Payload=%s", invokeResult.GetPayload())

	if invokeResult.Status != 200 {
		t.Errorf("returned non-OK status, got: %d, want: %d.", invokeResult.Status, 200)
	}
}
