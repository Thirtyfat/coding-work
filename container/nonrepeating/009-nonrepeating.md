###  寻找最长不含有重复字符的子串

* 对于每一个字母x
    * lastOccurred[x]不存在，或者< start 那么则无需操作
    * lastOccurred >= start 更新start
    * 更新lastOccurred[x] ,更新maxLength
    
