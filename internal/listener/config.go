package listener

import "fmt"

type ListenerConfig struct {
	// StartBlock specifies the block at which we start listening for event.
	// switch StartBlock:
	// 		case -1:
	//			LoadFromDB();
	//		case 0:
	//			StartFromLatest()
	//		default:
	//			StartRightWhereItIsDefined()
	StartBlock int64 `json:"StartBlock"`
}

func DefaultConfig() ListenerConfig {
	return ListenerConfig{
		StartBlock: 0,
	}
}
func (cfg ListenerConfig) IsValid() (bool, error) {
	if cfg.StartBlock < -1 {
		return false, fmt.Errorf("invalid StartBlock")
	}
	return true, nil
}
