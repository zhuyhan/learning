package sort

import (
	"reflect"
	"testing"
)

func Test_shellSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "希尔排序",
			args: args{
				nums: []int{1, 4, 5, 6, 2, 3, 11, 45, 30, 31, 12, 11, 23, 12},
			},
			want: []int{1, 2, 3, 4, 5, 6, 11, 11, 12, 12, 23, 30, 31, 45},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
