package main

import (
	"github.com/artemkaxboy/go-hocon"
	"log"
)

type properties struct {
	Greeting string
	LogLevel string `hocon:"path=logLevel,default=debug"`

	AdvertURL     string `hocon:"path=advert.url"`
	AdvertEnabled bool   `hocon:"path=advert.enabled,default=true"`

	Add struct {
		Add1 int32 `hocon:"node=first"`
		Add2 int32 `hocon:"node=second"`
		Sum  int32 `hocon:"node=sum"`
	} `hocon:"node=numbers"`

	Multi struct {
		Multi1  int64 `hocon:"path=numbers.first"`
		Multi2  int64 `hocon:"path=numbers.second"`
		Product int64 `hocon:"path=numbers.product"`
	}

	Div struct {
		Div1 float32 `hocon:"node=first"`
		Div2 float32 `hocon:"path=numbers.second"`
		Quot float32 `hocon:"node=quotient"`
	} `hocon:"path=numbers"`
}

func main() {
	var props properties
	err := hocon.LoadConfigFile("hocon.conf", &props)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(props.Greeting)
	log.Printf("log level is %s", props.LogLevel)

	isSum := props.Add.Add1+props.Add.Add2 == props.Add.Sum
	log.Printf("%d + %d = %d is %t", props.Add.Add1, props.Add.Add2, props.Add.Sum, isSum)

	isProduct := props.Multi.Multi1+props.Multi.Multi2 == props.Multi.Product
	log.Printf("%d * %d = %d is %t", props.Multi.Multi1, props.Multi.Multi2, props.Multi.Product, isProduct)

	isDivision := props.Div.Div1/props.Div.Div2 == props.Div.Quot
	log.Printf("%d / %d = %d is %t", props.Add.Add1, props.Add.Add2, props.Add.Sum, isDivision)

	if props.AdvertEnabled {
		log.Printf("visit %s for more details ...", props.AdvertURL)
	}
}
