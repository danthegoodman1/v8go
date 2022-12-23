package v8go

import (
	"fmt"
	"testing"
)

func TestIsolationConfig(t *testing.T) {
	iso := NewIsolate()
	defer iso.Dispose()
	ctx := NewContext(iso)
	defer ctx.Close()

	_, err := ctx.RunScript(`
	let a = []
	for (let i =0; i < 10000000; i++) {
		a.push('hello')
	}
`, "yeye.js")
	if err != nil {
		fmt.Println("ERR")
		t.Error(err)
	}

	hs := iso.GetHeapStatistics()
	fmt.Printf("%+v \n", hs)

	if hs.NumberOfNativeContexts != 2 {
		t.Error("expect NumberOfNativeContexts return 3, got", hs.NumberOfNativeContexts)
	}

	if hs.NumberOfDetachedContexts != 0 {
		t.Error("expect NumberOfDetachedContexts return 0, got", hs.NumberOfDetachedContexts)
	}
}

// 24772608 - 1M
// 37359616 - 2M
// 47140864 - 3M

// 10MB - 22544384
// 20MB - 45088768
// 30MB - 80216064
// 1000MB - 200540160
