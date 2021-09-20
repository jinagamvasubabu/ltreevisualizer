package ltreevisualizer

//VisualizerSchema Contract to send to ltreevisualizer
type VisualizerSchema struct {
	Data []Data `json:"Data"`
}

//Data Visualizer data
type Data struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}
