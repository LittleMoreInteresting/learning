package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func downloadFile(ctx http.Context) error {
	disposition := fmt.Sprintf("attachment; filename=%s", "files.zip")
	ctx.Response().Header().Set("Content-Type", "application/zip")
	ctx.Response().Header().Set("Content-Disposition", disposition)
	ctx.Response().Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	zipWriter := zip.NewWriter(ctx.Response())
	defer zipWriter.Close()
	zipEntry, errZip := zipWriter.Create("08.png")
	if errZip != nil {
		return errZip
	}
	url1 := "http://uxbpubtest.uxiaobang.com/public/avatar/08.png"
	content, err := readRemoteContent(url1)
	if err != nil {
		return err
	}
	_, err = zipEntry.Write(content)
	if err != nil {
		return err
	}
	url2 := "http://www.uxiaobang.com/teacher/static/img/pic_people.2e56c03f.png"
	zipEntry, errZip = zipWriter.Create("747c82f5.png")
	if errZip != nil {
		return errZip
	}
	content2, err := readRemoteContent(url2)
	if err != nil {
		return err
	}
	_, err = zipEntry.Write(content2)
	if err != nil {
		return err
	}
	return nil
}

func readRemoteContent(url string) ([]byte, error) {
	response, err := stdhttp.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

func main() {
	var opts = []http.ServerOption{
		http.Address(":8001"),
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}

	httpSrv := http.NewServer(
		opts...,
	)
	route := httpSrv.Route("/")
	route.GET("/download", downloadFile)

	app := kratos.New(
		kratos.Name("download"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
