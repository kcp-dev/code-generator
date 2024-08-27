package static

var StaticFakesClientSetEventExpansion = `
func (c *eventsClient) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := core.NewRootCreateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = core.NewCreateAction(eventsResource, c.ClusterPath, c.Namespace, event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *eventsClient) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := core.NewRootUpdateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = core.NewUpdateAction(eventsResource, c.ClusterPath, c.Namespace, event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// PatchWithEventNamespace patches an existing event. Returns the copy of the event the server returns, or an error.
// TODO: Should take a PatchType as an argument probably.
func (c *eventsClient) PatchWithEventNamespace(event *v1.Event, data []byte) (*v1.Event, error) {
	// TODO: Should be configurable to support additional patch strategies.
	pt := types.StrategicMergePatchType
	action := core.NewRootPatchAction(eventsResource, c.ClusterPath, event.Name, pt, data)
	if c.Namespace != "" {
		action = core.NewPatchAction(eventsResource, c.ClusterPath, c.Namespace, event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
func (c *eventsClient) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	action := core.NewRootListAction(eventsResource, eventsKind, c.ClusterPath, metav1.ListOptions{})
	if c.Namespace != "" {
		action = core.NewListAction(eventsResource, eventsKind, c.ClusterPath, c.Namespace, metav1.ListOptions{})
	}
	obj, err := c.Fake.Invokes(action, &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.EventList), err
}

func (c *eventsClient) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	action := core.GenericActionImpl{}
	action.Verb = "get-field-selector"
	action.Resource = eventsResource
	action.ClusterPath = c.ClusterPath

	_, _ = c.Fake.Invokes(action, nil)
	return fields.Everything()
}



`
