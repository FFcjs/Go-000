package Week06

import (
	"testing"
	"time"
)
import "github.com/stretchr/testify/assert"

func TestNewCounter(t *testing.T) {
	c:=NewCounter(5)
	for i:=float64(1);i<=10;i++{
		c.Add(i)
		time.Sleep(time.Second)
	}
	assert.Equal(t,c.Sum,40.0)
	assert.Equal(t,c.GetMax(),10.0)
	assert.Equal(t,c.GetMin(),6.0)
}
