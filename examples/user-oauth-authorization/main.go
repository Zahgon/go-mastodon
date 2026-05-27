package main

func ConfigureClient() { _ = "STUB: not implemented"; return }

// Have the user manually get the token and send it back to us

// Create the client

// Exchange the User authentication code with an access token, that can be used to interact with the api on behalf of the user

// Lets Export the secrets so we can use them later to preform actions on behalf of the user
// Without having to request authroization all the time.
// Exporting this as Environment variables, but it can be a configuration file, or database, anywhere you'd like to keep this credentials

// Preform user actions wihtout having to re-authenticate again
func doUserActions() {
	_ = "STUB: not implemented"
	// Load Environment variables, config file, secrets from db
	return
}

// instanciate the new client

// Let's do some actions on behalf of the user!

// Post a toot

func main() {
	ConfigureClient()
	doUserActions()
}
