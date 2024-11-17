package mongoDB

import (
	"context"
	"github.com/04Akaps/trading_bot.git/config"
	"github.com/04Akaps/trading_bot.git/types"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	cfg config.Config

	client *mongo.Client
	db     *mongo.Database

	volume *mongo.Collection
	scan   *mongo.Collection
}

var defaultScanRes = map[string]bool{
	"POLUSDT": true,
}

func NewMongoDB(cfg config.Config) MongoDB {
	ctx := context.Background()

	log.Println("Try Connect MongoDB")

	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.Uri)); err != nil {
		panic(err)
	} else if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	} else {
		db := client.Database(cfg.MongoDB.DB)

		volume := db.Collection("volume")

		defer log.Println("Success To connect mongoDB")

		return MongoDB{
			cfg:    cfg,
			client: client,
			db:     db,
			volume: volume,
		}
	}

}

func (m MongoDB) ScanTokenList() map[string]bool {

	ctx := context.Background()

	cursor, err := m.scan.Find(ctx, bson.M{})

	if err != nil {
		return defaultScanRes
	}

	length := cursor.RemainingBatchLength()
	res := make(map[string]bool, length)

	for cursor.Next(ctx) {
		var dec struct {
			Symbol string `bson:"symbol"`
		}

		if err = cursor.Decode(&dec); err != nil {
			return defaultScanRes
		}

		res[dec.Symbol] = true
	}

	return res
}

func (m MongoDB) UpdateBulk(models []mongo.WriteModel) {

	_, err := m.volume.BulkWrite(context.Background(), models)

	if err != nil {
		log.Println("Failed to set volume", "err", err)
	}
}

func (m MongoDB) GetVolumeInfo(symbol string) (avgVolume, currentVolume, diff float64) {

	filter := bson.M{"symbol": symbol}
	ctx := context.Background()

	cursor, err := m.volume.Find(ctx, filter)

	if err != nil {
		log.Println("Failed to get volume", "symbol", symbol, "err", err)
		return
	}

	total := cursor.RemainingBatchLength()

	var totalVolume decimal.Decimal

	for cursor.Next(ctx) {
		var res types.VolumeDocument

		if err = cursor.Decode(&res); err != nil {
			log.Println("Failed to decode volume document", "err", err)
		} else {

			v, _ := decimal.NewFromString(res.Volume)
			totalVolume = totalVolume.Add(v)

			currentVolume, _ = v.Round(2).Float64()
		}
	}

	if total > 0 {
		avgVolumeDecimal := totalVolume.Div(decimal.NewFromInt(int64(total)))
		avgVolume, _ = avgVolumeDecimal.Round(2).Float64()

		diffDecimal := decimal.NewFromFloat(currentVolume).Sub(avgVolumeDecimal)
		diffDecimal = diffDecimal.Div(avgVolumeDecimal).Mul(decimal.NewFromInt(100))

		diff, _ = diffDecimal.Round(2).Float64()
	}

	return avgVolume, currentVolume, diff
}
