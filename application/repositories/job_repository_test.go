package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db: db}
	videoRepo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	jobRepository := repositories.JobRepositoryDb{Db: db}
	savedJob, err := jobRepository.Insert(job)

	require.Nil(t, err)
	require.NotEmpty(t, savedJob.ID)

	jobFromDb, err := jobRepository.Find(savedJob.ID)

	require.Nil(t, err)
	require.NotEmpty(t, jobFromDb.ID)
	require.Equal(t, savedJob.ID, jobFromDb.ID)
	require.Equal(t, jobFromDb.Video.ID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	videoRepo := repositories.VideoRepositoryDb{Db: db}
	videoRepo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	jobRepository := repositories.JobRepositoryDb{Db: db}
	savedJob, err := jobRepository.Insert(job)

	require.Nil(t, err)
	require.NotEmpty(t, savedJob.ID)

	job.Status = "Complete"
	jobRepository.Update(job)

	jobFromDb, err := jobRepository.Find(savedJob.ID)

	require.Nil(t, err)
	require.Equal(t, jobFromDb.Status, job.Status)
}
