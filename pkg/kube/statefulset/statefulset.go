package statefulset

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(client k8sclient.Client) Client {
	return Client{
		client: client,
	}
}

type Client struct {
	client k8sclient.Client
}

// Get provides a thin wrapper and client.Client to access appsv1.StatefulSet types
func (c Client) Get(key k8sclient.ObjectKey) (appsv1.StatefulSet, error) {
	set := appsv1.StatefulSet{}
	if err := c.client.Get(context.TODO(), key, &set); err != nil {
		return appsv1.StatefulSet{}, err
	}
	return set, nil
}

// Update provides a thin wrapper and client.Client to update appsv1.StatefulSet types
func (c Client) Update(set appsv1.StatefulSet) error {
	if err := c.client.Update(context.TODO(), &set); err != nil {
		return err
	}
	return nil
}

// Create provides a thin wrapper and client.Client to create appsv1.StatefulSet types
func (c Client) Create(set appsv1.StatefulSet) error {
	if err := c.client.Create(context.TODO(), &set); err != nil {
		return err
	}
	return nil
}

// Delete provides a thin wrapper and client.Client to delete appsv1.StatefulSet types
func (c Client) Delete(set appsv1.StatefulSet) error {
	if err := c.client.Delete(context.TODO(), &set); err != nil {
		return err
	}
	return nil
}

// CreateOrUpdate will either Create the stateful set if it doesn't exist, or Update it
// if it does
func (c Client) CreateOrUpdate(set appsv1.StatefulSet) error {
	_, err := c.Get(types.NamespacedName{Name: set.Name, Namespace: set.Namespace})
	if err != nil && errors.IsNotFound(err) {
		return c.Create(set)
	} else if err != nil {
		return err
	}
	if err := c.Update(set); err != nil {
		return err
	}
	return nil
}