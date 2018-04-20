package kubernetes

//Curator holds configuration data and state
type Curator struct {
	kubernetesConfig *Configuration
	Filters          *Filters
}

//NewCurator creates a new struct
func NewCurator(conf *Configuration, f *Filters) *Curator {

	return &Curator{kubernetesConfig: conf, Filters: f}
}
