/*
Copyright 2017 Gustavo Hoirisch.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package fake

import (
	v1alpha1 "github.com/gugahoi/memento/pkg/apis/registry/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRegistries implements RegistryInterface
type FakeRegistries struct {
	Fake *FakeMementoV1alpha1
	ns   string
}

var registriesResource = schema.GroupVersionResource{Group: "memento.gugahoi.com", Version: "v1alpha1", Resource: "registries"}

var registriesKind = schema.GroupVersionKind{Group: "memento.gugahoi.com", Version: "v1alpha1", Kind: "Registry"}

// Get takes name of the registry, and returns the corresponding registry object, and an error if there is any.
func (c *FakeRegistries) Get(name string, options v1.GetOptions) (result *v1alpha1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(registriesResource, c.ns, name), &v1alpha1.Registry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Registry), err
}

// List takes label and field selectors, and returns the list of Registries that match those selectors.
func (c *FakeRegistries) List(opts v1.ListOptions) (result *v1alpha1.RegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(registriesResource, registriesKind, c.ns, opts), &v1alpha1.RegistryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RegistryList{}
	for _, item := range obj.(*v1alpha1.RegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested registries.
func (c *FakeRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(registriesResource, c.ns, opts))

}

// Create takes the representation of a registry and creates it.  Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Create(registry *v1alpha1.Registry) (result *v1alpha1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(registriesResource, c.ns, registry), &v1alpha1.Registry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Registry), err
}

// Update takes the representation of a registry and updates it. Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Update(registry *v1alpha1.Registry) (result *v1alpha1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(registriesResource, c.ns, registry), &v1alpha1.Registry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Registry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegistries) UpdateStatus(registry *v1alpha1.Registry) (*v1alpha1.Registry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(registriesResource, "status", c.ns, registry), &v1alpha1.Registry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Registry), err
}

// Delete takes name of the registry and deletes it. Returns an error if one occurs.
func (c *FakeRegistries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(registriesResource, c.ns, name), &v1alpha1.Registry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegistries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(registriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RegistryList{})
	return err
}

// Patch applies the patch and returns the patched registry.
func (c *FakeRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(registriesResource, c.ns, name, data, subresources...), &v1alpha1.Registry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Registry), err
}
