package count

import (
	"log"
	"testing"
)

const (
	testFilePath      = "test.txt"
	errorFileNotExist = "failed to read file file/not/exist: open file/not/exist: no such file or directory"
)

func TestCountLines(t *testing.T) {
	t.Run("test for line counter func without error", func(t *testing.T) {
		got, err := CountLines(testFilePath)
		want := 3
		if err != nil {
			t.Errorf("wasn't expected error but got one: %v", err)
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test for line counter func with error", func(t *testing.T) {
		got, err := CountLines("file/not/exist")
		want := 0
		if err.Error() != errorFileNotExist {
			t.Errorf("expected error but haven't got one. Got: %v", err.Error())
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestCountChars(t *testing.T) {
	t.Run("test count chars func", func(t *testing.T) {
		got, err := CountChars(testFilePath)
		want := 39
		if err != nil {
			t.Errorf("wasn't expected error but got one: %v", err)
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test for char counter func with error", func(t *testing.T) {
		got, err := CountChars("file/not/exist")
		want := 0
		if err.Error() != errorFileNotExist {
			t.Errorf("expected error but haven't got one. Got: %v", err.Error())
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestCountWords(t *testing.T) {
	t.Run("test CountWords func without error", func(t *testing.T) {
		got, err := CountWords(testFilePath)
		want := 8

		if err != nil {
			t.Errorf("wasn't expected error but got one: %v", err)
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test CountWords func with error", func(t *testing.T) {
		got, err := CountChars("file/not/exist")
		want := 0
		if err.Error() != errorFileNotExist {
			t.Errorf("expected error but haven't got one. Got: %v", err.Error())
		}
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestDefineFunc(t *testing.T) {
	t.Run("test DefineFunc with lines counting", func(t *testing.T) {

		linesFlag := true
		wordsFlag := false
		charsFlag := false

		gotFunc, err := DefineFunc(linesFlag, wordsFlag, charsFlag)
		wantFunc := CountLines

		got, err1 := gotFunc(testFilePath)
		want, err2 := wantFunc(testFilePath)
		if err != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err1 != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err2 != nil {
			t.Errorf("didn't expect error but got one %v", err2)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test DefineFunc with words counting", func(t *testing.T) {

		linesFlag := false
		wordsFlag := true
		charsFlag := false

		gotFunc, err := DefineFunc(linesFlag, wordsFlag, charsFlag)
		wantFunc := CountWords

		got, err1 := gotFunc(testFilePath)
		want, err2 := wantFunc(testFilePath)

		if err != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err1 != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err2 != nil {
			t.Errorf("didn't expect error but got one %v", err2)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test DefineFunc with chars counting", func(t *testing.T) {

		linesFlag := false
		wordsFlag := false
		charsFlag := true

		gotFunc, err := DefineFunc(linesFlag, wordsFlag, charsFlag)
		wantFunc := CountChars

		got, err1 := gotFunc(testFilePath)
		want, err2 := wantFunc(testFilePath)

		if err != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err1 != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err2 != nil {
			t.Errorf("didn't expect error but got one %v", err2)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("test DefineFunc with wrong flags", func(t *testing.T) {

		linesFlag := false
		wordsFlag := true
		charsFlag := true

		_, err := DefineFunc(linesFlag, wordsFlag, charsFlag)

		if err == nil {
			log.Fatal()
		}
		if err.Error() != "wrong flags or their combination specified" {
			t.Errorf("expect error but wrong one %v", err)
		}
	})

	t.Run("test DefineFunc with wrong flags 2", func(t *testing.T) {

		linesFlag := true
		wordsFlag := true
		charsFlag := true

		_, err := DefineFunc(linesFlag, wordsFlag, charsFlag)

		if err == nil {
			log.Fatal()
		}
		if err.Error() != "wrong flags or their combination specified" {
			t.Errorf("expect error but wrong one %v", err)
		}
	})

	t.Run("test DefineFunc with wrong flags 3", func(t *testing.T) {

		linesFlag := true
		wordsFlag := true
		charsFlag := false

		_, err := DefineFunc(linesFlag, wordsFlag, charsFlag)

		if err == nil {
			log.Fatal()
		}
		if err.Error() != "wrong flags or their combination specified" {
			t.Errorf("expect error but wrong one %v", err)
		}
	})

	t.Run("test DefineFunc with wrong flags 3", func(t *testing.T) {

		linesFlag := true
		wordsFlag := false
		charsFlag := true

		_, err := DefineFunc(linesFlag, wordsFlag, charsFlag)

		if err == nil {
			log.Fatal()
		}
		if err.Error() != "wrong flags or their combination specified" {
			t.Errorf("expect error but wrong one %v", err)
		}
	})

	t.Run("test DefineFunc with no specified flags", func(t *testing.T) {

		linesFlag := false
		wordsFlag := false
		charsFlag := false

		gotFunc, err := DefineFunc(linesFlag, wordsFlag, charsFlag)
		wantFunc := CountWords

		got, err1 := gotFunc(testFilePath)
		want, err2 := wantFunc(testFilePath)

		if err != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err1 != nil {
			t.Errorf("didn't expect error but got one %v", err1)
		}
		if err2 != nil {
			t.Errorf("didn't expect error but got one %v", err2)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
