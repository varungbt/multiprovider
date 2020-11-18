// Package muxutil provides helpers for providermux ...
package muxutil

// SchemeMap holds the cloud provdier id and the reference to the concrete type ...
type SchemeMap struct {
	provider string
	m        map[string]interface{}
}

// Register a concrete type
func (m *SchemeMap) Register(provider string, value interface{}) {
	if m.m == nil {
		m.m = map[string]interface{}{}
	}

	m.m[provider] = value

}
