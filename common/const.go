package common

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "client"
)

type Requester interface {
	GetUserID() int
	GetUserEmail() string
	GetUserRole() string
}

const (
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDisLikeRestaurant = "TopicUserDisLikeRestaurant"
)