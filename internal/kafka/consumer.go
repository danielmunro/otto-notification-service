package kafka

import (
	"encoding/json"
	"github.com/danielmunro/otto-notification-service/internal/model"
	"github.com/danielmunro/otto-notification-service/internal/service"
	"github.com/google/uuid"
	"log"
)

func InitializeAndRunLoop() {
	err := loopKafkaReader()
	if err != nil {
		log.Fatal(err)
	}
}

func loopKafkaReader() error {
	reader := GetReader()
	consumerService := service.CreateConsumerService()
	for {
		log.Print("listening for kafka messages")
		data, err := reader.ReadMessage(-1)
		if err != nil {
			log.Print(err)
			return nil
		}
		log.Print("message received on topic :: ", data.TopicPartition.String())
		log.Print("data :: ", string(data.Value))
		if *data.TopicPartition.Topic == "users" {
			readUser(consumerService, data.Value)
		} else if *data.TopicPartition.Topic == "images" {
			updateUserImage(consumerService, data.Value)
		} else if *data.TopicPartition.Topic == "follows" {
			userFollowed(consumerService, data.Value)
		} else if *data.TopicPartition.Topic == "posts" {
			readPost(consumerService, data.Value)
		} else if *data.TopicPartition.Topic == "postLikes" {
			readPostLikes(consumerService, data.Value)
		}
	}
}

func readPostLikes(consumerService *service.ConsumerService, data []byte) {
	postLikeModel, err := model.DecodeMessageToPostLike(data)
	if err != nil {
		log.Print("error reading post like topic :: ", err)
		return
	}
	log.Print("create post like notification :: ", postLikeModel)
	consumerService.CreatePostLikeNotification(postLikeModel)
}

func readPost(consumerService *service.ConsumerService, data []byte) {
	postModel, err := model.DecodeMessageToPost(data)
	if err != nil {
		log.Print("error reading post kafka topic :: ", err)
		return
	}
	log.Print("upsert post :: ", postModel.Uuid)
	consumerService.UpsertPost(postModel)
}

func userFollowed(consumerService *service.ConsumerService, data []byte) {
	log.Print("consuming user followed message :: ", string(data))
	followModel, err := model.DecodeMessageToFollow(data)
	if err != nil {
		log.Print("error decoding follow :: ", err)
		return
	}
	consumerService.UpsertFollow(followModel)
}

func updateUserImage(consumerService *service.ConsumerService, data []byte) {
	result := decodeToMap(data)
	user := result["user"].(map[string]interface{})
	userUuidStr := user["uuid"].(string)
	s3Key := result["s3_key"].(string)
	userUuid, err := uuid.Parse(userUuidStr)
	if err != nil {
		log.Print(err)
		return
	}
	consumerService.UpdateProfilePic(userUuid, s3Key)
}

func readUser(consumerService *service.ConsumerService, data []byte) {
	log.Print("consuming user message ", string(data))
	userModel, err := model.DecodeMessageToUser(data)
	if err != nil {
		log.Print("error decoding message to user error :: ", err)
		return
	}
	_, err = uuid.Parse(userModel.Uuid)
	if err != nil {
		log.Print(err)
		return
	}
	consumerService.UpsertUser(userModel)
}

func decodeToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}
