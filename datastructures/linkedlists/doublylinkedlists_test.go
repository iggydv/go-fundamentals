package linkedlists

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedList_Prepend(t *testing.T) {
	type fields struct {
		head   *DoubleNode
		tail   *DoubleNode
		length int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Insert value at front",
			fields: fields{
				head: &DoubleNode{
					value: 8,
					next:  nil,
					prev:  nil,
				},
				length: 0,
			},
			args: args{
				value: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &DoublyLinkedList{
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				length: tt.fields.length,
			}
			dl.Prepend(tt.args.value)
			if dl.head.value != 22 {
				fmt.Errorf("DoublyLinkedList.Prepend() = %v, want %v", dl.head.value, 22)
			}
		})
	}
}

func TestDoublyLinkedList_Append(t *testing.T) {
	type fields struct {
		head   *DoubleNode
		tail   *DoubleNode
		length int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Insert value at end",
			fields: fields{
				head: &DoubleNode{
					value: 8,
					next:  nil,
					prev:  nil,
				},
				length: 0,
			},
			args: args{
				value: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &DoublyLinkedList{
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				length: tt.fields.length,
			}
			dl.Append(tt.args.value)
			current := dl.tail
			for current.next != nil {
				current = dl.tail.next
			}
			if current.value != 22 {
				fmt.Errorf("DoublyLinkedList.Append() = %v, want %v", current.value, 22)
			}
		})
	}
}
