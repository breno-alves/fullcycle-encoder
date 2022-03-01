package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"log"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	conn, err := db.DB()

	if err != nil {
		log.Fatalf("error on db initialization on test. error: %v", err)
	}
	defer conn.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
