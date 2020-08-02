package actor

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ComponentConfig map[string]VariableConfig

type VariableConfig map[interface{}]interface{}

func (m VariableConfig) GetString(key string) (string, error) {
	v, ok := m[key].(string)
	if !ok {
		return "", fmt.Errorf("error while converting key '%s' to string", key)
	}

	return v, nil
}

func (m VariableConfig) GetStringOr(key string, fallback string) string {
	v, ok := m[key].(string)
	if !ok {
		return fallback
	}

	return v
}

func (m VariableConfig) GetInt(key string) (int, error) {
	v, ok := m[key].(int)
	if !ok {
		return 0, fmt.Errorf("error while converting key '%s' to int", key)
	}

	return v, nil
}

func (m VariableConfig) GetIntOr(key string, fallback int) int {
	v, ok := m[key].(int)
	if !ok {
		return fallback
	}

	return v
}

func (m VariableConfig) GetFloat64(key string) (float64, error) {
	v, ok := m[key].(float64)
	if !ok {
		return 0, fmt.Errorf("error while converting key '%s' to float64", key)
	}

	return v, nil
}

func (m VariableConfig) GetFloat64Or(key string, fallback float64) float64 {
	v, ok := m[key].(float64)
	if !ok {
		return fallback
	}

	return v
}

func (m VariableConfig) GetVariableConfig(key string) (VariableConfig, error) {
	v, ok := m[key].(VariableConfig)
	if !ok {
		return nil, fmt.Errorf("error while converting key '%s' to VariableConfig", key)
	}

	return v, nil
}

func (m VariableConfig) GetVariableConfigOr(key string, fallback VariableConfig) VariableConfig {
	v, ok := m[key].(VariableConfig)
	if !ok {
		return fallback
	}

	return v
}

func (m VariableConfig) Extract(key string, v interface{}) error {
	config, err := m.GetVariableConfig(key)
	if err != nil {
		return nil
	}

	err = mapstructure.Decode(config, v)

	if err != nil {
		return fmt.Errorf("invalid config for %v at key '%s'", v, key)
	}

	return nil
}
