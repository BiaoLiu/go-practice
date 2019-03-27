package unittest1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestM1(t *testing.T) {
	asserts:=assert.New(t)

	c:=M1(3,5)
	asserts.Equal(8,c)
}