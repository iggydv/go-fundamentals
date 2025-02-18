package linkedlists

import "testing"

func TestCircularLinkedList_Prepend(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		value int
	}
	first := &Node{
		value: 8,
		next:  nil,
	}
	first.next = first

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Insert value at front",
			fields: fields{
				head:   first,
				length: 1,
			},
			args: args{
				value: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CircularLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			cl.Prepend(tt.args.value)
			if cl.head.value != 22 {
				t.Errorf("CircularLinkedList.Prepend() = %v, want %v", cl.head.value, 22)
			}
			if cl.head.next.value != 8 {
				t.Errorf("CircularLinkedList.Prepend() = %v, want %v", cl.head.next.value, 8)
			}
		})
	}
}

func TestCircularLinkedList_Append(t *testing.T) {
	type fields struct {
		head   *Node
		length int
	}
	type args struct {
		value int
	}
	first := &Node{
		value: 8,
		next:  nil,
	}
	first.next = first
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Insert value at front",
			fields: fields{
				head:   first,
				length: 1,
			},
			args: args{
				value: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CircularLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			cl.Append(tt.args.value)
			if cl.head.value != 8 {
				t.Errorf("CircularLinkedList.Append() = %v, want %v", cl.head.value, 8)
			}
			if cl.head.next.value != 22 {
				t.Errorf("CircularLinkedList.Append() = %v, want %v", cl.head.next.value, 22)
			}
		})
	}
}
