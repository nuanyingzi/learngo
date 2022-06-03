package split

import (
	"reflect"
	"testing"
)

// 测试

func TestSplit(t *testing.T) {
	got := Split("我爱你", "爱")
	want := []string{"我", "你"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want:%v got:%v\n", want, got)
	}
}
