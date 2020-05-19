###rune
> rune相当于go的char
* 使用range遍历pos，rune对，但是对于中文支持并不太友好，需要进行编解码
* 使用utf8.RuneCountInString 获得字符数量
* 使用len获得字节长度
* 使用[]byte获得字节
