package stack_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/stack"
)

func TestListStack(t *testing.T) {
	testInterfaceHelper(t, stack.NewListStack)
}
