// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./clients.go -destination mocks/clients.go

package v1alpha3

import (
	"context"

	split_smi_spec_io_v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha3"
	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulticlusterClientset for the split.smi-spec.io/v1alpha3 APIs
type MulticlusterClientset interface {
	// Cluster returns a Clientset for the given cluster
	Cluster(cluster string) (Clientset, error)
}

type multiclusterClientset struct {
	client multicluster.Client
}

func NewMulticlusterClientset(client multicluster.Client) MulticlusterClientset {
	return &multiclusterClientset{client: client}
}

func (m *multiclusterClientset) Cluster(cluster string) (Clientset, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

// clienset for the split.smi-spec.io/v1alpha3 APIs
type Clientset interface {
	// clienset for the split.smi-spec.io/v1alpha3/v1alpha3 APIs
	TrafficSplits() TrafficSplitClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (Clientset, error) {
	scheme := scheme.Scheme
	if err := split_smi_spec_io_v1alpha3.AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) Clientset {
	return &clientSet{client: client}
}

// clienset for the split.smi-spec.io/v1alpha3/v1alpha3 APIs
func (c *clientSet) TrafficSplits() TrafficSplitClient {
	return NewTrafficSplitClient(c.client)
}

// Reader knows how to read and list TrafficSplits.
type TrafficSplitReader interface {
	// Get retrieves a TrafficSplit for the given object key
	GetTrafficSplit(ctx context.Context, key client.ObjectKey) (*split_smi_spec_io_v1alpha3.TrafficSplit, error)

	// List retrieves list of TrafficSplits for a given namespace and list options.
	ListTrafficSplit(ctx context.Context, opts ...client.ListOption) (*split_smi_spec_io_v1alpha3.TrafficSplitList, error)
}

// TrafficSplitTransitionFunction instructs the TrafficSplitWriter how to transition between an existing
// TrafficSplit object and a desired on an Upsert
type TrafficSplitTransitionFunction func(existing, desired *split_smi_spec_io_v1alpha3.TrafficSplit) error

// Writer knows how to create, delete, and update TrafficSplits.
type TrafficSplitWriter interface {
	// Create saves the TrafficSplit object.
	CreateTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, opts ...client.CreateOption) error

	// Delete deletes the TrafficSplit object.
	DeleteTrafficSplit(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given TrafficSplit object.
	UpdateTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, opts ...client.UpdateOption) error

	// Patch patches the given TrafficSplit object.
	PatchTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all TrafficSplit objects matching the given options.
	DeleteAllOfTrafficSplit(ctx context.Context, opts ...client.DeleteAllOfOption) error

	// Create or Update the TrafficSplit object.
	UpsertTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, transitionFuncs ...TrafficSplitTransitionFunction) error
}

// StatusWriter knows how to update status subresource of a TrafficSplit object.
type TrafficSplitStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given TrafficSplit object.
	UpdateTrafficSplitStatus(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, opts ...client.UpdateOption) error

	// Patch patches the given TrafficSplit object's subresource.
	PatchTrafficSplitStatus(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on TrafficSplits.
type TrafficSplitClient interface {
	TrafficSplitReader
	TrafficSplitWriter
	TrafficSplitStatusWriter
}

type trafficSplitClient struct {
	client client.Client
}

func NewTrafficSplitClient(client client.Client) *trafficSplitClient {
	return &trafficSplitClient{client: client}
}

func (c *trafficSplitClient) GetTrafficSplit(ctx context.Context, key client.ObjectKey) (*split_smi_spec_io_v1alpha3.TrafficSplit, error) {
	obj := &split_smi_spec_io_v1alpha3.TrafficSplit{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *trafficSplitClient) ListTrafficSplit(ctx context.Context, opts ...client.ListOption) (*split_smi_spec_io_v1alpha3.TrafficSplitList, error) {
	list := &split_smi_spec_io_v1alpha3.TrafficSplitList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *trafficSplitClient) CreateTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *trafficSplitClient) DeleteTrafficSplit(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &split_smi_spec_io_v1alpha3.TrafficSplit{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *trafficSplitClient) UpdateTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *trafficSplitClient) PatchTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *trafficSplitClient) DeleteAllOfTrafficSplit(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &split_smi_spec_io_v1alpha3.TrafficSplit{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *trafficSplitClient) UpsertTrafficSplit(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, transitionFuncs ...TrafficSplitTransitionFunction) error {
	genericTxFunc := func(existing, desired runtime.Object) error {
		for _, txFunc := range transitionFuncs {
			if err := txFunc(existing.(*split_smi_spec_io_v1alpha3.TrafficSplit), desired.(*split_smi_spec_io_v1alpha3.TrafficSplit)); err != nil {
				return err
			}
		}
		return nil
	}
	_, err := controllerutils.Upsert(ctx, c.client, obj, genericTxFunc)
	return err
}

func (c *trafficSplitClient) UpdateTrafficSplitStatus(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *trafficSplitClient) PatchTrafficSplitStatus(ctx context.Context, obj *split_smi_spec_io_v1alpha3.TrafficSplit, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Provides TrafficSplitClients for multiple clusters.
type MulticlusterTrafficSplitClient interface {
	// Cluster returns a TrafficSplitClient for the given cluster
	Cluster(cluster string) (TrafficSplitClient, error)
}

type multiclusterTrafficSplitClient struct {
	client multicluster.Client
}

func NewMulticlusterTrafficSplitClient(client multicluster.Client) MulticlusterTrafficSplitClient {
	return &multiclusterTrafficSplitClient{client: client}
}

func (m *multiclusterTrafficSplitClient) Cluster(cluster string) (TrafficSplitClient, error) {
	client, err := m.client.Cluster(cluster)
	if err != nil {
		return nil, err
	}
	return NewTrafficSplitClient(client), nil
}
