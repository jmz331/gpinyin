HELLO WORLD

实现一个golang版本的汉字转拼音类

TODO：
》实现汉字转拼音部分功能

已完成：
》简繁体转换
》带音标与不带音标的转换
》单元测试

使用方式：
go get github.com/jmz331/gpinyin
```
import "github.com/jmz331/gpinyin"

const s = "台我要1234!#$翻译成繁体的汉字堡垒asdf"
r1 := ConvertToPinyinString(s, "-", PINYIN_WITHOUT_TONE)
//tai-wo-yao-1234!#$-fan-yi-cheng-fan-ti-de-han-zi-bao-lei-asdf
r2 := ConvertToPinyinString(s, "-", PINYIN_WITH_TONE_MARK)
//tái-wǒ-yào-1234!#$-fān-yì-chéng-fán-tǐ-de-hàn-zì-bǎo-lěi-asdf
```

###参考
https://github.com/stuxuhai/jpinyin
