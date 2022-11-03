package modules

import (
	"context"
	"errors"

	"github.com/416-C/terraform-registry-go/modules/apiv1"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type Modules struct {
	client *resty.Client
}

func NewModulesService(client *resty.Client) *Modules {
	return &Modules{
		client: client,
	}
}

// Search Modules
// https://developer.hashicorp.com/terraform/registry/api-docs#search-modules
func (s *Modules) Search(ctx context.Context, q *apiv1.SearchParams) (*apiv1.ListModules, error) {
	values, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	response, err := s.client.NewRequest().SetContext(ctx).SetResult(&apiv1.ListModules{}).SetQueryParamsFromValues(values).Get("search")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.ListModules), nil
}

// Get a Specific Module
// https://developer.hashicorp.com/terraform/registry/api-docs#get-a-specific-module
func (s *Modules) Get(ctx context.Context, namespace, name, provider, version string) (*apiv1.Module, error) {
	pathParams := map[string]string{
		"namespace": namespace,
		"name":      name,
		"provider":  provider,
		"version":   version,
	}
	response, err := s.client.NewRequest().SetContext(ctx).SetResult(&apiv1.Module{}).SetPathParams(pathParams).Get("/{namespace}/{name}/{provider}/{version}")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.Module), nil
}

// List Modules
// https://developer.hashicorp.com/terraform/registry/api-docs#list-modules
func (s *Modules) List(ctx context.Context, q *apiv1.ListParams, namespace ...string) (*apiv1.ListModules, error) {
	var ns string
	if len(namespace) > 0 {
		ns = namespace[0]
	}
	values, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	response, err := s.client.NewRequest().SetContext(ctx).SetResult(&apiv1.ListModules{}).SetQueryParamsFromValues(values).Get(ns)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.ListModules), nil
}

// ListAvailableVersions List Available Versions for a Specific Module
// https://developer.hashicorp.com/terraform/registry/api-docs#list-available-versions-for-a-specific-module
func (s *Modules) ListAvailableVersions(ctx context.Context, namespace, name, provider string) (*apiv1.ListAvailableVersions, error) {
	pathParams := map[string]string{
		"namespace": namespace,
		"name":      name,
		"provider":  provider,
	}
	response, err := s.client.NewRequest().SetContext(ctx).SetResult(&apiv1.ListAvailableVersions{}).SetPathParams(pathParams).Get("/{namespace}/{name}/{provider}/versions")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.ListAvailableVersions), nil
}

// ListLatestVersionForAllProviders List Latest Version of Module for All Providers
// https://developer.hashicorp.com/terraform/registry/api-docs#list-latest-version-of-module-for-all-providers
func (s *Modules) ListLatestVersionForAllProviders(ctx context.Context, q *apiv1.ListLatestVersionParams, namespace, name string) (*apiv1.ListModules, error) {
	values, err := query.Values(q)
	if err != nil {
		return nil, err
	}
	pathParams := map[string]string{
		"namespace": namespace,
		"name":      name,
	}
	response, err := s.client.NewRequest().SetContext(ctx).SetResult(&apiv1.ListModules{}).SetPathParams(pathParams).SetQueryParamsFromValues(values).Get("/{namespace}/{name}")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.ListModules), nil
}

// LatestVersionForSpecificProvider Latest Version of Module for Specific Provider
// https://developer.hashicorp.com/terraform/registry/api-docs#latest-version-for-a-specific-module-provider
func (s *Modules) LatestVersionForSpecificProvider(ctx context.Context, namespace, name, provider string) (*apiv1.Module, error) {
	pathParams := map[string]string{
		"namespace": namespace,
		"name":      name,
		"provider":  provider,
	}
	response, err := s.client.NewRequest().SetContext(ctx).SetResult(&apiv1.Module{}).SetPathParams(pathParams).Get("/{namespace}/{name}/{provider}")
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	return response.Result().(*apiv1.Module), nil
}
