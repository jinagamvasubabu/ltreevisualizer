package ltree_visualizer

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilsTestSuite struct {
	suite.Suite
}

func TestUtilsTestSuite(t *testing.T) {
	tests := new(UtilsTestSuite)
	suite.Run(t, tests)
}

func (suite *UtilsTestSuite) SetupTest() {
}

func (suite *UtilsTestSuite) TestContains_Key_Found() {
	//Given
	list := []string{"apple", "orange", "banana"}
	toFind := "banana"

	//When
	res := Contains(list, toFind)

	//Then
	suite.True(res)
}

func (suite *UtilsTestSuite) TestContains_Key_Not_Found() {
	//Given
	list := []string{"apple", "orange", "banana"}
	toFind := "banana1"

	//When
	res := Contains(list, toFind)

	//Then
	suite.False(res)
}
