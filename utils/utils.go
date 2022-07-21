package utils

import (
	"encoding/json"
	"example/fiber/docs"
	"fmt"
	"os"
)

func PrintJson(v any) {
	if s, e := json.MarshalIndent(v, "", "  "); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(string(s))
	}
}

func SetSwaggerInfos() {
	docs.SwaggerInfo.Title = os.Getenv("APP_NAME")
	docs.SwaggerInfo.Description = os.Getenv("APP_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
