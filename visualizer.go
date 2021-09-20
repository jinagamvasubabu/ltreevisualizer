package ltreevisualizer

import (
	"context"
	"errors"
	"fmt"
	"github.com/emicklei/dot"
	"github.com/goccy/go-graphviz"
	"github.com/jinagamvasubabu/ltreevisualizer/database"
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
	LogLevel    log.Level
	RankDir     string
	PostgresURI string //Example postgresql://postgres:postgres@localhost:5432/taxonomy?sslmode=disable
	Query       string //select id, name, path from table1 //columns specified in this example should match or use resultset alias if your column names are different
	FetchFromDB bool
}

//GenerateDotGraph generates a DOT graph string
func (v *Visualizer) GenerateDotGraph(ctx context.Context, ltreeData VisualizerSchema) (string, error) {
	log.SetLevel(v.LogLevel)
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "GenerateDotGraph"})
	defer CalculateTimeTaken(ctx, time.Now(), "Time Taken by GenerateDotGraph")
	//Check if DB details are present in the config then convert that Data to ltreeData
	if err := v.validateDBRequest(); err != nil {
		return "", err
	}
	if err := v.validateRequest(ltreeData); err != nil {
		logger.Debugf("Validation failed = %v", err.Error())
		return "", err
	}
	if v.FetchFromDB {
		var err error
		ltreeData, err = v.getLtreeDataFromPostgres(ctx)
		if err != nil {
			return "", err
		}
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
		logger.Debugf("Error while converting graph Data to image = %v", err.Error())
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

//getLtreeDataFromPostgres private method which helps to get the ltree data from the postgres DB
func (v *Visualizer) getLtreeDataFromPostgres(ctx context.Context) (VisualizerSchema, error) {
	log.SetLevel(v.LogLevel)
	logger := log.WithContext(ctx).WithFields(log.Fields{"Method": "getLtreeDataFromPostgres"})
	helper := database.NewDBHelper(v.PostgresURI)
	conn, err := helper.CreateDBConn()
	if err != nil {
		return VisualizerSchema{}, err
	}
	var data []Data
	if err := conn.Raw(v.Query).Scan(&data).Error; err != nil {
		logger.Error("error while fetching the ltree data from postgres", err.Error())
		return VisualizerSchema{}, fmt.Errorf("error while fetching the ltree data from postgres = %s", err.Error())
	}
	if len(data) == 0 {
		return VisualizerSchema{}, errors.New("no data available in postgres DB")
	}
	schema := VisualizerSchema{data}
	return schema, nil
}

//validateRequest validate the request
func (v *Visualizer) validateRequest(ltreeData VisualizerSchema) error {
	//validations
	if v.FetchFromDB == true {
		return nil
	}
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

//validateDBRequest check  all necessary information provided to fetch the information from DB
func (v *Visualizer) validateDBRequest() error {
	//validations
	if v.FetchFromDB == false {
		return nil
	}
	if v.PostgresURI == "" {
		return errors.New("PostgresURI is missing in the configuration")
	}
	if v.Query == "" {
		return errors.New("query is missing in the configuration")
	}
	return nil
}
