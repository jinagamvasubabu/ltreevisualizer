# ltree data visualizer

A golang library to visualize postgres ltree type data using DOT language and Graphviz

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

## How to use?

* get `LtreeVisualizer`

```
  go get github.com/jinagamvasubabu/LtreeVisualizer
```

* import and use it like below:

```
  import "github.com/jinagamvasubabu/LtreeVisualizer"
  import "github.com/sirupsen/logrus"
 
  l := LtreeVisualizer.Visualizer{}
  resp, err := l.GenerateDotGraph(context.Background(), //json data)
  fmt.Println(resp)
```
