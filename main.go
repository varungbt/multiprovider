package main

import (
	"context"
	"multiprovider/database"
	_ "multiprovider/database/datastore"
)

func main() {
	ctx := context.Background()
	d, _ := database.OpenConnection(ctx, "google")
	d.Client.RegisterJob("Job1", "conveterjob")

}
