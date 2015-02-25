package gpinyin

import (
	"testing"
	// "unicode/utf8"
)

func equal(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("转换繁体失败,%q == %q，期待值：%q", got, want, want)
	}
}

func TestT2SAndS2TConvert(t *testing.T) {
	const s1 = "台我要1234!#$翻译成繁体的汉字asdf"
	const s2 = "臺我要1234!#$翻譯成繁體的漢字asdf"

	r1 := ConvertToTraditionalChinese(s1)
	r2 := ConvertToSimplifiedChinese(s2)

	equal(t, r1, s2)
	equal(t, r2, s1)
}

func TestHello(t *testing.T) {
	t.Logf("r:%v", isChinese("我"))
	t.Logf("r:%v", isChinese("臺"))
	t.Logf("r:%v", isChinese("A"))
}
