// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	federationv1 "maistra.io/api/federation/v1"
)

// FakeImportedServiceSets implements ImportedServiceSetInterface
type FakeImportedServiceSets struct {
	Fake *FakeFederationV1
	ns   string
}

var importedservicesetsResource = schema.GroupVersionResource{Group: "federation.maistra.io", Version: "v1", Resource: "importedservicesets"}

var importedservicesetsKind = schema.GroupVersionKind{Group: "federation.maistra.io", Version: "v1", Kind: "ImportedServiceSet"}

// Get takes name of the importedServiceSet, and returns the corresponding importedServiceSet object, and an error if there is any.
func (c *FakeImportedServiceSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *federationv1.ImportedServiceSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(importedservicesetsResource, c.ns, name), &federationv1.ImportedServiceSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.ImportedServiceSet), err
}

// List takes label and field selectors, and returns the list of ImportedServiceSets that match those selectors.
func (c *FakeImportedServiceSets) List(ctx context.Context, opts v1.ListOptions) (result *federationv1.ImportedServiceSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(importedservicesetsResource, importedservicesetsKind, c.ns, opts), &federationv1.ImportedServiceSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &federationv1.ImportedServiceSetList{ListMeta: obj.(*federationv1.ImportedServiceSetList).ListMeta}
	for _, item := range obj.(*federationv1.ImportedServiceSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested importedServiceSets.
func (c *FakeImportedServiceSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(importedservicesetsResource, c.ns, opts))

}

// Create takes the representation of a importedServiceSet and creates it.  Returns the server's representation of the importedServiceSet, and an error, if there is any.
func (c *FakeImportedServiceSets) Create(ctx context.Context, importedServiceSet *federationv1.ImportedServiceSet, opts v1.CreateOptions) (result *federationv1.ImportedServiceSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(importedservicesetsResource, c.ns, importedServiceSet), &federationv1.ImportedServiceSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.ImportedServiceSet), err
}

// Update takes the representation of a importedServiceSet and updates it. Returns the server's representation of the importedServiceSet, and an error, if there is any.
func (c *FakeImportedServiceSets) Update(ctx context.Context, importedServiceSet *federationv1.ImportedServiceSet, opts v1.UpdateOptions) (result *federationv1.ImportedServiceSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(importedservicesetsResource, c.ns, importedServiceSet), &federationv1.ImportedServiceSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.ImportedServiceSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImportedServiceSets) UpdateStatus(ctx context.Context, importedServiceSet *federationv1.ImportedServiceSet, opts v1.UpdateOptions) (*federationv1.ImportedServiceSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(importedservicesetsResource, "status", c.ns, importedServiceSet), &federationv1.ImportedServiceSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.ImportedServiceSet), err
}

// Delete takes name of the importedServiceSet and deletes it. Returns an error if one occurs.
func (c *FakeImportedServiceSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(importedservicesetsResource, c.ns, name, opts), &federationv1.ImportedServiceSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImportedServiceSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(importedservicesetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &federationv1.ImportedServiceSetList{})
	return err
}

// Patch applies the patch and returns the patched importedServiceSet.
func (c *FakeImportedServiceSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *federationv1.ImportedServiceSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(importedservicesetsResource, c.ns, name, pt, data, subresources...), &federationv1.ImportedServiceSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.ImportedServiceSet), err
}
