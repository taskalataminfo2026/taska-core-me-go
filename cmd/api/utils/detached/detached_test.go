package detached_test

import (
	"context"
	"taska-core-me-go/cmd/api/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDetached(t *testing.T) {
	assert := assert.New(t)

	t.Run("It should not be cancelled the new context when the original context is cancelled", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "field_1", "value_1")
		ctx, cancel := context.WithCancel(ctx)
		newCtx := utils.NewContextFlow(ctx)
		cancel()

		assert.Equal(newCtx.Value("field_1"), "value_1", "they should be equal")

		deadLine, isDead := newCtx.Deadline()
		assert.Equal(deadLine, time.Time{}, "they should be equal")
		assert.Equal(isDead, false, "they should be equal")

		err := newCtx.Err()
		assert.NoError(err, "they should not be equal error")

		done := newCtx.Done()
		assert.Nil(done, "they should be equal nil")
	})
}
