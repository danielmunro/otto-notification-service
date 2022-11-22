package service

import (
	"github.com/danielmunro/otto-notification-service/internal/db"
	"github.com/danielmunro/otto-notification-service/internal/mapper"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/danielmunro/otto-notification-service/internal/repository"
	"github.com/google/uuid"
	"log"
)

type PostService struct {
	userRepository *repository.UserRepository
	postRepository *repository.PostRepository
}

func CreatePostService() *PostService {
	conn := db.CreateDefaultConnection()
	return &PostService{
		repository.CreateUserRepository(conn),
		repository.CreatePostRepository(conn),
	}
}

func (p *PostService) UpsertPost(postModel *model.Post) {
	postEntity, err := p.postRepository.FindOneByUuid(uuid.MustParse(postModel.Uuid))
	if err == nil {
		postEntity.UpdatePostFromModel(postModel)
		p.postRepository.Save(postEntity)
	} else {
		user, err := p.userRepository.FindOneByUuid(uuid.MustParse(postModel.User.Uuid))
		if err != nil {
			log.Print("user not found when upserting post :: ", postModel)
			return
		}
		postEntity = mapper.GetPostEntityFromModel(user.ID, postModel)
		p.postRepository.Create(postEntity)
	}
}
