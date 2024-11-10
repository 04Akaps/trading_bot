package mongoDB

import (
	"context"
	"github.com/04Akaps/trading_bot.git/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	cfg config.Config

	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDB(cfg config.Config) MongoDB {
	ctx := context.Background()

	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.Uri)); err != nil {
		panic(err)
	} else if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		db := client.Database(cfg.MongoDB.DB)

		return MongoDB{
			cfg:    cfg,
			client: client,
			db:     db,
		}
	}

}

func (m MongoDB) ScanTokenList() map[string]bool {
	// TODO query

	return map[string]bool{
		//"ETHBTC": true,
		//"BNBBTC":  true,
		"POLUSDT": true,
	}
}
