package static

var StaticFakesClientSetEventNamespacedExpansion = `
func (c *eventsClient) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := kcptesting.NewRootCreateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = kcptesting.NewCreateAction(eventsResource, c.ClusterPath, c.Namespace, event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *eventsClient) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := kcptesting.NewRootUpdateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = kcptesting.NewUpdateAction(eventsResource, c.ClusterPath, c.Namespace, event)
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
	action := kcptesting.NewRootPatchAction(eventsResource, c.ClusterPath, event.Name, pt, data)
	if c.Namespace != "" {
		action = kcptesting.NewPatchAction(eventsResource, c.ClusterPath, c.Namespace, event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
func (c *eventsClient) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	action := kcptesting.NewRootListAction(eventsResource, eventsKind, c.ClusterPath, metav1.ListOptions{})
	if c.Namespace != "" {
		action = kcptesting.NewListAction(eventsResource, eventsKind, c.ClusterPath, c.Namespace, metav1.ListOptions{})
	}
	obj, err := c.Fake.Invokes(action, &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.EventList), err
}

func (c *eventsClient) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	action := kcptesting.GenericActionImpl{}
	action.Verb = "get-field-selector"
	action.Resource = eventsResource
	action.ClusterPath = c.ClusterPath

	_, _ = c.Fake.Invokes(action, nil)
	return fields.Everything()
}
`

var StaticFakesClientSetEventExpansion = `

func (c *eventsClient) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := kcptesting.NewRootCreateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = kcptesting.NewCreateAction(eventsResource, c.ClusterPath, c.Namespace, event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *eventsClient) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := kcptesting.NewRootUpdateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = kcptesting.NewUpdateAction(eventsResource, c.ClusterPath, c.Namespace, event)
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
	action := kcptesting.NewRootPatchAction(eventsResource, c.ClusterPath, event.Name, pt, data)
	if c.Namespace != "" {
		action = kcptesting.NewPatchAction(eventsResource, c.ClusterPath, c.Namespace, event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
func (c *eventsClient) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	action := kcptesting.NewRootListAction(eventsResource, eventsKind, c.ClusterPath, metav1.ListOptions{})
	if c.Namespace != "" {
		action = kcptesting.NewListAction(eventsResource, eventsKind, c.ClusterPath, c.Namespace, metav1.ListOptions{})
	}
	obj, err := c.Fake.Invokes(action, &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.EventList), err
}

func (c *eventsClient) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	action := kcptesting.GenericActionImpl{}
	action.Verb = "get-field-selector"
	action.Resource = eventsResource
	action.ClusterPath = c.ClusterPath

	_, _ = c.Fake.Invokes(action, nil)
	return fields.Everything()
}
`

var StaticFakesClientSetNamespacesExpansion = `
func (c *namespacesClient) Finalize(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error) {
	action := kcptesting.CreateActionImpl{}
	action.Verb = "create"
	action.Resource = namespacesResource
	action.Subresource = "finalize"
	action.Object = namespace
	action.ClusterPath = c.ClusterPath

	obj, err := c.Fake.Invokes(action, namespace)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Namespace), err
}
`

var StaticFakesClientSetNodesExpansion = `
// TODO: Should take a PatchType as an argument probably.
func (c *nodesClient) PatchStatus(_ context.Context, nodeName string, data []byte) (*v1.Node, error) {
	// TODO: Should be configurable to support additional patch strategies.
	pt := types.StrategicMergePatchType
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(nodesResource, c.ClusterPath, nodeName, pt, data, "status"), &v1.Node{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Node), err
}
`

var StaticFakesClientSetPodsExpansion = `

func (c *podsClient) Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error {
	action := kcptesting.CreateActionImpl{}
	action.Verb = "create"
	action.Namespace = binding.Namespace
	action.Resource = podsResource
	action.Subresource = "binding"
	action.Object = binding
	action.ClusterPath = c.ClusterPath

	_, err := c.Fake.Invokes(action, binding)
	return err
}

func (c *podsClient) GetBinding(name string) (result *v1.Binding, err error) {
	obj, err := c.Fake.
		Invokes(kcptesting.NewGetSubresourceAction(podsResource, c.ClusterPath, c.Namespace, "binding", name), &v1.Binding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Binding), err
}

func (c *podsClient) GetLogs(name string, opts *v1.PodLogOptions) *restclient.Request {
	action := kcptesting.GenericActionImpl{}
	action.Verb = "get"
	action.Namespace = c.Namespace
	action.Resource = podsResource
	action.Subresource = "log"
	action.Value = opts
	action.ClusterPath = c.ClusterPath

	_, _ = c.Fake.Invokes(action, &v1.Pod{})
	fakeClient := &fakerest.RESTClient{
		Client: fakerest.CreateHTTPClient(func(request *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("fake logs")),
			}
			return resp, nil
		}),
		NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
		GroupVersion:         podsKind.GroupVersion(),
		VersionedAPIPath:     fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/log", c.Namespace, name),
	}
	return fakeClient.Request()
}

func (c *podsClient) Evict(ctx context.Context, eviction *policyv1beta1.Eviction) error {
	return c.EvictV1beta1(ctx, eviction)
}

func (c *podsClient) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	action := kcptesting.CreateActionImpl{}
	action.Verb = "create"
	action.Namespace = c.Namespace
	action.Resource = podsResource
	action.Subresource = "eviction"
	action.Object = eviction
	action.ClusterPath = c.ClusterPath

	_, err := c.Fake.Invokes(action, eviction)
	return err
}

func (c *podsClient) EvictV1beta1(ctx context.Context, eviction *policyv1beta1.Eviction) error {
	action := kcptesting.CreateActionImpl{}
	action.Verb = "create"
	action.Namespace = c.Namespace
	action.Resource = podsResource
	action.Subresource = "eviction"
	action.Object = eviction
	action.ClusterPath = c.ClusterPath

	_, err := c.Fake.Invokes(action, eviction)
	return err
}

func (c *podsClient) ProxyGet(scheme, name, port, path string, params map[string]string) restclient.ResponseWrapper {
	return c.Fake.InvokesProxy(kcptesting.NewProxyGetAction(podsResource, c.ClusterPath, c.Namespace, scheme, name, port, path, params))
}

func (c *podsClient) UpdateEphemeralContainers(ctx context.Context, podName string, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(podsResource, c.ClusterPath, "ephemeralcontainers", c.Namespace, pod), &v1.Pod{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}
`

var StaticFakesClientSetServicesExpansion = `
func (c *servicesClient) ProxyGet(scheme, name, port, path string, params map[string]string) restclient.ResponseWrapper {
	return c.Fake.InvokesProxy(kcptesting.NewProxyGetAction(servicesResource, c.ClusterPath, c.Namespace, scheme, name, port, path, params))
}
`

var StaticFakesClientSetServiceAccountsExpansion = `
func (c *serviceAccountsClient) CreateToken(ctx context.Context, serviceAccountName string, tokenRequest *authenticationv1.TokenRequest, opts metav1.CreateOptions) (*authenticationv1.TokenRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateSubresourceAction(serviceaccountsResource, c.ClusterPath, serviceAccountName, "token", c.Namespace, tokenRequest), &authenticationv1.TokenRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authenticationv1.TokenRequest), err
}
`
