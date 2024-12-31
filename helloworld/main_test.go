package main

import "testing"

func TestHello(t *testing.T) {

	cases := []struct {
		description string
		name        string
		language    string
		want        string
	}{
		{description: "say hello to a specific person", name: "Chris", language: "", want: "Hello, Chris"},
		{description: "default to 'Hello, World' when name is empty", name: "", language: "", want: "Hello, World"},
		{description: "say hello in Spanish", name: "Elodia", language: "ES", want: "Hola, Elodia"},
		{description: "say hello in French", name: "Jean", language: "FR", want: "Bonjour, Jean"},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			got := Hello(tc.name, tc.language)
			want := tc.want
			assertCorrectMessage(t, got, want)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() //puts line number of function call upon fail in report
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
