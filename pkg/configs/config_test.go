package configs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) SetupSuite() {
	setup(".", "test")
	load()
}
func (s *ConfigTestSuite) TestNested() {
	s.Assertions.Equal(Get().DB.SSL, true)
}
func (s *ConfigTestSuite) TestDefaultNested() {
	s.Assertions.Equal(Get().DB.HOST, "localhost")
}
func TestConfigLoader(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
