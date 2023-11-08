package mongodb

import (
	"context"
	"fmt"
	"github.com/prometheus/alertmanager/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/go-kit/log"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/notify"
	"github.com/prometheus/alertmanager/template"
)

type Notifier struct {
	conf    *config.MongoDbConfig
	tmpl    *template.Template
	logger  log.Logger
	retrier *notify.Retrier
}

type Document struct {
	Time      primitive.DateTime `json:"time" bson:"time"`
	AlertData *template.Data
}

func (n Notifier) Notify(ctx context.Context, alert ...*types.Alert) (bool, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", n.conf.Username, n.conf.Password, n.conf.Url, n.conf.Port)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongodb.Connect(context.Background(), clientOptions)
	collection := client.Database(n.conf.Database).Collection(n.conf.Collection)
	data := notify.GetTemplateData(ctx, n.tmpl, alert, n.logger)

	doc := Document{
		Time:      primitive.NewDateTimeFromTime(time.Now()),
		AlertData: data,
	}

	res, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		fmt.Println("InsertOne ERROR:", err)
	}
	if err := client.Disconnect(context.Background()); err != nil {
		// 处理错误
		panic(err)
	}
	//fmt.Println("Inserted document with time:", res.InsertedID)
	n.logger.Log("msg", "Inserted document with time:", "time", res.InsertedID)
	return true, nil
}

// New returns a new notifier that uses the mongodb.
func New(c *config.MongoDbConfig, t *template.Template, l log.Logger) (*Notifier, error) {
	n := &Notifier{
		conf:    c,
		tmpl:    t,
		logger:  l,
		retrier: &notify.Retrier{},
	}

	return n, nil
}
