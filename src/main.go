package main

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    // "go.mongodb.org/mongo-driver/mongo/findopt"
)

// func insertBsonD(col *mongo.Collection) error {
//     bsonD := bson.D{
//         {"str1", "abc"},
//         {"num1", 1},
//         {"str2", "xyz"},
//         {"num2", bson.A{2, 3, 4}},
//         {"subdoc", bson.D{{"str", "subdoc"}, {"num", 987}}},
//         {"date", time.Now()},
//     }
//     _, err := col.InsertOne(context.Background(), bsonD)
//     return err
// }

func insertBsonM(col *mongo.Collection) error {
    bsonM := bson.M{
        "str1": "efg",
        "num1": 11,
        "str2": "opq",
        "num2": bson.A{12, 13, 14},
        "subdoc": bson.M{"str": "subdoc", "num": 987},
        "date": time.Now(),
    }
    for i := 0; i < 10; i++ {
        _, err := col.InsertOne(context.Background(), bsonM)
        if err != nil {
            return err
        }
    }
    return nil
}

// type myType struct {
//     Str1 string
//     Num1 int
//     Str2 string
//     Num2 []int
//     Subdoc struct {
//         Str string
//         Num int
//     }
//     Date time.Time
// }
//
// func insertStruct(col *mongo.Collection) error {
//     doc := myType{
//         "hij",
//         21,
//         "rst",
//         []int{22, 23, 24},
//         struct {
//             Str string
//             Num int
//         }{"subdoc", 987},
//         time.Now(),
//     }
//     _, err := col.InsertOne(context.Background(), doc)
//     return err
// }

// func findMaxNum(col *mongo.Collection) error {
//     // https://github.com/mongodb/mongo-go-driver/blob/5fea1444e52844a15513c0d9490327b2bd89ed7c/mongo/crud_spec_test.go#L364
//     var opts []findopt.Find
//     opts = append(opts, findopt.Sort("num1".(map[string]interface{})))
//     opts = append(opts, findopt.Limit(int64(1.(float64))))
//     cursor, err := col.Find(context.Background(), filter, opts...)
// }

func mainMain() error {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/?connect=direct"))
    if err != nil {
        return err
    }
    if err = client.Connect(context.Background()); err != nil {
        return err
    }
    defer client.Disconnect(context.Background())

    col := client.Database("test").Collection("col")
    // if err = insertBsonD(col); err != nil {
    //     return err
    // }
    if err = insertBsonM(col); err != nil {
        return err
    }
    // if err = insertStruct(col); err != nil {
    //     return err
    // }
    // if err = findMaxNum(col); err != nil {
    //     return err
    // }
    return nil
}

func main() {
    if err := mainMain(); err != nil {
        log.Fatal(err)
    }
    log.Println("normal end.")
}
