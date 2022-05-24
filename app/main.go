package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Adress struct {
	Prefecture     string `json:"都道府県"`
	Municipalities string `json:"市区町村"`
	Town           string `json:"町名"`
}

type PostInfo struct {
	PostalCode string `json:"郵便番号"`
	Adress     `json:"住所"`
}

func (p *PostInfo) searchAdress(s string) {
	f, _ := os.Open("./KEN_ALL.csv")
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Printf("IO ERROR")
			break
		}

		if record[2] == s {
			prefecture, municipalities, town := record[6], record[7], record[8]
			if record[8] == "以下に掲載がない場合" {
				town = ""
			}

			p.PostalCode = s
			p.Adress = Adress{prefecture, municipalities, town}
			break
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/search", func(ctx *gin.Context) {
		postalCode := ctx.Query("param")
		var postInfo PostInfo
		postInfo.searchAdress(postalCode)
		ctx.JSON(200, postInfo)
	})

	err := router.Run(":8080")
	if err == nil {
		fmt.Printf("ROUTER RUN ERROR")
	}
}
