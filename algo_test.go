package cvm

import (
	"strings"
	"testing"
)

var source = `Let’s return to Hamlet, but this time your working memory — consisting of a whiteboard — has room for just 100 words. Once the play starts, you write down the first 100 words you hear, again skipping any repeats. When the space is full, press pause and flip a coin for each word. Heads, and the word stays on the list; tails, and you delete it. After this preliminary round, you’ll have about 50 distinct words left.

Now you move forward with what the team calls Round 1. Keep going through Hamlet, adding new words as you go. If you come to a word that’s already on your list, flip a coin again. If it’s tails, delete the word; heads, and the word stays on the list. Proceed in this fashion until you have 100 words on the whiteboard. Then randomly delete about half again, based on the outcome of 100 coin tosses. That concludes Round 1.

Next, move to Round 2. Continue as in Round 1, only now we’ll make it harder to keep a word. When you come to a repeated word, flip the coin again. Tails, and you delete it, as before. But if it comes up heads, you’ll flip the coin a second time. Only keep the word if you get a second heads. Once you fill up the board, the round ends with another purge of about half the words, based on 100 coin tosses.

In the third round, you’ll need three heads in a row to keep a word. In the fourth round you’ll need four heads in a row. And so on.

RELATED:
Scientists Find Optimal Balance of Data Storage and Time
Researchers Approach New Speed Limit for Seminal Problem
Mathematicians Discover the Perfect Way to Multiply
Eventually, in the kth round, you’ll reach the end of Hamlet. The point of the exercise has been to ensure that every word, by virtue of the random selections you’ve made, has the same probability of being there: 1/2k. If, for instance, you have 61 words on your list at the conclusion of Hamlet, and the process took six rounds, you can divide 61 by the probability, 1/26, to estimate the number of distinct words — which comes out to 3,904 in this case. (It’s easy to see how this procedure works: Suppose you start with 100 coins and flip each one individually, keeping only those that come up heads. You’ll end up with close to 50 coins, and if someone divides that number by the probability, ½, they can guess that there were about 100 coins originally.)

Variyam and Meel mathematically proved that the accuracy of this technique scales with the size of the memory. Hamlet has exactly 3,967 unique words. (They counted.) In experiments using a memory of 100 words, the average estimate after five runs was 3,955 words. With a memory of 1,000 words, the average improved to 3,964. “Of course,” Variyam said, “if the [memory] is so big that it fits all the words, then we can get 100% accuracy.”

“This is a great example of how, even for very basic and well-studied problems, there are sometimes very simple but non-obvious solutions still waiting to be discovered,” said William Kuszmaul of Harvard University.
`

// This test should fail because it's memory is significantly lower than the unique count.
func TestSmallStack(t *testing.T) {

	vx := strings.ReplaceAll(source, ".", "")
	ax := strings.Split(vx, " ")

	og := make(map[string]struct{})
	for _, word := range ax {
		if _, ok := og[word]; !ok {
			og[word] = struct{}{}
		}
	}

	uniqueCount := float64(len(og))

	// At this scale, with such small memory, inaccuracies are expected.

	if n := Count(ax, 10); n == uniqueCount {
		t.Errorf("wanted %f, got %f", uniqueCount, n)
	}

}

// This test should work because it's memory is significantly higher than the unique count.
// TODO: I need to tweak the algorithm so that this works.
func TestFullStack(t *testing.T) {
	vx := strings.ReplaceAll(source, ".", "")
	ax := strings.Split(vx, " ")

	og := make(map[string]struct{})
	for _, word := range ax {
		if _, ok := og[word]; !ok {
			og[word] = struct{}{}
		}
	}

	uniqueCount := float64(len(og))

	// At this scale, with such small memory, inaccuracies are expected.

	if n := Count(ax, 10_000); n != uniqueCount {
		t.Errorf("wanted %f, got %f", uniqueCount, n)
	}
}


