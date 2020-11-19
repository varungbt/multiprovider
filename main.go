package main

import (
	"context"
	"fmt"
	"multiprovider/database"
	_ "multiprovider/database/cosmosdb"
	_ "multiprovider/database/datastore"
)

func main() {
	ctx := context.Background()

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("Registering a processing job in Google Cloud Datastore")
	d, _ := database.OpenConnection(ctx, "google")
	d.Client.RegisterJob("ConverterJobinGCP", "conveterjob")

	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("Registering a processing job in Azure CosmosDb")
	c, _ := database.OpenConnection(ctx, "azure")
	c.Client.RegisterJob("ConverterJobinAzure", "conveterjob")
	fmt.Println("--------------------------------------------------------------------------")

}
