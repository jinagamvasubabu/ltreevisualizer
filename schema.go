package ltreevisualizer

//VisualizerSchema Contract to send to ltreevisualizer
type VisualizerSchema struct {
	Data []data `json:"data"`
}

type data struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}
