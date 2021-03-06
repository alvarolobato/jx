package v1

import (
	v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PipelineActivitiesGetter has a method to return a PipelineActivityInterface.
// A group's client should implement this interface.
type PipelineActivitiesGetter interface {
	PipelineActivities(namespace string) PipelineActivityInterface
}

// PipelineActivityInterface has methods to work with PipelineActivity resources.
type PipelineActivityInterface interface {
	Create(*v1.PipelineActivity) (*v1.PipelineActivity, error)
	Update(*v1.PipelineActivity) (*v1.PipelineActivity, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.PipelineActivity, error)
	List(opts meta_v1.ListOptions) (*v1.PipelineActivityList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PipelineActivity, err error)
	PipelineActivityExpansion
}

// pipelineActivities implements PipelineActivityInterface
type pipelineActivities struct {
	client rest.Interface
	ns     string
}

// newPipelineActivities returns a PipelineActivities
func newPipelineActivities(c *JenkinsV1Client, namespace string) *pipelineActivities {
	return &pipelineActivities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pipelineActivity, and returns the corresponding pipelineActivity object, and an error if there is any.
func (c *pipelineActivities) Get(name string, options meta_v1.GetOptions) (result *v1.PipelineActivity, err error) {
	result = &v1.PipelineActivity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pipelineactivities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PipelineActivities that match those selectors.
func (c *pipelineActivities) List(opts meta_v1.ListOptions) (result *v1.PipelineActivityList, err error) {
	result = &v1.PipelineActivityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pipelineactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pipelineActivities.
func (c *pipelineActivities) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pipelineactivities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pipelineActivity and creates it.  Returns the server's representation of the pipelineActivity, and an error, if there is any.
func (c *pipelineActivities) Create(pipelineActivity *v1.PipelineActivity) (result *v1.PipelineActivity, err error) {
	result = &v1.PipelineActivity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pipelineactivities").
		Body(pipelineActivity).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pipelineActivity and updates it. Returns the server's representation of the pipelineActivity, and an error, if there is any.
func (c *pipelineActivities) Update(pipelineActivity *v1.PipelineActivity) (result *v1.PipelineActivity, err error) {
	result = &v1.PipelineActivity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pipelineactivities").
		Name(pipelineActivity.Name).
		Body(pipelineActivity).
		Do().
		Into(result)
	return
}

// Delete takes name of the pipelineActivity and deletes it. Returns an error if one occurs.
func (c *pipelineActivities) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pipelineactivities").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pipelineActivities) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pipelineactivities").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pipelineActivity.
func (c *pipelineActivities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PipelineActivity, err error) {
	result = &v1.PipelineActivity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pipelineactivities").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
