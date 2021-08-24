package ltree_visualizer

//VisualizerSchema: Contract to send to ltree-visualizer
type VisualizerSchema struct {
	Data []data `json:"data"`
}

type data struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}