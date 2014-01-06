package gologgly

import ()

type inputs map[string]string

var (
	loggly = new()
)

// Initialize map for storing the application's inputs.
func new() (i inputs) {
	return make(inputs)
}

// Exported function for adding inputs to the loggy inputs map.
func AddInput(name string, customer_token string) (err error) {
	loggly.add(name, customer_token)
	return nil
}

// Add a new input to the input map.
func (i inputs) add(name string, customer_token string) (points inputs) {

	// Store in the map.
	i[name] = customer_token

	return i
}
