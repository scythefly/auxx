/*
请你设计一个迭代器类，包括以下内容：

一个构造函数，输入参数包括：一个 有序且字符唯一 的字符串 characters
（该字符串只包含小写英文字母）和一个数字 combinationLength 。
函数 next() ，按 字典序 返回长度为 combinationLength 的下一个字母组合。
函数 hasNext() ，只有存在长度为 combinationLength 的下一个字母组合时，才返回 True；
否则，返回 False。
 

示例：

CombinationIterator iterator = new CombinationIterator("abc", 2); // 创建迭代器 iterator

iterator.next(); // 返回 "ab"
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 "ac"
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 "bc"
iterator.hasNext(); // 返回 false
 

提示：

1 <= combinationLength <= characters.length <= 15
每组测试数据最多包含 10^4 次函数调用。
题目保证每次调用函数 next 时都存在下一个字母组合。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/iterator-for-combination
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

type CombinationIterator struct {
	idx    int
	length int
	clen   int
	ct     string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	return CombinationIterator{
		idx:    0,
		length: len(characters),
		clen:   combinationLength,
		ct:     characters,
	}
}

func (c *CombinationIterator) Next() string {

}

func (c *CombinationIterator) HasNext() bool {

}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func printCombinationIterator() {
}

func TestCombinationIterator() {
}
