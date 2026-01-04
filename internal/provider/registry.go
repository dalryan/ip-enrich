package provider

import (
	"fmt"
	"sort"
	"sync"
)

// Registry manages the collection of available providers.
type Registry struct {
	mu        sync.RWMutex
	providers map[string]Provider
}

// Global default registry
var defaultRegistry = NewRegistry()

// NewRegistry creates a new empty provider registry.
func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string]Provider),
	}
}

// Register adds a provider to the registry.
// Panics if a provider with the same ID is already registered.
func (r *Registry) Register(p Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := p.ID()
	if _, exists := r.providers[id]; exists {
		panic(fmt.Sprintf("provider already registered: %s", id))
	}
	r.providers[id] = p
}

// Get retrieves a provider by ID.
func (r *Registry) Get(id string) (Provider, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.providers[id]
	return p, ok
}

// All returns all registered providers, sorted by name.
func (r *Registry) All() []Provider {
	r.mu.RLock()
	defer r.mu.RUnlock()

	providers := make([]Provider, 0, len(r.providers))
	for _, p := range r.providers {
		providers = append(providers, p)
	}

	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Name() < providers[j].Name()
	})

	return providers
}

// IDs returns all registered provider IDs, sorted alphabetically.
func (r *Registry) IDs() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ids := make([]string, 0, len(r.providers))
	for id := range r.providers {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

// Filter returns providers matching the given IDs.
// If ids is empty, returns all providers.
func (r *Registry) Filter(ids []string) []Provider {
	if len(ids) == 0 {
		return r.All()
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	seen := make(map[string]struct{})
	providers := make([]Provider, 0, len(ids))

	for _, id := range ids {
		if _, alreadySeen := seen[id]; alreadySeen {
			continue
		}
		seen[id] = struct{}{}

		if p, ok := r.providers[id]; ok {
			providers = append(providers, p)
		}
	}

	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Name() < providers[j].Name()
	})

	return providers
}

// Validate validates a list of providers by their IDs
func (r *Registry) Validate(ids []string) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var unknown []string
	for _, id := range ids {
		if _, exists := r.providers[id]; !exists {
			unknown = append(unknown, id)
		}
	}
	return unknown
}

// Register adds a provider to the default registry.
func Register(p Provider) {
	defaultRegistry.Register(p)
}

// Get retrieves a provider from the default registry.
func Get(id string) (Provider, bool) {
	return defaultRegistry.Get(id)
}

// All returns all providers from the default registry.
func All() []Provider {
	return defaultRegistry.All()
}

// IDs returns all provider IDs from the default registry.
func IDs() []string {
	return defaultRegistry.IDs()
}

// Filter returns filtered providers from the default registry.
func Filter(ids []string) []Provider {
	return defaultRegistry.Filter(ids)
}

// Validate validates each provider in the default registry.
func Validate(ids []string) []string {
	return defaultRegistry.Validate(ids)
}
