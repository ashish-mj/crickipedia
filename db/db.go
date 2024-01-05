package db

import (
	"fmt"
	"log"
	"github.com/couchbase/gocb/v2"
	"crickipedia/models"
	"time"
)

var (
    collection *gocb.Collection
    clster *gocb.Cluster
	bucket_name string
)

func InitDb(connectionString,bucketName,username,password,scope,collectionName string) error {

	cluster, err := gocb.Connect("couchbase://"+connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Println("Connection Failed", err)
		return err
	}

	bucket := cluster.Bucket(bucketName)
	col := bucket.Scope(scope).Collection(collectionName) 

	collection = col
	clster = cluster
	bucket_name = bucketName
	return nil
}

func GetDocument(docID string) (models.Player, error) {
	var player models.Player
	response, err := collection.Get(docID, nil)
	if err != nil {
		return player, err
	}
	response.Content(&player)
	return player, nil
}

func GetAllDocuments() ([]models.Player, error){
	var player[] models.Player
	queryResult, err := clster.Query(
		fmt.Sprintf("SELECT x.* FROM "+bucket_name+ " x"),
		&gocb.QueryOptions{Adhoc: true},
	)
	if err != nil {
		log.Println(err)
		return player,err
       
	}
	for queryResult.Next() {
		var ply models.Player
		queryResult.Row(&ply)
		player = append(player,ply) 
}
    return player, nil
}

func InsertDocument(player models.Player) error{
	key := player.Id
	_, err := collection.Insert(key, &player, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteDocument(key string) error{
	_ , err := collection.Remove(key, &gocb.RemoveOptions{
		Timeout:         1000 * time.Millisecond,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateDocument(player models.Player) error{
	key := player.Id
	_ , err := collection.Upsert(key, &player, nil)
	if err != nil {
		return err
	}
	return nil
}
