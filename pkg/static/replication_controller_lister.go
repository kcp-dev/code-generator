package static

var StaticReplicationControllerListerExpansion = `
// ReplicationControllerListerExpansion allows custom methods to be added to
// ReplicationControllerLister.
type ReplicationControllerListerExpansion interface {
	GetPodControllers(pod *v1.Pod) ([]*v1.ReplicationController, error)
}

// GetPodControllers returns a list of ReplicationControllers that potentially match a pod.
// Only the one specified in the Pod's ControllerRef will actually manage it.
// Returns an error only if no matching ReplicationControllers are found.
func (s *replicationControllerLister) GetPodControllers(pod *v1.Pod) ([]*v1.ReplicationController, error) {
	if len(pod.Labels) == 0 {
		return nil, fmt.Errorf("no controllers found for pod %v because it has no labels", pod.Name)
	}

	items, err := s.ReplicationControllers(pod.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var controllers []*v1.ReplicationController
	for i := range items {
		rc := items[i]
		selector := labels.Set(rc.Spec.Selector).AsSelectorPreValidated()

		// If an rc with a nil or empty selector creeps in, it should match nothing, not everything.
		if selector.Empty() || !selector.Matches(labels.Set(pod.Labels)) {
			continue
		}
		controllers = append(controllers, rc)
	}

	if len(controllers) == 0 {
		return nil, fmt.Errorf("could not find controller for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
	}

	return controllers, nil
}

// ReplicationControllerClusterListerExpansion allows custom methods to be added to
// ReplicationControllerLister.
type ReplicationControllerClusterListerExpansion interface {
	// Cluster returns a lister that can list and get ReplicationController in one workspace.
	Cluster(clusterName logicalcluster.Name) ReplicationControllerLister
}

// ReplicationControllerNamespaceListerExpansion allows custom methods to be added to
// ReplicationControllerNamespaceLister.
type ReplicationControllerNamespaceListerExpansion interface{}
`
