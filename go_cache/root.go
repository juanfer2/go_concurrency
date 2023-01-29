package go_cache

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type CacheUser struct {
	User
	ExpireAtTimestamp int64
}

type LocalCache struct{
	
}