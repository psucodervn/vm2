package pm2

import (
  "testing"
  "github.com/stretchr/testify/assert"
  )

func TestList(t *testing.T) {
  as := assert.New(t)

  procs, err := List()
  as.NoError(err)
  as.IsType(new([]Process), &procs)
}
