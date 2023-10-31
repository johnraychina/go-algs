package sub_str

// Brute-force algorithm needs backup for every mismatch.
//DFA is abstract string-searching machine.
//	・Finite number of states (including start and halt).
//	・Exactly one transition for each char in alphabet.
//	・Accept if sequence of transitions leads to halt state.

// Key differences from brute-force implementation.
// ・Need to precompute dfa[][] from pattern.
// ・Text pointer i never decrements.

// Include one state for each character in pattern (plus accept state).

// Mismatch transition.
// If in state j and next char c != pat.charAt(j),
// then the last j-1 characters of input are pat[1..j-1], followed by c.

// To compute dfa[c][j]: Simulate pat[1..j-1] on DFA and take transition c.
// Running time. Seems to require j steps.

// Mismatch transition.
// If in state j and next char c != pat.charAt(j),
// then the last j-1 characters of input are pat[1..j-1], followed by c. state X
// To compute dfa[c][j]: Simulate pat[1..j-1] on DFA and take transition c.
// Running time. Takes only constant time if we maintain state X.

// 对于重复的模式，如果到最后匹配不上，每次都要回拨指针.
// text:     AAAAAAABCCCCCDDDDEFFFGGG
//                i
// pattern:  AAAAAB
//                j
// 因为DFA有限状态自动机可以看成一个匹配状态机，KMP算法使用DFA记录匹配状态，解决回拨重复匹配的开销.
// 这种方法需要提前构建dfa[][]状态转移数组.

// KMPSearch 从文本txt中匹配模式pat, 给出txt的匹配位置.
func (d *DFA) KMPSearch(txt string) int {
	N := len(txt)
	M := d.states
	i := 0
	j := 0
	for ; i < N && j < M; i++ {
		c := txt[i]
		j = d.dfa[c][j]
	}

	if j == M {
		return i - M
	} else {
		return N
	}
}

// DFA 如何构造？
// 每个模式字符一个状态state，加上一个accept state.
// dfa[radix][state]
type DFA struct {
	dfa    [][]int
	radix  int
	states int
}

// 构造DFA关键1：使用pattern字符串构建DFA，匹配到accept节点则返回匹配成功。
// 考虑text与pattern不匹配时，回拨指针backup i,j
// 在状态j时，遇到不匹配的字符( c != pat.charAt(j) ), text最后的j-1个字符为pat[1...j-1]，例如
// text:     AAAAAAABCCCCCDDDDEFFFGGG
//            ^   i
// pattern:  AAAAAB
//            1   j
// 回拨指针i时，不会回到原点，而是原点+1位置， 而text[原点+1...i] 和 pat[1...j-1]是匹配的。
// 在DFA上再跑一遍pat[1...j-1]字符匹配，加上一个字符 c, 看起来要j个步骤.
// 直觉上看，很多情况下是没必要的，

// 构造DFA关键2：如果我们记住[1...j-1]的匹配状态X？遇到不匹配的情况，直接从状态X处继续匹配c就行了。
// 这是平时多流汗（维护状态X），战时少流血（mismatch时)的思路。

const Radix = 256

func BuildDFA(pat string) *DFA {
	// dfa[radix][states]
	states := len(pat)
	dfa := make([][]int, Radix)
	for i := range dfa {
		dfa[i] = make([]int, states)
	}

	X := 0
	j := 1
	dfa[pat[0]][0] = 1
	//For each state j:
	//・Copy dfa[][X] to dfa[][j] for mismatch case.
	//・Set dfa[pat.charAt(j)][j] to j+1
	//for match case. ・Update X.
	for ; j < states; j++ {
		for c := 0; c < Radix; c++ {
			dfa[c][j] = dfa[c][X]
		}
		dfa[pat[j]][j] = j + 1
		X = dfa[pat[j]][X]
	}

	return &DFA{
		dfa:    dfa,
		radix:  Radix,
		states: states,
	}
}

//Proposition. KMP substring search accesses no more than M + N chars to search for a pattern of length M in a text of length N.
//Pf. Each pattern char accessed once when constructing the DFA;
//each text char accessed once (in the worst case) when simulating the DFA.

//Proposition. KMP constructs dfa[][] in time and space proportional to R M.
//Larger alphabets. Improved version of KMP constructs nfa[] in time and
//space proportional to M.

// 命题：KMP字符串搜索，最多访问M+N这个字符
// 证明：对于模式字符串，只在构建DFA时访问一次，对于目标文本，每个字符用DFA最多匹配一次。

// 命题：KMP构造DFA的空间和时间与 R*M成正比.
