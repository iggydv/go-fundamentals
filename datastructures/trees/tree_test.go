package trees

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTreeNode_Search(t *testing.T) {
	type fields struct {
		Key      int
		Children []*TreeNode
	}
	type args struct {
		val int
	}

	testTree := &TreeNode{
		Key: 1,
		Children: []*TreeNode{
			{
				Key: 2,
				Children: []*TreeNode{
					{
						Key: 3,
					},
				},
			},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "DFSPreOrder for a value in a tree",
			fields: fields{
				Key:      1,
				Children: []*TreeNode{testTree},
			},
			args: args{
				val: 3,
			},
			want: true,
		},
		{
			name: "DFSPreOrder for a value in a tree",
			fields: fields{
				Key:      1,
				Children: []*TreeNode{testTree},
			},
			args: args{
				val: 7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &TreeNode{
				Key:      tt.fields.Key,
				Children: tt.fields.Children,
			}
			if got := n.DFSPreOrder(tt.args.val); got != tt.want {
				t.Errorf("DFSPreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_DFSPostOrder(t *testing.T) {
	type fields struct {
		Key      int
		Children []*TreeNode
	}
	type args struct {
		val int
	}
	testTree := &TreeNode{
		Key: 2,
		Children: []*TreeNode{
			{
				Key: 3,
				Children: []*TreeNode{
					{
						Key: 4,
					},
				},
			},
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "DFSPreOrder for a value in a tree",
			fields: fields{
				Key:      1,
				Children: []*TreeNode{testTree},
			},
			args: args{
				val: 3,
			},
			want: true,
		},
		{
			name: "DFSPreOrder for a value in a tree",
			fields: fields{
				Key:      1,
				Children: []*TreeNode{testTree},
			},
			args: args{
				val: 7,
			},
			want: false,
		},
		{
			name: "DFSPreOrder for a value in a tree",
			fields: fields{
				Key:      1,
				Children: []*TreeNode{testTree},
			},
			args: args{
				val: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &TreeNode{
				Key:      tt.fields.Key,
				Children: tt.fields.Children,
			}
			if got := n.DFSPostOrder(tt.args.val); got != tt.want {
				t.Errorf("DFSPreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_DFSPostOrderSlice(t *testing.T) {
	type fields struct {
		Key      int
		Children []*TreeNode
	}
	testTree := &TreeNode{
		Key: 2,
		Children: []*TreeNode{
			{
				Key: 3,
				Children: []*TreeNode{
					{
						Key: 4,
					},
				},
			},
		},
	}
	fmt.Println(testTree)
	tt := &TreeNode{
		Key:      1,
		Children: []*TreeNode{testTree},
	}
	fmt.Println(tt)

	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "DFSPreOrder for a value in a tree",
			fields: fields{
				Key:      1,
				Children: []*TreeNode{testTree},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &TreeNode{
				Key:      tt.fields.Key,
				Children: tt.fields.Children,
			}
			if got := n.DFSPostOrderSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFSPostOrderSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
