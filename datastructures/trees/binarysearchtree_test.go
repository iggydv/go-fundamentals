package trees

import (
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
				Val:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			n.Insert(tt.args.val)
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
			name: "Search for a value",
			fields: fields{
				Key: 100,
				Left: &Node{
					Val: 50,
				},
				Right: &Node{
					Val: 150,
				},
			},
			args: args{
				val: 50,
			},
			want: true,
		},
		{
			name: "Search for a value that doesn't exist",
			fields: fields{
				Key: 100,
				Left: &Node{
					Val: 50,
				},
				Right: &Node{
					Val: 150,
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
				Val:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.Search(tt.args.val); got != tt.want {
				t.Errorf("Search(%d) = %v, want %v", tt.args.val, got, tt.want)
			}
		})
	}
}
