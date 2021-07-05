package database

import "github.com/alidevjimmy/go-echo-train/domains"

var (
	Products []domains.Product = []domains.Product{
		{Id: 1, Title: "book1", Price: 24000},
		{Id: 2, Title: "laptop", Price: 900},
		{Id: 3, Title: "cpu", Price: 4000},
		{Id: 4, Title: "ram", Price: 243000},
	}
)
