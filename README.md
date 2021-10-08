# Ltree Visualizer
[![Build Status](https://circleci.com/gh/jinagamvasubabu/ltreevisualizer.svg?style=svg)](https://circleci.com/gh/jinagamvasubabu/ltreevisualizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/jinagamvasubabu/ltreevisualizer)](https://goreportcard.com/report/github.com/jinagamvasubabu/ltreevisualizer)
[![GoDoc](https://godoc.org/github.com/jinagamvasubabu/ltreevisualizer?status.svg)](https://godoc.org/github.com/jinagamvasubabu/ltreevisualizer) 
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Visits Badge](https://badges.pufler.dev/visits/jinagamvasubabu/ltreevisualizer)](https://badges.pufler.dev)

A golang library to visualize or display postgres ltree type data directly from Postgres DB using DOT language and Graphviz. 

![alt text](https://github.com/jinagamvasubabu/ltreevisualizer/blob/main/images/LtreeVisualizer.jpg?raw=true)

# What is Ltree?

Ltree is a data type which is used to represent the hierarchical tree-like structures in flat rows and columns in
postgres DB For more info-refer this https://www.postgresql.org/docs/9.1/ltree.html

Sample Hierarchy:
![alt text](https://github.com/jinagamvasubabu/ltreevisualizer/blob/main/examples/graph.png?raw=true)

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
![alt text](https://github.com/jinagamvasubabu/ltreevisualizer/blob/main/images/DotLanguageDirected.png?raw=true)

# Config:
```go
//Visualizer config
type Visualizer struct {
    LogLevel    log.Level
    RankDir     string
    PostgresURI string //Example postgresql://postgres:postgres@localhost:5432/taxonomy?sslmode=disable
    Query       string //select id, name, path from table1 //columns specified in this example should match or use resultset alias if your column names are different
    FetchFromDB bool
}
```
RankDir: Sets the direction of tree layout(https://www.graphviz.org/docs/attrs/rankdir/) and supported values are
* TB (Top to Bottom)
* RL (Right to Left)
* LR (Left to Right)
* BT (Bottom to Top)
  Note: Default is TB

FilePath: FilePath to save the image, this parameter is optional for `GenerateDotGraph`.
Note: Default value of FilePath is `graph.png`

You can generate image of your ltree data using two ways:
* Directly fetch the data from your Postgres DB
* using Interim JSON file

## DB Way (Directly fetch the data from your Postgres DB):
```go
//Visualizer config
type Visualizer struct {
    LogLevel    log.Level
    RankDir     string
    PostgresURI string //Example postgresql://postgres:postgres@localhost:5432/taxonomy?sslmode=disable
    Query       string //select id, name, path from table1 //columns specified in this example should match or use resultset alias if your column names are different
    FetchFromDB bool
}
```
1. Provide `PostgresURI` which your app can connect Eg: `postgresql://postgres:postgres@localhost:5432/taxonomy?sslmode=disable`
2. Provide `Query` to fetch the data of your Ltree and your query result set should contain id, name, path
3. Set `FetchFromDB` to true

```go
  import "github.com/jinagamvasubabu/ltreevisualizer"
  import "github.com/sirupsen/logrus"
 
  l := ltreevisualizer.Visualizer{
        PostgresURI: "postgresql://postgres:postgres@localhost:5432/taxonomy?sslmode=disable",
        Query: "select id as id, name as name,path as path from table"
  }
  resp, err := l.ConvertLtreeDataToImage(context.Background(), ltreevisualizer.VisualizerSchema{})
  fmt.Println(resp)
```

## Using Interim Json File:
If you don't want to connect to DB and fetch the results then you can follow this way by using interim JSON file
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


* get `LtreeVisualizer`

```
  go get github.com/jinagamvasubabu/ltreevisualizer
```

* import and use it like below for to generate the output in DOT graph string:

```
  import "github.com/jinagamvasubabu/ltreevisualizer"
  import "github.com/sirupsen/logrus"
 
  l := ltreevisualizer.Visualizer{}
  resp, err := l.GenerateDotGraph(context.Background(), //json data)
  fmt.Println(resp)
```

* import and use it like below for to generate the output as an image:

```
  import "github.com/jinagamvasubabu/ltreevisualizer"
  import "github.com/sirupsen/logrus"
 
  l := ltreevisualizer.Visualizer{}
  err := l.ConvertLtreeDataToImage(context.Background(), //json data)
```

Note: This will create a graph.png image if you don't specify Filepath

You can refer `examples` directory for more info

# How to test?
Refer `visualizer_test.go` for sample tests
