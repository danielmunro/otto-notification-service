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
	notificationService := service.CreateNotificationService()
	userService := service.CreateUserService()
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
			readUser(userService, data.Value)
		} else if *data.TopicPartition.Topic == "images" {
			updateUserImage(userService, data.Value)
		} else if *data.TopicPartition.Topic == "follows" {
			userFollowed(notificationService, data.Value)
		}
	}
}

func userFollowed(notificationService *service.NotificationService, data []byte) {
	log.Print("consuming user followed message :: ", string(data))
	result := decodeToMap(data)
	userData := result["user"].(map[string]interface{})
	userUuid := uuid.MustParse(userData["uuid"].(string))
	followingData := result["following"].(map[string]interface{})
	followingUuid := uuid.MustParse(followingData["uuid"].(string))
	notificationService.CreateFollowNotification(userUuid, followingUuid)
}

func updateUserImage(userService *service.UserService, data []byte) {
	result := decodeToMap(data)
	user := result["user"].(map[string]interface{})
	userUuidStr := user["uuid"].(string)
	s3Key := result["s3_key"].(string)
	userUuid, err := uuid.Parse(userUuidStr)
	if err != nil {
		log.Print(err)
		return
	}
	userService.UpdateProfilePic(userUuid, s3Key)
}

func readUser(userService *service.UserService, data []byte) {
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
	userService.UpsertUser(userModel)
}

func decodeToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}
