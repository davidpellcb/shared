package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation/v2"
)

func main() {

	stringAppId := os.Getenv("APP_ID")
	appId, err := strconv.Atoi(stringAppId)
	if err != nil {
		panic(err)
	}
	stringInstallationId := os.Getenv("INSTALLATION_ID")
	installationId, err := strconv.Atoi(stringInstallationId)
	if err != nil {
		panic(err)
	}

	pemFile, err := os.Create("/tmp/private-key.pem")
	if err != nil {
		panic(err)
	}
	defer pemFile.Close()

	keyFromEnv := os.Getenv("PRIVATE_KEY")
	pemFile.Write([]byte(keyFromEnv))
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, int64(appId), int64(installationId), "/tmp/private-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	token, err := itr.Token(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}
