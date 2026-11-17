package main

import "fmt"

// Doc ist die gemeinsame Schnittstelle für Blätter und Komposita.
type Doc interface {
	Size() int
}

// TxtDoc ist ein Blatt mit fester Größe.
type TxtDoc struct {
	size int
}

func (d TxtDoc) Size() int { return d.size }

// Folder kann Textdokumente und weitere Ordner enthalten.
type Folder struct {
	docs []Doc
}

func (f *Folder) AddDoc(doc Doc) {
	f.docs = append(f.docs, doc)
}

func (f *Folder) Size() int {
	total := 0
	for _, doc := range f.docs {
		total += doc.Size()
	}
	return total
}

func main() {
	root := &Folder{}
	root.AddDoc(TxtDoc{size: 10})
	root.AddDoc(TxtDoc{size: 50})
	subfolder := &Folder{}
	subfolder.AddDoc(TxtDoc{size: 100})
	root.AddDoc(subfolder)
	fmt.Println(root.Size())
}
