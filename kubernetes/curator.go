package kubernetes

//Curator holds configuration data and state
type Curator struct {
	kubernetesConfig *Configuration
	Filters          *Filters
	//Namespace
	Namespaces [][]string
	Pods       map[string][]string
}

//NewCurator creates a new struct
func NewCurator(conf *Configuration, f *Filters) *Curator {

	return &Curator{kubernetesConfig: conf, Filters: f, Namespaces: [][]string{
		[]string{"Namespaces"}}, Pods: make(map[string][]string)}
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
		psl, err := c.kubernetesConfig.GetPods(ns.Name)
		if err != nil {
			continue
		}
		//Fetch Pods
		for _, p := range psl.Items {
			c.Pods[ns.Name] = append(c.Pods[ns.Name], p.Name)
		}
	}

	//--------------------------------------------------------------------------
	return nil

}
