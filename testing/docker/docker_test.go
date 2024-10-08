package docker

import (
	"testing"

	"github.com/stretchr/testify/suite"

	configmocks "github.com/goravel/framework/mocks/config"
	foundationmocks "github.com/goravel/framework/mocks/foundation"
)

type DockerTestSuite struct {
	suite.Suite
	mockApp *foundationmocks.Application
	docker  *Docker
}

func TestDockerTestSuite(t *testing.T) {
	suite.Run(t, new(DockerTestSuite))
}

func (s *DockerTestSuite) SetupTest() {
	s.mockApp = &foundationmocks.Application{}
	s.docker = NewDocker(s.mockApp)
}

func (s *DockerTestSuite) TestDatabase() {
	mockConfig := &configmocks.Config{}
	mockConfig.On("GetString", "database.default").Return("mysql").Once()
	mockConfig.On("GetString", "database.connections.mysql.driver").Return("mysql").Once()
	mockConfig.On("GetString", "database.connections.mysql.database").Return("goravel").Once()
	mockConfig.On("GetString", "database.connections.mysql.username").Return("goravel").Once()
	mockConfig.On("GetString", "database.connections.mysql.password").Return("goravel").Once()
	s.mockApp.On("MakeConfig").Return(mockConfig).Once()

	database, err := s.docker.Database()
	s.Nil(err)
	s.NotNil(database)
	databaseImpl := database.(*Database)
	s.Equal("mysql", databaseImpl.connection)

	mockConfig = &configmocks.Config{}
	mockConfig.On("GetString", "database.connections.postgres.driver").Return("postgres").Once()
	mockConfig.On("GetString", "database.connections.postgres.database").Return("goravel").Once()
	mockConfig.On("GetString", "database.connections.postgres.username").Return("goravel").Once()
	mockConfig.On("GetString", "database.connections.postgres.password").Return("goravel").Once()
	s.mockApp.On("MakeConfig").Return(mockConfig).Once()

	database, err = s.docker.Database("postgres")
	s.Nil(err)
	s.NotNil(database)
	databaseImpl = database.(*Database)
	s.Equal("postgres", databaseImpl.connection)
}
