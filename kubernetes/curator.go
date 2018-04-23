package kubernetes

//Curator holds configuration data and state
type Curator struct {
	kubernetesConfig *Configuration
	Filters          *Filters
	//Namespace
	Namespaces [][]string
	//Mapped by namespace
	Pods map[string][]string
	//Mapped by Namespace
	Services map[string][]string
	//Mapped by Namespace
	Deployments map[string][]string
	//Mapped by Namespace
	StatefulSets map[string][]string
}

//NewCurator creates a new struct
func NewCurator(conf *Configuration, f *Filters) *Curator {

	return &Curator{kubernetesConfig: conf, Filters: f, Namespaces: [][]string{
		[]string{"Namespaces"}}, Pods: make(map[string][]string), Services: make(map[string][]string),
		Deployments: make(map[string][]string), StatefulSets: make(map[string][]string)}
}

//Do runs the current filter set over the curator
func (c *Curator) Do() error {

	//Fetch namespaces----------------------------------------------------------
	nsl, err := c.kubernetesConfig.GetNamespaces()
	if err != nil {
		return err
	}
	c.Namespaces = [][]string{
		[]string{"Namespaces"}}

	for _, ns := range nsl.Items {
		sr := []string{ns.Name}
		//Fetch namespaces
		c.Namespaces = append(c.Namespaces, sr)
		//Fetch Deployments
		dsl, err := c.kubernetesConfig.GetDeployments(ns.Name)
		if err != nil {
			continue
		}
		for _, p := range dsl.Items {
			c.Deployments[ns.Name] = append(c.Deployments[ns.Name], p.Name)
		}
		stsl, err := c.kubernetesConfig.GetStatefulSets(ns.Name)
		if err != nil {
			continue
		}
		for _, p := range stsl.Items {
			c.StatefulSets[ns.Name] = append(c.StatefulSets[ns.Name], p.Name)
		}

		//Fetch Services
		ssl, err := c.kubernetesConfig.GetServices(ns.Name)
		if err != nil {
			continue
		}
		for _, p := range ssl.Items {
			c.Services[ns.Name] = append(c.Services[ns.Name], p.Name)
		}
		//Fetch Pods
		psl, err := c.kubernetesConfig.GetPods(ns.Name)
		if err != nil {
			continue
		}
		for _, p := range psl.Items {
			c.Pods[ns.Name] = append(c.Pods[ns.Name], p.Name)
		}
	}

	//--------------------------------------------------------------------------
	return nil

}
