package find

import (
	"bytes"
	"slices"
	"testing"
)

const testFilepath = "foo"

func TestFind(t *testing.T) {
	// t.Run("test all the content in directory", func(t *testing.T) {
	// 	got := ReadContent(testFilepath)
	// 	want := []string{
	// 		"foo/bar",
	// 		"foo/bar/baz",
	// 		"foo/bar/broken_sl",
	// 		"foo/bar/buzz",
	// 		"foo/bar/test.txt",
	// 		"foo/bar/baz/deep",
	// 		"foo/bar/baz/deep/directory"}

	// 	if !slices.Equal(got.allContent, want) {
	// 		t.Errorf("got %s want %s", got, want)
	// 	}
	// })
	t.Run("test the list of folders", func(t *testing.T) {
		dc, err := ReadContent(testFilepath)
		got := dc.Folders()
		want := []string{
			"foo/bar",
			"foo/bar/baz",
			"foo/bar/baz/deep",
			"foo/bar/baz/deep/directory",
		}
		if err != nil {
			t.Errorf("wasn't expected error but got one %v", err)
		}
		if !slices.Equal(got, want) {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test the list of symbol links", func(t *testing.T) {
		dc, err := ReadContent(testFilepath)
		if err != nil {
			t.Errorf("wasn't expected error but got one %v", err)
		}
		got := dc.SymbolLinks()
		want := map[string]string{
			"foo/bar/broken_sl": "[broken]",
			"foo/bar/buzz":      "foo/bar/baz",
		}
		// want := []string{
		// 	"foo/bar/broken_sl",
		// 	"foo/bar/buzz",
		// }
		if !CompareMaps(got, want) {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test the list of files without folders and symbol links", func(t *testing.T) {
		dc, err := ReadContent(testFilepath)
		if err != nil {
			t.Errorf("wasn't expected error but got one %v", err)
		}
		got := dc.Files()
		want := []string{
			"foo/bar/test.txt",
		}
		if !slices.Equal(got, want) {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test PrintFolders", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		dc, err := ReadContent(testFilepath)
		if err != nil {
			t.Errorf("wasn't expected error but got one %v", err)
		}
		dc.PrintFolders(buffer)

		got := buffer.String()
		want := `foo/bar
foo/bar/baz
foo/bar/baz/deep
foo/bar/baz/deep/directory
`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("test PrintFiles", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		ext := "txt"
		dc, err := ReadContent(testFilepath)
		if err != nil {
			t.Errorf("wasn't expected error but got one %v", err)
		}
		dc.PrintFiles(buffer, ext)

		got := buffer.String()
		want := "foo/bar/test.txt\n"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test PrintSL", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		dc, err := ReadContent(testFilepath)
		if err != nil {
			t.Errorf("wasn't expected error but got one %v", err)
		}
		dc.PrintSymbolLinks(buffer)

		got := buffer.String()
		want1 := `foo/bar/broken_sl -> [broken]
foo/bar/buzz -> foo/bar/baz
`
		want2 := `foo/bar/buzz -> foo/bar/baz
foo/bar/broken_sl -> [broken]
`

		if !(got == want1 || got == want2) {
			t.Errorf("got %s want %s", got, want1)
		}
	})

	t.Run("test incorrect dir", func(t *testing.T) {

		_, got := ReadContent("non/exist/dir")
		want := "failed to read directory non/exist/dir: open non/exist/dir: no such file or directory"
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}

func CompareMaps(map1, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value := range map1 {
		if v, ok := map2[key]; !ok || v != value {
			return false
		}
	}

	return true
}
