// Copyright 2022 fishfinal <contact.fishfinal@gmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formatalign

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestFormatAlignSuite(t *testing.T) {
	suite.Run(t, &FormatAlignSuite{})
}

type FormatAlignSuite struct {
	suite.Suite
	alignOfRight *FormatAlign
	alignOfLeft  *FormatAlign
}

func (suite *FormatAlignSuite) BeforeTest(suiteName, testName string) {
	suite.alignOfLeft = NewFormatAlign()
	suite.alignOfRight = NewFormatAlign(WithAlignment(AlignmentOfRight))
}

func (suite *FormatAlignSuite) Test_FormatAlign_Left() {
	actual := suite.alignOfLeft.Format("username", "fishfinal")
	assert.Equal(suite.T(), `username                  : fishfinal`, actual)

	actual = suite.alignOfLeft.SetSeparator("").Format("username", "fishfinal")
	assert.Equal(suite.T(), `username                   fishfinal`, actual)

	actual = suite.alignOfLeft.SetSeparator("-").Format("username", "fishfinal")
	assert.Equal(suite.T(), `username                  - fishfinal`, actual)
}

func (suite *FormatAlignSuite) Test_FormatAlign_Right() {
	actual := suite.alignOfRight.Format("username", "fishfinal")
	assert.Equal(suite.T(), `                  username: fishfinal`, actual)

	actual = suite.alignOfRight.SetSeparator("").Format("username", "fishfinal")
	assert.Equal(suite.T(), `                  username fishfinal`, actual)

	actual = suite.alignOfRight.SetSeparator("-").Format("username", "fishfinal")
	assert.Equal(suite.T(), `                  username- fishfinal`, actual)
}
