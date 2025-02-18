package trees

import (
	"fmt"
	"testing"
)

func TestNode_Insert(t *testing.T) {
	type fields struct {
		Key   int
		Left  *Node
		Right *Node
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Insert a smaller value",
			fields: fields{
				Key: 100,
			},
			args: args{
				val: 50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Key:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			n.Insert(tt.args.val)
			fmt.Print(n)
		})
	}
}

func TestNode_Search(t *testing.T) {
	type fields struct {
		Key   int
		Left  *Node
		Right *Node
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "DFSPreOrder for a value",
			fields: fields{
				Key: 100,
				Left: &Node{
					Key: 50,
				},
				Right: &Node{
					Key: 150,
				},
			},
			args: args{
				val: 50,
			},
			want: true,
		},
		{
			name: "DFSPreOrder for a value that doesn't exist",
			fields: fields{
				Key: 100,
				Left: &Node{
					Key: 50,
				},
				Right: &Node{
					Key: 150,
				},
			},
			args: args{
				val: 350,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Key:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.Search(tt.args.val); got != tt.want {
				t.Errorf("DFSPreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
