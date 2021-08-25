package ltreevisualizer

import (
	"context"
	"errors"
	"fmt"
	"github.com/emicklei/dot"
	"github.com/goccy/go-graphviz"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

//IVisualizer interface for to interact ltree visualizer
type IVisualizer interface {
	GenerateDotGraph(ctx context.Context, ltreeData VisualizerSchema) (string, error)
	ConvertLtreeDataToImage(ctx context.Context, ltreeData VisualizerSchema) error
}

//Visualizer config
type Visualizer struct {
	LogLevel log.Level
	RankDir  string
}

//GenerateDotGraph generates a DOT graph string
func (v *Visualizer) GenerateDotGraph(ctx context.Context, ltreeData VisualizerSchema) (string, error) {
	log.SetLevel(v.LogLevel)
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "GenerateDotGraph"})
	defer CalculateTimeTaken(ctx, time.Now(), "Time Taken by GenerateDotGraph")
	if err := v.validateRequest(ltreeData); err != nil {
		logger.Debugf("Validation failed = %v", err.Error())
		return "", err
	}
	//New Dot Graph instance
	g := dot.NewGraph(dot.Directed) //Directed graph
	if v.RankDir == "" || !Contains(GetSupportedRankDir(), v.RankDir) {
		logger.Debugf("Setting Default Rankdir to TB")
		v.RankDir = "TB" //Default is TB (top to bottom)
	}
	g.Attr("rankdir", v.RankDir)
	//Create unique edges
	edgeMap := map[string]dot.Edge{}
	//create a map for names to show in the nodes
	nodeMap := map[string]string{}
	for _, d := range ltreeData.Data {
		nodeMap[strconv.Itoa(int(d.ID))] = d.Name
	}
	for _, d := range ltreeData.Data {
		values := strings.Split(d.Path, ".")
		for i := 0; i+1 < len(values); i++ {
			if _, ok := edgeMap[values[i]+"->"+values[i+1]]; !ok {
				n1 := g.Node(nodeMap[values[i]])
				n2 := g.Node(nodeMap[values[i+1]])
				edgeMap[values[i]+"->"+values[i+1]] = g.Edge(n1, n2)
			}
		}
	}
	return g.String(), nil
}

//ConvertLtreeDataToImage Converts Ltree Data to an image
func (v *Visualizer) ConvertLtreeDataToImage(ctx context.Context, ltreeData VisualizerSchema) error {
	log.SetLevel(v.LogLevel)
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "ConvertLtreeDataToImage"})
	defer CalculateTimeTaken(ctx, time.Now(), "ConvertLtreeDataToImage")
	dotGraphStr, err := v.GenerateDotGraph(ctx, ltreeData)
	if err != nil {
		logger.Debugf("Error while converting graph data to image = %v", err.Error())
		return err
	}
	graph, err := graphviz.ParseBytes([]byte(dotGraphStr))
	g := graphviz.New()
	if err := g.RenderFilename(graph, graphviz.PNG, "graph.png"); err != nil {
		log.Fatal(err)
		return fmt.Errorf("error while generating image")
	}
	return nil
}

//validateRequest validate the request
func (v *Visualizer) validateRequest(ltreeData VisualizerSchema) error {
	//validations
	if len(ltreeData.Data) == 0 {
		return errors.New("ltreeData is missing in the request")
	}
	var uniqueNames []string
	for _, e := range ltreeData.Data {
		if Contains(uniqueNames, e.Name) {
			return errors.New("names should be unique")
		}
		uniqueNames = append(uniqueNames, e.Name)
	}
	return nil
}