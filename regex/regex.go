package regex

//Following Thompson's paper, the compiler builds an NFA from a regular expression in postfix notation with dot (.) added as an explicit concatenation operator.
//A separate function re2post rewrites infix regular expressions like “a(bb)+a” into equivalent postfix expressions like “abb.+.a.”.
//(A “real” implementation would certainly need to use dot as the “any character” metacharacter rather than as a concatenation operator.
//A real implementation would also probably build the NFA during parsing rather than build an explicit postfix expression.
//However, the postfix version is convenient and follows Thompson's paper more closely.)

// As the compiler scans the postfix expression, it maintains a stack of computed NFA fragments.
// Literals push new NFA fragments onto the stack, while operators pop fragments off the stack and then push a new fragment.
// For example, after compiling the abb in abb.+.a., the stack contains NFA fragments for a, b, and b.
// The compilation of the . that follows pops the two b NFA fragment from the stack and pushes an NFA fragment for the concatenation bb..
// Each NFA fragment is defined by its start state and its outgoing arrows:

type State struct {
	c        int
	out      *State
	out1     *State
	lastList int //LastList is used during execution
}

type Frag struct {
	start *State
	out   *Ptrlist
}

type Ptrlist struct {
	list []*State
}

// List1 creates a new pointer list containing the single pointer outp.
func List1(outp *State) *Ptrlist {
	return &Ptrlist{list: []*State{outp}}
}

// Append concatenates two pointer lists, returning the result
func Append(l1 *Ptrlist, l2 *Ptrlist) *Ptrlist {
	l1.list = append(l1.list, l2.list...)
	return l1
}

// Patch connects the dangling arrows in the pointer list l to the state s:
// it sets *outp = s for each pointer outp in l.
func Patch(l *Ptrlist, s *State) {
	for _, e := range l.list {
		e.out = s
	}
}

func post2nfa(postfix rune) *State {
	panic("todo")
}

//State* post2nfa(char *postfix) {
//char *p;
//Frag stack[1000], *stackp, e1, e2, e;
//State *s;
//
//#define push(s) *stackp++ = s
//#define pop()   *--stackp
//
//stackp = stack;
//for(p=postfix; *p; p++){
//switch(*p){
///* compilation cases, described below */
//}
//}
//
//e = pop();
//patch(e.out, matchstate);
//return e.start;
//}
