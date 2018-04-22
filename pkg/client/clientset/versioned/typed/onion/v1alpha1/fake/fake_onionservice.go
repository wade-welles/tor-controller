/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	v1alpha1 "github.com/kragniz/kube-onions/pkg/apis/onion/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOnionServices implements OnionServiceInterface
type FakeOnionServices struct {
	Fake *FakeOnionV1alpha1
	ns   string
}

var onionservicesResource = schema.GroupVersionResource{Group: "onion.kragniz.eu", Version: "v1alpha1", Resource: "onionservices"}

var onionservicesKind = schema.GroupVersionKind{Group: "onion.kragniz.eu", Version: "v1alpha1", Kind: "OnionService"}

// Get takes name of the onionService, and returns the corresponding onionService object, and an error if there is any.
func (c *FakeOnionServices) Get(name string, options v1.GetOptions) (result *v1alpha1.OnionService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(onionservicesResource, c.ns, name), &v1alpha1.OnionService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnionService), err
}

// List takes label and field selectors, and returns the list of OnionServices that match those selectors.
func (c *FakeOnionServices) List(opts v1.ListOptions) (result *v1alpha1.OnionServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(onionservicesResource, onionservicesKind, c.ns, opts), &v1alpha1.OnionServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OnionServiceList{}
	for _, item := range obj.(*v1alpha1.OnionServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested onionServices.
func (c *FakeOnionServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(onionservicesResource, c.ns, opts))

}

// Create takes the representation of a onionService and creates it.  Returns the server's representation of the onionService, and an error, if there is any.
func (c *FakeOnionServices) Create(onionService *v1alpha1.OnionService) (result *v1alpha1.OnionService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(onionservicesResource, c.ns, onionService), &v1alpha1.OnionService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnionService), err
}

// Update takes the representation of a onionService and updates it. Returns the server's representation of the onionService, and an error, if there is any.
func (c *FakeOnionServices) Update(onionService *v1alpha1.OnionService) (result *v1alpha1.OnionService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(onionservicesResource, c.ns, onionService), &v1alpha1.OnionService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnionService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOnionServices) UpdateStatus(onionService *v1alpha1.OnionService) (*v1alpha1.OnionService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(onionservicesResource, "status", c.ns, onionService), &v1alpha1.OnionService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnionService), err
}

// Delete takes name of the onionService and deletes it. Returns an error if one occurs.
func (c *FakeOnionServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(onionservicesResource, c.ns, name), &v1alpha1.OnionService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOnionServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(onionservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OnionServiceList{})
	return err
}

// Patch applies the patch and returns the patched onionService.
func (c *FakeOnionServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OnionService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(onionservicesResource, c.ns, name, data, subresources...), &v1alpha1.OnionService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnionService), err
}
