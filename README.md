# Ltree Visualizer
A golang library to visualize postgres ltree type data using DOT language and Graphviz

![alt text](https://github.com/jinagamvasubabu/LtreeVisualizer/blob/main/images/LtreeVisualizer.jpg?raw=true)

# What is Ltree?

Ltree is a data type which is used to represent the hierarchical tree-like structures in flat rows and columns in
postgres DB For more info-refer this https://www.postgresql.org/docs/9.1/ltree.html

Sample Hierarchy:
![alt text](https://github.com/jinagamvasubabu/LtreeVisualizer/blob/main/examples/graph.png?raw=true)

# why do we need this library ?

Ltree Labels are separated using Dot `.` like `1.2.3.4` and it is not easy to visualize like a tree.

This Library can produce output into two different formats:
* DOT Graph
* Image

# DOT Graph:
DOT is a graph description language, using this language we can represent Directed, Undirected and FlowCharts. https://en.wikipedia.org/wiki/DOT_(graph_description_language)
```go
digraph graphname {
    a -> b -> c;
    b -> d;
}
```
Using GraphViz we can visualize it like below:
![alt text](https://github.com/jinagamvasubabu/LtreeVisualizer/blob/main/images/DotLanguageDirected.png?raw=true)

## How to use?
* get `LtreeVisualizer`

```
  go get github.com/jinagamvasubabu/LtreeVisualizer
```

* import and use it like below for to generate the output in DOT graph string:

```
  import "github.com/jinagamvasubabu/LtreeVisualizer"
  import "github.com/sirupsen/logrus"
 
  l := LtreeVisualizer.Visualizer{}
  resp, err := l.GenerateDotGraph(context.Background(), //json data)
  fmt.Println(resp)
```

* import and use it like below for to generate the output as an image:

```
  import "github.com/jinagamvasubabu/LtreeVisualizer"
  import "github.com/sirupsen/logrus"
 
  l := LtreeVisualizer.Visualizer{}
  err := l.ConvertLtreeDataToImage(context.Background(), //json data)
```

Note: This will create a graph.png image if you don't specify Filepath 

# Config:
```go
//Visualizer config
type Visualizer struct {
	LogLevel log.Level
	RankDir  string
	FilePath string
}
```
RankDir: Supported values are
* TB (Top to Bottom)
* RL (Right to Left)
* LR (Left to Right)
* BT (Bottom to Top)
Note: Default is TB

FilePath: FilePath to save the image, this parameter is optional for `GenerateDotGraph`. 
Note: Default value of FilePath is `graph.png`

#Input:
Ltree Visualizer accepts data in this format
```go
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
```

Refer `data.json` file under examples directory for sample data

# How to test?
Refer `visualizer_test.go` for sample tests