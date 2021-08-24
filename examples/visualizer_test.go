package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"ltree-visualizer"
	"testing"
)

type VisualizerTestSuite struct {
	suite.Suite
	visualizer ltree_visualizer.Visualizer
}

func TestNewVisualizerTestSuite(t *testing.T) {
	tests := new(VisualizerTestSuite)

	suite.Run(t, tests)
}

func (suite *VisualizerTestSuite) SetupTest() {
	suite.visualizer = ltree_visualizer.Visualizer{

	}
}

//Example1: This test will generate the Dot Graph and print to the console
func (suite *VisualizerTestSuite) TestVisualizer_Generate_Dot_Graph_String() {
	//Given
	ltreeData := ltree_visualizer.VisualizerSchema{}
	data, _ := ioutil.ReadFile("data.json")
	err := json.Unmarshal(data, &ltreeData)
	suite.Nil(err)

	//When
	graphString, err := suite.visualizer.GenerateDotGraph(context.Background(), ltreeData)

	//Then
	suite.Nil(err)
	fmt.Printf("%s", graphString)
	suite.NotNil(graphString)
}

//Example2: This test will generate an image under examples directory
func (suite *VisualizerTestSuite) TestConvertLtreeDataToImage_Generate_Image_Success() {
	//Given
	ltreeData := ltree_visualizer.VisualizerSchema{}
	data, _ := ioutil.ReadFile("data.json")
	err := json.Unmarshal(data, &ltreeData)
	suite.Nil(err)

	//When
	err = suite.visualizer.ConvertLtreeDataToImage(context.Background(), ltreeData)

	//Then
	suite.Nil(err)
}

func (suite *VisualizerTestSuite) TestVisualizer_Validation_failure() {
	//Given
	ltreeData := ltree_visualizer.VisualizerSchema{}

	//When
	_, err := suite.visualizer.GenerateDotGraph(context.Background(), ltreeData)

	//Then
	suite.NotNil(err)
}


func (suite *VisualizerTestSuite) TestConvertLtreeDataToImage_Validation_Failure() {
	//Given
	ltreeData := ltree_visualizer.VisualizerSchema{}


	//When
	err := suite.visualizer.ConvertLtreeDataToImage(context.Background(), ltreeData)

	//Then
	suite.NotNil(err)
}

