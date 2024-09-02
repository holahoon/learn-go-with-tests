package main

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

// -- Type 1 --
// func TestSearch(t *testing.T) {
// 	// dictionary := map[string]string{"test": "this is just a test"}
// 	dictionary := Dictionary{"test": "this is just a test"}

// 	got := dictionary.Search("test")
// 	want := "this is just a test"

// 	assertStrings(t, got, want)
// }

// -- Type 3 --
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
}

// -- Type 1 --
// func TestAdd(t *testing.T) {
// 	dictionary := Dictionary{}
// 	dictionary.Add("test", "this is just a test")

// 	want := "this is just a test"
// 	got, err := dictionary.Search("test")
// 	if err != nil {
// 		t.Fatal("should find the added word:", err)
// 	}

// 	assertStrings(t, got, want)
// }

// -- Type 2 --
// func TestAdd(t *testing.T) {
// 	dictionary := Dictionary{}
// 	word := "test"
// 	definition := "this is just a test"

// 	dictionary.Add(word, definition)

// 	assetDefinition(t, dictionary, word, definition)
// }

func assetDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assetDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictioanry := Dictionary{word: definition}
		err := dictioanry.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assetDefinition(t, dictioanry, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assetDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}
