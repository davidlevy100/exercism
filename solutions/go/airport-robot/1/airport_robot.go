// Package airportrobot provides a simple example of using Go interfaces
// to support greetings in multiple languages.
package airportrobot

import "fmt"

// Greeter defines behavior for greeting visitors in a specific language.
type Greeter interface {
	// LanguageName returns the name of the language.
	LanguageName() string
	// Greet returns a greeting for the given visitor name.
	Greet(name string) string
}

// SayHello produces a greeting message using the provided Greeter.
func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}

// German greets visitors in German.
type German struct{}

// LanguageName returns "German".
func (German) LanguageName() string { return "German" }

// Greet returns a German greeting for the visitor.
func (German) Greet(name string) string { return fmt.Sprintf("Hallo %s!", name) }

// Italian greets visitors in Italian.
type Italian struct{}

// LanguageName returns "Italian".
func (Italian) LanguageName() string { return "Italian" }

// Greet returns an Italian greeting for the visitor.
func (Italian) Greet(name string) string { return fmt.Sprintf("Ciao %s!", name) }

// Portuguese greets visitors in Portuguese.
type Portuguese struct{}

// LanguageName returns "Portuguese".
func (Portuguese) LanguageName() string { return "Portuguese" }

// Greet returns a Portuguese greeting for the visitor.
func (Portuguese) Greet(name string) string { return fmt.Sprintf("Olá %s!", name) }
