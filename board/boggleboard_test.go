package board

import (
	"fmt"
	"testing"
)

func TestBoggle1(t *testing.T) {

	board := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	words := []string{"SEE", "ABCCED", "ABCB"}

	resolver := NewBoggleSolver(words, board)
	results := resolver.Solve()
	fmt.Println(results)
	// Output: [SEE ABCCED]
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
	for _, r := range results {
		if r != "SEE" && r != "ABCCED" {
			t.Errorf("expected SEE or ABCCED, got %s", r)
		}
	}
}

func TestBoggle2(t *testing.T) {

	board := [][]rune{
		{'G', 'I', 'Z'},
		{'U', 'E', 'K'},
		{'Q', 'S', 'E'},
	}
	words := []string{"GEEKS", "FOR", "QUIZ", "GO"}

	resolver := NewBoggleSolver(words, board)
	results := resolver.Solve()
	fmt.Println(results)
	// Output: [GEEKS QUIZ]
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
	for _, r := range results {
		if r != "GEEKS" && r != "QUIZ" {
			t.Errorf("expected GEEKS or QUIZ, got %s", r)
		}
	}
}
