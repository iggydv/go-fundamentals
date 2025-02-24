package trees

import (
	"testing"
)

func TestBinaryTreeNode_BFS(t *testing.T) {
	type fields struct {
		Key   int
		Left  *BinaryTreeNode
		Right *BinaryTreeNode
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
			name: "BFS",
			fields: fields{
				Key: 10,
				Left: &BinaryTreeNode{
					Key: 5,
					Left: &BinaryTreeNode{
						Key: 3,
						Left: &BinaryTreeNode{
							Key: 1,
						},
					},
				},
			},
			args: args{
				val: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryTreeNode{
				Key:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.BFS(tt.args.val); got != tt.want {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTreeNode_DFSPostOrderSearch(t *testing.T) {
	type fields struct {
		Key   int
		Left  *BinaryTreeNode
		Right *BinaryTreeNode
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
			name: "DFSPostOrderSearch",
			fields: fields{
				Key: 10,
				Left: &BinaryTreeNode{
					Key: 5,
					Left: &BinaryTreeNode{
						Key: 3,
						Left: &BinaryTreeNode{
							Key: 1,
						},
					},
				},
			},
			args: args{
				val: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryTreeNode{
				Key:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.DFSPostOrderSearch(tt.args.val); got != tt.want {
				t.Errorf("DFSPostOrderSearch(%d) = %v, want %v", tt.args.val, got, tt.want)
			}
		})
	}
}

func TestBinaryTreeNode_DFSPreOrderSearch(t *testing.T) {
	type fields struct {
		Key   int
		Left  *BinaryTreeNode
		Right *BinaryTreeNode
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
			name: "DFSPostOrderSearch",
			fields: fields{
				Key: 10,
				Left: &BinaryTreeNode{
					Key: 5,
					Left: &BinaryTreeNode{
						Key: 3,
						Left: &BinaryTreeNode{
							Key: 1,
						},
					},
				},
			},
			args: args{
				val: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryTreeNode{
				Key:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.DFSPreOrderSearch(tt.args.val); got != tt.want {
				t.Errorf("DFSPreOrderSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTreeNode_InsertLevelOrder(t *testing.T) {
	type fields struct {
		Key   int
		Left  *BinaryTreeNode
		Right *BinaryTreeNode
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
			name: "InsertLevelOrder",
			fields: fields{
				Key: 10,
				Left: &BinaryTreeNode{
					Key: 5,
					Left: &BinaryTreeNode{
						Key: 3,
						Left: &BinaryTreeNode{
							Key: 1,
						},
					},
				},
			},
			args: args{
				val: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &BinaryTreeNode{
				Key:   tt.fields.Key,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			n.InsertLevelOrder(tt.args.val)

			if !n.BFS(tt.args.val) {
				t.Errorf("InsertLevelOrder() = %v, want %v", false, true)
			}
		})
	}
}
