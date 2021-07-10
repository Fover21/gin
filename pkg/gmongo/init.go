package gmongo

import (
	"context"
	"time"

	"gin_one/pkg/setting"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mgo *mongo.Database
)

func Setup() {
	Mgo = NewClient()
}
func NewClient() (Cli *mongo.Database) {
	/*op:=options.ClientOptions{uri: setting.MongoSetting.ApplyURI,
		MaxPoolSize: &(setting.MongoSetting.MaxPoolSize),
	}*/

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(setting.MongoSetting.ApplyURI))
	if err != nil {
		return nil
	}
	// ping5次，都不通认为失败
	for i := 0; i < 5; i++ {
		err := client.Ping(context.TODO(), nil)
		if err == nil {
			return client.Database("bigdata")
		}
		time.Sleep(100 * time.Microsecond)

	}
	return nil
}
