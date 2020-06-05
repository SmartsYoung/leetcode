package _648_repalce_words

import (
	"sort"
	"strings"
)

/*
648. 单词替换
在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。

你需要输出替换之后的句子。



示例：

输入：dict(词典) = ["cat", "bat", "rat"] sentence(句子) = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
*/

func replaceWords(dict []string, sentence string) string {
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) < len(dict[j])
	})

	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(dict); j++ {
			if strings.HasPrefix(words[i], dict[j]) {
				words[i] = dict[j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
