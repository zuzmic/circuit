package circuit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneralConfig_Merge(t *testing.T) {

	t.Run("respect ForceOpen field of args cfg", func(t *testing.T) {
		cfg := GeneralConfig{}

		cfg.merge(GeneralConfig{ForceOpen: true})

		assert.True(t, cfg.ForceOpen, "expect to be true")
	})

	t.Run("respect ForceOpen field of receiver cfg", func(t *testing.T) {
		cfg := GeneralConfig{ForceOpen: true}

		cfg.merge(GeneralConfig{ForceOpen: false})

		assert.True(t, cfg.ForceOpen, "expect to be true")
	})

	t.Run("respect ForceClosed field of args cfg", func(t *testing.T) {
		cfg := GeneralConfig{}

		cfg.merge(GeneralConfig{ForcedClosed: true})

		assert.True(t, cfg.ForcedClosed, "expect to be true")
	})

	t.Run("respect ForceClosed field of receiver cfg", func(t *testing.T) {
		cfg := GeneralConfig{ForcedClosed: true}

		cfg.merge(GeneralConfig{ForceOpen: false})

		assert.True(t, cfg.ForcedClosed, "expect to be true")
	})

	t.Run("badRequestFunc not empty after merge with default config", func(t *testing.T) {
		cfg := GeneralConfig{BadRequestFunc: nil}

		cfg.merge(defaultGoSpecificConfig)

		assert.NotNil(t, cfg.BadRequestFunc, "should not be nil")
	})
}
