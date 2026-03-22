package service

// RegisterFetchers sets up the fetcher dependencies for the application.
// In a real application, this might involve dependency injection or a service registry.
func RegisterFetchers() GroceryFetcher {
	return NewMockGroceryFetcher()
}
