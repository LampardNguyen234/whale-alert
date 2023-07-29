package common

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"reflect"
)

type EVMEvent interface {
	Name() string
	GetLog() types.Log
}

type BaseEVMEvent struct {
	log types.Log
}

func (e *BaseEVMEvent) GetLog() types.Log {
	return e.log
}

// EVMEventHandler specifies method for an event handler.
type EVMEventHandler interface {
	Handle(event EVMEvent) error
}

// UnpackLog unpacks a retrieved log into the provided output structure.
func UnpackLog(contractAbi abi.ABI, out interface{}, event string, log types.Log) error {
	if log.Topics[0] != contractAbi.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}
	fmt.Println("UnpackLog", reflect.ValueOf(out).Kind())
	if len(log.Data) > 0 {
		if err := contractAbi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range contractAbi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
