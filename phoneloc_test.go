package phoneloc

import (
	"math/rand"
	"testing"
	"time"
)

type Target struct {
	Phone int
	City  string
	Sp    string
}

func TestFind(t *testing.T) {
	var results = []Target{
		{1891508, "常州", "电信"},
		{1525198, "常州", "移动"},
	}
	parser, err := NewParser("phone.dat")
	if err != nil {
		panic(err)
	}
	t.Log("开始测试Find函数")
	errFlag := false
	for index, result := range results {
		loc, err := parser.Find(result.Phone)
		if err != nil {
			t.Logf("  查询失败：%s\n", err.Error())
			errFlag = true
			break
		}
		t.Logf("第[%d]组查询 [%v]\n", index+1, result.Phone)
		t.Logf("  |-预期结果：[%s] [%s]\n", result.City, result.Sp)
		t.Logf("  |-查询结果：[%s] [%s]\n", loc.City, loc.Sp)

		if loc.City != result.City || loc.Sp != result.Sp {
			errFlag = true
		}
	}
	if errFlag {
		t.Fatal("\x1b[31m测试失败！\x1b[0m")
	}
	t.Log("\x1b[32m测试通过！\x1b[0m")
}

func BenchmarkFind(b *testing.B) {
	b.StopTimer()

	loc, err := NewParser("phone.dat")
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = loc.Find(1300006)
	}
}
