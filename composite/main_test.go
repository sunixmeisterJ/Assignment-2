package main

import "testing"

func TestTxtDocSize(t *testing.T) {
	if got := (TxtDoc{size: 50}).Size(); got != 50 {
		t.Fatalf("Size() = %d, want 50", got)
	}
}

func TestEmptyFolderSize(t *testing.T) {
	if got := (&Folder{}).Size(); got != 0 {
		t.Fatalf("Size() = %d, want 0", got)
	}
}

func TestNestedFolderSize(t *testing.T) {
	root, child, grandchild := &Folder{}, &Folder{}, &Folder{}
	root.AddDoc(TxtDoc{size: 10})
	root.AddDoc(TxtDoc{size: 50})
	grandchild.AddDoc(TxtDoc{size: 100})
	child.AddDoc(grandchild)
	child.AddDoc(&Folder{})
	root.AddDoc(child)
	if got := root.Size(); got != 160 {
		t.Fatalf("Size() = %d, want 160", got)
	}
	// Änderungen im Unterordner müssen bei erneuter Berechnung sichtbar sein.
	child.AddDoc(TxtDoc{size: 20})
	if got := root.Size(); got != 180 {
		t.Fatalf("Size() after AddDoc = %d, want 180", got)
	}
}

func Example_main() {
	main()
	// Output: 160
}
