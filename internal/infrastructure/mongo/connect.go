package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(_ context.Context, uri, user, pwd string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	if user != "" && pwd != "" {
		opts.SetAuth(options.Credential{
			Username: user,
			Password: pwd,
		})
	}

	return mongo.Connect(context.TODO(), opts)
}
