package main

import (
	"context"
	"log"
	"net/http"

	albumv1 "example/gen/album/v1"
	"example/gen/album/v1/albumv1connect"

	"connectrpc.com/connect"
)

func main() {
	client := albumv1connect.NewAlbumServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(),
	)
	resAlbum, errGet2 := GetAlbum(client, "2")
	if errGet2 != nil {
		log.Println(errGet2)
	} else {
		log.Println(resAlbum)
	}

	resAlbums, _ := GetAlbums(client)
	for _, response := range resAlbums {
		log.Println(response.Id, response.Artist, response.Price, response.Title)
	}

	resAdd, errAdd := AddAlbum(client, &albumv1.Album{
		Id: "999",
		Title: "New title in the city",
		Artist: "New artist in the city",
		Price: 12.34,
	})
	if errAdd != nil {
		log.Println(errAdd)
		return
	}
	log.Println(resAdd)

	resAlbum999, errGet999 := GetAlbum(client, "999")
	if errGet999 != nil {
		log.Println(errGet999)
	} else {
		log.Println(resAlbum999)
	}

	resDelete, errDelete := Delete(client, "999")
	if errDelete != nil {
		log.Println(errDelete)
		return
	}
	log.Println(resDelete)

	resAlbums2, _ := GetAlbums(client)
	for _, response := range resAlbums2 {
		log.Println(response.Id, response.Artist, response.Price, response.Title)
	}
}

func GetAlbum(client albumv1connect.AlbumServiceClient, id string) (*albumv1.Album, error) {
	res, err := client.Get(
		context.Background(),
		connect.NewRequest(&albumv1.GetRequest{Id: id}),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("== Get Id %s ==", id)
	return res.Msg.Album, nil
}

func GetAlbums(client albumv1connect.AlbumServiceClient) (map[string]*albumv1.Album, error) {
	res, err := client.GetAll(
		context.Background(),
		connect.NewRequest(&albumv1.GetAllRequest{}),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("== Get All ==")
	return res.Msg.Albums, nil
}

func AddAlbum(client albumv1connect.AlbumServiceClient, album *albumv1.Album) (string, error) {
	log.Printf("== Add Album %v ==", album);
	res, err := client.Add(
		context.Background(),
		connect.NewRequest(&albumv1.AddRequest{
			Album: album,
		}),
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Msg.Id, nil
}

func Delete(client albumv1connect.AlbumServiceClient, id string) (string, error) {
	log.Printf("== Delete Id %s ==", id);
	res, err := client.Delete(
		context.Background(),
		connect.NewRequest(&albumv1.DeleteRequest{Id: id}),
	)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.Msg.Id, nil
}