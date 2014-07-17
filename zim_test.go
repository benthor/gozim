package zim

import "testing"

func TestOpen(t *testing.T) {
	z, err := NewReader("test.zim")
	if err != nil {
		t.Errorf("Can't read %v", err)
	}

	if z.ArticleCount == 0 {
		t.Errorf("No article found")
	}
}