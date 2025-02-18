package linkedlists

import "testing"

func TestLinkedList_Prepend(t *testing.T) {
	type fields struct {
		head   *Node
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
				head: &Node{
					value: 8,
					next:  nil,
				},
				length: 1,
			},
			args: args{
				value: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			l.Prepend(tt.args.value)
			if l.head.value != 22 {
				t.Errorf("LinkedList.Prepend() = %v, want %v", l.head.value, 22)
			}
		})

	}
}

func TestLinkedList_Append(t *testing.T) {
	type fields struct {
		head   *Node
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
				head: &Node{
					value: 8,
					next:  nil,
				},
				length: 1,
			},
			args: args{
				value: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			l.Append(tt.args.value)
			current := l.head
			for current.next != nil {
				current = current.next
			}
			if current.value != 22 {
				t.Errorf("LinkedList.Prepend() = %v, want %v", current.value, 22)
			}
		})
	}
}
