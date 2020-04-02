/*
Copyright 2020 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	messagingv1alpha1 "knative.dev/eventing-contrib/natss/pkg/apis/messaging/v1alpha1"
	versioned "knative.dev/eventing-contrib/natss/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing-contrib/natss/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/eventing-contrib/natss/pkg/client/listers/messaging/v1alpha1"
)

// NatssChannelInformer provides access to a shared informer and lister for
// NatssChannels.
type NatssChannelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NatssChannelLister
}

type natssChannelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNatssChannelInformer constructs a new informer for NatssChannel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNatssChannelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNatssChannelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNatssChannelInformer constructs a new informer for NatssChannel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNatssChannelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MessagingV1alpha1().NatssChannels(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MessagingV1alpha1().NatssChannels(namespace).Watch(options)
			},
		},
		&messagingv1alpha1.NatssChannel{},
		resyncPeriod,
		indexers,
	)
}

func (f *natssChannelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNatssChannelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *natssChannelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&messagingv1alpha1.NatssChannel{}, f.defaultInformer)
}

func (f *natssChannelInformer) Lister() v1alpha1.NatssChannelLister {
	return v1alpha1.NewNatssChannelLister(f.Informer().GetIndexer())
}