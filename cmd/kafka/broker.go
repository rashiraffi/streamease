package kafka

import (
	"fmt"

	"github.com/spf13/viper"
)

func getBrokerAddr(broker string, profile string) (hostAddr string, err error) {
	if broker != "" {
		hostAddr = broker
		return
	}
	if profile != "" {
		hostAddr, ok := viper.Get(fmt.Sprintf("kafka.profile.%s.host", profile)).(string)
		if !ok || hostAddr == "" {
			err = fmt.Errorf("profile %s not found", profile)
			return "", err
		} else {
			return hostAddr, nil
		}
	}

	defaultProfile, ok := viper.Get("kafka.default_profile").(string)
	if !ok || defaultProfile == "" {
		err = fmt.Errorf("default profile not found")
		return "", err
	} else {
		hostAddr, ok := viper.Get(fmt.Sprintf("kafka.profile.%s.host", defaultProfile)).(string)
		if !ok || hostAddr == "" {
			err = fmt.Errorf("profile %s not found", defaultProfile)
			return "", err
		} else {
			return hostAddr, nil
		}
	}

}
