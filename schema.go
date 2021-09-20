package ltreevisualizer

//VisualizerSchema Contract to send to ltreevisualizer
type VisualizerSchema struct {
	Data []Data `json:"Data"`
}

type Data struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}
