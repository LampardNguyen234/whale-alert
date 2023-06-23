package volume_watch

import (
	"encoding/json"
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"time"
)

type VolumeCheck struct {
	Period time.Duration
	Volume float64
}

func (cfg *VolumeCheck) MarshalJSON() ([]byte, error) {
	type strHolder struct {
		Period string  `json:"Period"`
		Volume float64 `json:"Volume"`
	}
	tmp := strHolder{
		Period: cfg.Period.String(),
		Volume: cfg.Volume,
	}

	return json.Marshal(tmp)
}

func (cfg *VolumeCheck) UnmarshalJSON(data []byte) error {
	type strHolder struct {
		Period string  `json:"Period"`
		Volume float64 `json:"Volume"`
	}
	var tmp strHolder
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	period, err := time.ParseDuration(tmp.Period)
	if err != nil {
		return fmt.Errorf("failed to parse period %v: %v", tmp.Period, err)
	}

	cfg.Period = period
	cfg.Volume = tmp.Volume

	return nil
}

type VolumeWatchConfig struct {
	common.BaseConfig
	VolumeChecks []*VolumeCheck
}

func DefaultConfig() VolumeWatchConfig {
	return VolumeWatchConfig{
		BaseConfig: common.BaseConfig{
			Enabled:   true,
			QueueSize: 1024,
		},
		VolumeChecks: []*VolumeCheck{
			{
				24 * time.Hour,
				1000000,
			},
			{
				1 * time.Hour,
				200000,
			},
			{
				30 * time.Minute,
				100000,
			},
			{
				5 * time.Minute,
				50000,
			},
		},
	}
}

func (cfg VolumeWatchConfig) IsValid() (bool, error) {
	if _, err := cfg.BaseConfig.IsValid(); err != nil {
		return false, err
	}
	if len(cfg.VolumeChecks) == 0 {
		return false, fmt.Errorf("empty volume")
	}
	currentPeriod := 1000 * time.Hour
	for _, volume := range cfg.VolumeChecks {
		if volume.Period > currentPeriod {
			return false, fmt.Errorf("period must be descendingly sorted")
		}
		if volume.Volume == 0 {
			return false, fmt.Errorf("invalid volume for duration %v", volume.Period.String())
		}
		currentPeriod = volume.Period
	}

	return true, nil
}
