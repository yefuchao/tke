/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	platform "tkestack.io/tke/api/platform"
)

// CronHPAsGetter has a method to return a CronHPAInterface.
// A group's client should implement this interface.
type CronHPAsGetter interface {
	CronHPAs() CronHPAInterface
}

// CronHPAInterface has methods to work with CronHPA resources.
type CronHPAInterface interface {
	Create(ctx context.Context, cronHPA *platform.CronHPA, opts v1.CreateOptions) (*platform.CronHPA, error)
	Update(ctx context.Context, cronHPA *platform.CronHPA, opts v1.UpdateOptions) (*platform.CronHPA, error)
	UpdateStatus(ctx context.Context, cronHPA *platform.CronHPA, opts v1.UpdateOptions) (*platform.CronHPA, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*platform.CronHPA, error)
	List(ctx context.Context, opts v1.ListOptions) (*platform.CronHPAList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.CronHPA, err error)
	CronHPAExpansion
}

// cronHPAs implements CronHPAInterface
type cronHPAs struct {
	client rest.Interface
}

// newCronHPAs returns a CronHPAs
func newCronHPAs(c *PlatformClient) *cronHPAs {
	return &cronHPAs{
		client: c.RESTClient(),
	}
}

// Get takes name of the cronHPA, and returns the corresponding cronHPA object, and an error if there is any.
func (c *cronHPAs) Get(ctx context.Context, name string, options v1.GetOptions) (result *platform.CronHPA, err error) {
	result = &platform.CronHPA{}
	err = c.client.Get().
		Resource("cronhpas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CronHPAs that match those selectors.
func (c *cronHPAs) List(ctx context.Context, opts v1.ListOptions) (result *platform.CronHPAList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &platform.CronHPAList{}
	err = c.client.Get().
		Resource("cronhpas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cronHPAs.
func (c *cronHPAs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cronhpas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cronHPA and creates it.  Returns the server's representation of the cronHPA, and an error, if there is any.
func (c *cronHPAs) Create(ctx context.Context, cronHPA *platform.CronHPA, opts v1.CreateOptions) (result *platform.CronHPA, err error) {
	result = &platform.CronHPA{}
	err = c.client.Post().
		Resource("cronhpas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cronHPA).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cronHPA and updates it. Returns the server's representation of the cronHPA, and an error, if there is any.
func (c *cronHPAs) Update(ctx context.Context, cronHPA *platform.CronHPA, opts v1.UpdateOptions) (result *platform.CronHPA, err error) {
	result = &platform.CronHPA{}
	err = c.client.Put().
		Resource("cronhpas").
		Name(cronHPA.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cronHPA).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cronHPAs) UpdateStatus(ctx context.Context, cronHPA *platform.CronHPA, opts v1.UpdateOptions) (result *platform.CronHPA, err error) {
	result = &platform.CronHPA{}
	err = c.client.Put().
		Resource("cronhpas").
		Name(cronHPA.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cronHPA).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cronHPA and deletes it. Returns an error if one occurs.
func (c *cronHPAs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cronhpas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cronHPA.
func (c *cronHPAs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *platform.CronHPA, err error) {
	result = &platform.CronHPA{}
	err = c.client.Patch(pt).
		Resource("cronhpas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
