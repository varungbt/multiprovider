// Package muxutil provides helpers for providermux ...
package muxutil

// ProviderMap holds the cloud provdier id and the reference to the concrete type ...
type ProviderMap struct {
	provider string
	m        map[string]interface{}
}

// Register a concrete type
func (m *ProviderMap) Register(provider string, value interface{}) {
	if m.m == nil {
		m.m = map[string]interface{}{}
	}

	m.m[provider] = value

}

// FromString returns the DbClientOpener for the provider
func (m *ProviderMap) FromString(provider string) interface{} {
	providerDbOpener, ok := m.m[provider]

	if !ok {
		panic("Provider specific DB client opener no found. Check the CLOUDPROVIDER env variable")
	}
	return providerDbOpener
}
