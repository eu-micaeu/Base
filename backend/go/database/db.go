package database

import (
    "context"
    "errors"
    "os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
    client *mongo.Client
    db     *mongo.Database
    users  *mongo.Collection // convenience cached collection
}

// New cria uma nova conexão com o banco de dados MongoDB
func New(ctx context.Context) (*DB, error) {
    uri := getenv("MONGO_URI", "mongodb://mongo:27017")
    dbName := getenv("MONGO_DB", "base")
    usersColl := getenv("MONGO_USERS_COLLECTION", "users")

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }
    database := client.Database(dbName)
    return &DB{
        client: client,
        db:     database,
        users:  database.Collection(usersColl),
    }, nil
}

// Close fecha a conexão com o banco de dados
func (db *DB) Close(ctx context.Context) error {
    if db.client == nil {
        return nil
    }
    return db.client.Disconnect(ctx)
}

// Generic CRUD helpers (non-generic method signatures using interface{})
func (db *DB) InsertOne(ctx context.Context, collection string, doc interface{}) (primitive.ObjectID, error) {
    coll := db.db.Collection(collection)
    res, err := coll.InsertOne(ctx, doc)
    if err != nil {
        return primitive.NilObjectID, err
    }
    oid, _ := res.InsertedID.(primitive.ObjectID)
    return oid, nil
}

// FindAll preenche 'results' que deve ser ponteiro para slice
func (db *DB) FindAll(ctx context.Context, collection string, filter interface{}, results interface{}) error {
    coll := db.db.Collection(collection)
    cur, err := coll.Find(ctx, filter)
    if err != nil {
        return err
    }
    defer cur.Close(ctx)
    if err := cur.All(ctx, results); err != nil {
        return err
    }
    return cur.Err()
}

// FindByID decodifica em 'result' (ponteiro para struct) e retorna se encontrado
func (db *DB) FindByID(ctx context.Context, collection, idHex string, result interface{}) (bool, error) {
    oid, err := primitive.ObjectIDFromHex(idHex)
    if err != nil {
        return false, errors.New("invalid object id")
    }
    coll := db.db.Collection(collection)
    err = coll.FindOne(ctx, bson.M{"_id": oid}).Decode(result)
    if err == mongo.ErrNoDocuments {
        return false, nil
    }
    if err != nil {
        return false, err
    }
    return true, nil
}


// getenv retorna o valor da variável de ambiente ou um valor padrão se não estiver definida
func getenv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}
