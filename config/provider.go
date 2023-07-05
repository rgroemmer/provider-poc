/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	object_storage_project "github.com/rgroemmer/provider-poc/config/object-storage-project"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "poc"
	modulePath     = "github.com/rgroemmer/provider-poc"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		object_storage_project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
