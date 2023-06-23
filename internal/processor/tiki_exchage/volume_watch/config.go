package volume_watch

import (
	"fmt"
	"github.com/LampardNguyen234/whale-alert/internal/common"
	"time"
)

type VolumeCheck struct {
	Period time.Duration
	Volume float64
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
				30 * time.Minute,
				100000,
			},
			{
				24 * time.Hour,
				1000000,
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
	for _, volume := range cfg.VolumeChecks {
		if volume.Volume == 0 {
			return false, fmt.Errorf("invalid volume for duration %v", volume.Period.String())
		}
	}

	return true, nil
}
