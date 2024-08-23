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
		{Username: "JaneDoe", Email: "jane.doe@example.com", Name: "Jane Doe", Avatar: "/static/images/avatar.svg", Bio: "Enthusiastic learner", DateJoined: time.Now(), Password: password},
		{Username: "JohnSmith", Email: "john.smith@example.com", Name: "John Smith", Avatar: "/static/images/avatar.svg", Bio: "Loves coding", DateJoined: time.Now(), Password: password},
		{Username: "AliceW", Email: "alice.w@example.com", Name: "Alice W", Avatar: "/static/images/avatar.svg", Bio: "Avid reader", DateJoined: time.Now(), Password: password},
	}
	for _, user := range users {
		err := db.FirstOrCreate(&user, domain.User{Email: user.Email}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func createTopics(db *gorm.DB) error {
	topics := []domain.Topic{
		{Name: "Go Programming"},
		{Name: "Java Programming"},
		{Name: "Data Science"},
		{Name: "Mobile Development"},
		{Name: "Microservices Architecture"},
		{Name: "Go Best Practices"},
		{Name: "Web Development"},
		{Name: "Machine Learning"},
		{Name: "Cybersecurity"},
		{Name: "Cloud Computing"},
	}
	for _, topic := range topics {
		err := db.FirstOrCreate(&topic, topic).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func createRooms(db *gorm.DB) error {
	rooms := []domain.Room{
		{Name: "GoLang Beginners Workshop", HostID: 1, TopicID: 1, Description: "An introductory session for new Go developers."},
		{Name: "Advanced Java Programming", HostID: 2, TopicID: 2, Description: "Deep dive into advanced concepts in Java programming."},
		{Name: "Data Science with Python", HostID: 1, TopicID: 3, Description: "Exploring data science techniques using Python libraries."},
		{Name: "React Native for Mobile Development", HostID: 2, TopicID: 4, Description: "Building mobile apps with React Native."},
		{Name: "Microservices Architecture", HostID: 1, TopicID: 5, Description: "Designing and implementing microservices."},
		{Name: "Go Best Practices", HostID: 3, TopicID: 6, Description: "Discussing best practices in Go development."},
		{Name: "Modern Web Development", HostID: 1, TopicID: 7, Description: "Techniques and tools for building modern web applications."},
		{Name: "Machine Learning 101", HostID: 2, TopicID: 8, Description: "Introduction to machine learning concepts and algorithms."},
		{Name: "Introduction to Cybersecurity", HostID: 3, TopicID: 9, Description: "Basic concepts and best practices in cybersecurity."},
		{Name: "Cloud Computing Essentials", HostID: 2, TopicID: 10, Description: "Exploring key concepts in cloud computing."},
	}
	for _, room := range rooms {
		err := db.FirstOrCreate(&room, room).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func createRoomParticipants(db *gorm.DB) error {
	roomParticipants := []domain.RoomParticipant{
		{RoomID: 1, UserID: 1},
		{RoomID: 1, UserID: 3},
		{RoomID: 2, UserID: 1}, 
		{RoomID: 2, UserID: 2},
		{RoomID: 3, UserID: 1},
		{RoomID: 3, UserID: 2},
		{RoomID: 4, UserID: 2}, 
		{RoomID: 4, UserID: 3},
		{RoomID: 5, UserID: 1}, 
		{RoomID: 6, UserID: 3},
		{RoomID: 7, UserID: 1}, 
		{RoomID: 8, UserID: 2}, 
		{RoomID: 8, UserID: 3},
		{RoomID: 9, UserID: 1}, 
		{RoomID: 9, UserID: 3},
		{RoomID: 10, UserID: 2},
	}
	for _, rp := range roomParticipants {
		err := db.FirstOrCreate(&rp, rp).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func createMessages(db *gorm.DB) error {
	messages := []domain.Message{
		{RoomID: 1, UserID: 1, Body: "Great session for beginners!"},
		{RoomID: 2, UserID: 2, Body: "Loving the advanced Java concepts."},
		{RoomID: 3, UserID: 1, Body: "Data science is such an exciting field!"},
		{RoomID: 4, UserID: 2, Body: "React Native makes mobile development easier."},
		{RoomID: 5, UserID: 1, Body: "Microservices are the future of scalable architecture."},
		{RoomID: 6, UserID: 3, Body: "These Go best practices are really helpful."},
		{RoomID: 7, UserID: 2, Body: "Modern web development frameworks are powerful."},
		{RoomID: 8, UserID: 3, Body: "Machine learning is revolutionizing industries."},
		{RoomID: 9, UserID: 1, Body: "Understanding cybersecurity is so important."},
		{RoomID: 10, UserID: 2, Body: "Cloud computing enables flexible infrastructure."},
		{RoomID: 1, UserID: 3, Body: "I found this workshop very insightful."},
		{RoomID: 2, UserID: 1, Body: "Can we cover more on Java security features?"},
		{RoomID: 3, UserID: 2, Body: "What are the best Python libraries for data science?"},
		{RoomID: 4, UserID: 3, Body: "Let's build a sample app with React Native."},
	}
	for _, message := range messages {
		err := db.FirstOrCreate(&message, message).Error
		if err != nil {
			return err
		}
	}
	return nil
}
