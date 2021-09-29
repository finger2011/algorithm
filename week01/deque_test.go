package week01

import (
	"testing"
)

func TestMyCircularDeque_InsertFront(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			args: args{
				value: 5,
			},
			want: true,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 3,
			},
			args: args{
				value: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.InsertFront(tt.args.value); got != tt.want {
				t.Errorf("MyCircularDeque.InsertFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_InsertLast(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			args: args{
				value: 5,
			},
			want: true,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{3},
				capacity: 1,
			},
			args: args{
				value: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.InsertLast(tt.args.value); got != tt.want {
				t.Errorf("MyCircularDeque.InsertLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_DeleteFront(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 10,
			},
			want: true,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.DeleteFront(); got != tt.want {
				t.Errorf("MyCircularDeque.DeleteFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_DeleteLast(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 10,
			},
			want: true,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.DeleteLast(); got != tt.want {
				t.Errorf("MyCircularDeque.DeleteLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_GetFront(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 10,
			},
			want: 1,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.GetFront(); got != tt.want {
				t.Errorf("MyCircularDeque.GetFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_GetRear(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 10,
			},
			want: 3,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.GetRear(); got != tt.want {
				t.Errorf("MyCircularDeque.GetRear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_IsEmpty(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 10,
			},
			want: false,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.IsEmpty(); got != tt.want {
				t.Errorf("MyCircularDeque.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_IsFull(t *testing.T) {
	type fields struct {
		myQueue  []int
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				myQueue:  []int{1, 2, 3},
				capacity: 3,
			},
			want: true,
		},
		{
			name: "test02",
			fields: fields{
				myQueue:  []int{},
				capacity: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deque := &MyCircularDeque{
				myQueue:  tt.fields.myQueue,
				capacity: tt.fields.capacity,
			}
			if got := deque.IsFull(); got != tt.want {
				t.Errorf("MyCircularDeque.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
