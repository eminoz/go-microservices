package redisContoller

import (
	"encoding/json"
	"fmt"
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type User struct {
	Id      string
	Name    string
	Surname string
	Email   string
}

func (i User) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
func (c *Client) SetUser(ctx *gin.Context, id string, user *model.User) *redis.StatusCmd {
	u := User{
		Id:      id,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}
	json.Marshal(user)
	set := c.client.Set(ctx, u.Id, u, 0)
	fmt.Println(set)
	return set
}
func (c *Client) GetUser(ctx *gin.Context, id string) *redis.StringCmd {
	result := c.client.Get(ctx, id)

	return result
}
