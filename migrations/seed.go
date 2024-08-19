package migrations

import (
	"time"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := createUsers(db); err != nil {
		return err
	}
	if err := createTopics(db); err != nil {
		return err
	}
	if err := createRooms(db); err != nil {
		return err
	}
	if err := createRoomParticipants(db); err != nil {
		return err
	}
	if err := createMessages(db); err != nil {
		return err
	}
	return nil
}

func createUsers(db *gorm.DB) error {
	password, _ := bcrypt.HashPassword("test123")
	users := []domain.User{
		{ID: 1, Username: "Elyar 1", Email: "elyar@email.com", Avatar: "/static/images/avatar.svg", Name: "Elyar 1", Bio: "This is some test bio 1", DateJoined: time.Now(), Password: password},
		{ID: 2, Username: "Elyar 2", Email: "elyar2@email.com", Avatar: "/static/images/avatar.svg", Name: "Elyar 2", Bio: "This is some test bio 2", DateJoined: time.Now(), Password: password},
		{ID: 3, Username: "JaneDoe", Email: "jane.doe@example.com", Name: "Jane Doe", Avatar: "/static/images/avatar.svg", Bio: "Enthusiastic learner", DateJoined: time.Now(), Password: password},
		{ID: 4, Username: "JohnSmith", Email: "john.smith@example.com", Name: "John Smith", Avatar: "/static/images/avatar.svg", Bio: "Loves coding", DateJoined: time.Now(), Password: password},
		{ID: 5, Username: "AliceW", Email: "alice.w@example.com", Name: "Alice W", Avatar: "/static/images/avatar.svg", Bio: "Avid reader", DateJoined: time.Now(), Password: password},
	}
	return db.CreateInBatches(users, len(users)).Error
}

func createTopics(db *gorm.DB) error {
	topics := []domain.Topic{
		{ID: 1, Name: "Topic Test 1"},
		{ID: 2, Name: "Topic Test 2"},
		{ID: 3, Name: "Topic Test 3"},
		{ID: 4, Name: "Topic Test 4"},
		{ID: 5, Name: "Topic Test 5"},
		{ID: 6, Name: "Go Programming"},
		{ID: 7, Name: "Web Development"},
		{ID: 8, Name: "Machine Learning"},
		{ID: 9, Name: "Cybersecurity"},
		{ID: 10, Name: "Cloud Computing"},
	}
	return db.CreateInBatches(topics, len(topics)).Error
}

func createRooms(db *gorm.DB) error {
	rooms := []domain.Room{
		{ID: 1, Name: "Room 1", HostID: 1, TopicID: 1, Description: "This is Room 1"},
		{ID: 2, Name: "Room 2", HostID: 2, TopicID: 1, Description: "This is Room 2"},
		{ID: 3, Name: "Room 3", HostID: 1, TopicID: 2, Description: "This is Room 3"},
		{ID: 4, Name: "Room 4", HostID: 2, TopicID: 2, Description: "This is Room 4"},
		{ID: 5, Name: "Room 5", HostID: 1, TopicID: 3, Description: "This is Room 5"},
		{ID: 6, Name: "Room 6", HostID: 3, TopicID: 4, Description: "Discussing Go best practices"},
		{ID: 7, Name: "Room 7", HostID: 4, TopicID: 5, Description: "Web development techniques"},
		{ID: 8, Name: "Room 8", HostID: 3, TopicID: 6, Description: "Machine learning algorithms"},
		{ID: 9, Name: "Room 9", HostID: 5, TopicID: 7, Description: "Cybersecurity basics"},
		{ID: 10, Name: "Room 10", HostID: 4, TopicID: 8, Description: "Exploring cloud computing"},
	}
	return db.CreateInBatches(rooms, len(rooms)).Error
}

func createRoomParticipants(db *gorm.DB) error {
	roomParticipants := []domain.RoomParticipant{
		{ID: 1, RoomID: 1, UserID: 1},
		{ID: 2, RoomID: 1, UserID: 2},
		{ID: 3, RoomID: 2, UserID: 1},
		{ID: 4, RoomID: 3, UserID: 1},
		{ID: 5, RoomID: 4, UserID: 2},
		{ID: 6, RoomID: 5, UserID: 1},
		{ID: 7, RoomID: 6, UserID: 3},
		{ID: 8, RoomID: 7, UserID: 4},
		{ID: 9, RoomID: 8, UserID: 5},
		{ID: 10, RoomID: 9, UserID: 4},
		{ID: 11, RoomID: 10, UserID: 5},
		{ID: 12, RoomID: 1, UserID: 3},
		{ID: 13, RoomID: 2, UserID: 4},
		{ID: 14, RoomID: 3, UserID: 5},
		{ID: 15, RoomID: 4, UserID: 3},
	}
	return db.CreateInBatches(roomParticipants, len(roomParticipants)).Error
}

func createMessages(db *gorm.DB) error {
	messages := []domain.Message{
		{ID: 1, RoomID: 1, UserID: 1, Body: "Nice room"},
		{ID: 2, RoomID: 2, UserID: 2, Body: "Good one!"},
		{ID: 3, RoomID: 3, UserID: 1, Body: "Interesting topic"},
		{ID: 4, RoomID: 4, UserID: 2, Body: "Let's discuss further"},
		{ID: 5, RoomID: 5, UserID: 1, Body: "Great insights"},
		{ID: 6, RoomID: 6, UserID: 3, Body: "Learning a lot here"},
		{ID: 7, RoomID: 7, UserID: 4, Body: "Web development is fun"},
		{ID: 8, RoomID: 8, UserID: 5, Body: "AI is the future"},
		{ID: 9, RoomID: 9, UserID: 4, Body: "Security is crucial"},
		{ID: 10, RoomID: 10, UserID: 5, Body: "Cloud computing is fascinating"},
		{ID: 11, RoomID: 1, UserID: 3, Body: "I agree with that"},
		{ID: 12, RoomID: 2, UserID: 4, Body: "Let's dive deeper"},
		{ID: 13, RoomID: 3, UserID: 5, Body: "Can we explore this more?"},
		{ID: 14, RoomID: 4, UserID: 3, Body: "Sure, let's go ahead"},
	}
	return db.CreateInBatches(messages, len(messages)).Error
}
