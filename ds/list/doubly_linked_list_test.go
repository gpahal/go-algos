package list_test

import (
	"testing"

	"github.com/gpahal/go-algos/ds/list"
)

func TestDoublyLinkedList(t *testing.T) {
	testInterfaceHelper(t, list.NewDoublyLinkedList)
}
