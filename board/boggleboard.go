package board

// TrieNode represents a node in the trie data structure
type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

// NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}
}

// Insert inserts a word into the trie
func (node *TrieNode) Insert(word string) {
	curr := node
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = NewTrieNode()
		}
		curr = curr.children[ch]
	}
	curr.isWord = true
}

type BoggleSolver struct {
	board   [][]rune
	trie    *TrieNode
	visited [][]bool
	result  map[string]bool
}

func NewBoggleSolver(dictionary []string, board [][]rune) *BoggleSolver {
	trie := NewTrieNode()
	for _, word := range dictionary {
		trie.Insert(word)
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	return &BoggleSolver{
		board:   board,
		trie:    trie,
		visited: visited,
		result:  make(map[string]bool),
	}
}

func (solver *BoggleSolver) Solve() []string {
	for i := range solver.board {
		for j := range solver.board[i] {
			solver.search(i, j, solver.trie, "")
		}
	}

	words := make([]string, 0, len(solver.result))
	for word := range solver.result {
		words = append(words, word)
	}

	return words
}

func (solver *BoggleSolver) search(row, col int, node *TrieNode, prefix string) {
	if row < 0 || col < 0 || row >= len(solver.board) || col >= len(solver.board[row]) {
		return
	}
	if solver.visited[row][col] {
		return
	}

	ch := solver.board[row][col]
	childNode, ok := node.children[ch]
	if !ok {
		return
	}

	solver.visited[row][col] = true
	defer func() { solver.visited[row][col] = false }()

	newPrefix := prefix + string(ch)
	if childNode.isWord {
		solver.result[newPrefix] = true
	}

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range dirs {
		solver.search(row+dir[0], col+dir[1], childNode, newPrefix)
	}
}
