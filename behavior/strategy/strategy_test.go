package strategy

import "testing"

func TestStrategy(t *testing.T) {

	c := &Context{}
	c.SetStrategy(&ConcreteStrategyAdd{})
	c.ExecuteStrategy(1, 2)
}
