package main

import "testing"

func TestSLinkedList_isEmpty(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Populated list", fields{head: &SLLNode{val: 5, next: nil}, tail: &SLLNode{val: 15, next: nil}}, false},
		{"Empty list", fields{head: nil, tail: nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := list.isEmpty(); got != tt.want {
				t.Errorf("SLinkedList.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSLinkedList_addToHead(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	type args struct {
		el int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.addToHead(tt.args.el)
		})
	}
}

func TestSLinkedList_addToTail(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	type args struct {
		el int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.addToTail(tt.args.el)
		})
	}
}

func TestSLinkedList_deleteFromHead(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := list.deleteFromHead(); got != tt.want {
				t.Errorf("SLinkedList.deleteFromHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSLinkedList_deleteFromTail(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := list.deleteFromTail(); got != tt.want {
				t.Errorf("SLinkedList.deleteFromTail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSLinkedList_printAll(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.printAll()
		})
	}
}

func TestSLinkedList_isInList(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	type args struct {
		el int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := list.isInList(tt.args.el); got != tt.want {
				t.Errorf("SLinkedList.isInList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSLinkedList_delete(t *testing.T) {
	type fields struct {
		head *SLLNode
		tail *SLLNode
	}
	type args struct {
		el int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			list.delete(tt.args.el)
		})
	}
}
