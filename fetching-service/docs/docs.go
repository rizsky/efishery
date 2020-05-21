// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-21 21:18:37.346091 +0700 WIB m=+0.068700513

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/admin/list": {
            "get": {
                "description": "get data by param, data graph (mean, median, min, max) below \"products\" please notice it senpai",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fetch"
                ],
                "summary": "GET data by param and return it with filter and graphical data by Price",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWRtaW4gRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjEyIiwicm9sZSI6ImFkbWluIiwidGltZXN0YW1wIjoiMjAyMC0wNS0yMVQxNDoxMzoyOS4wNDlaIiwiaWF0IjoxNTkwMDcwNDI4fQ.qOoaipBvBQdJOCWxcHeTLjBFhMHg1EEtRHsdzX7dHVk",
                        "description": "Authentication header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "SUMATERA",
                        "description": "area_provinsi",
                        "name": "area_provinsi",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "2019-11-10T17:00:00.00Z",
                        "description": "dari_tanggal",
                        "name": "dari_tanggal",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "2020-08-30T13:35:00Z",
                        "description": "sampai_tanggal",
                        "name": "sampai_tanggal",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DataProductsWithGraph"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/list": {
            "get": {
                "description": "GET product data and convert price to idr",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fetch"
                ],
                "summary": "GET product data and convert price to idr",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjMyMCIsInJvbGUiOiJ1c2VyIiwidGltZXN0YW1wIjoiMjAyMC0wNS0xOVQwNDo1MzowNS40NDNaIiwiaWF0IjoxNTg5ODY0MDE0fQ.pnQXvaa9vFlIM4AKnptrbzbfztEGPXWU-ISDH-K1Vsg",
                        "description": "Authentication header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DataProducts"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/private-claim": {
            "get": {
                "description": "Get Private claim from JWT with data phone, role, name, timestamp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "GET Private claim from JWT",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiRXhjZWxlbmN5IiwicGhvbmUiOiIrNjI4MjI3MjQyMjMyMCIsInJvbGUiOiJ1c2VyIiwidGltZXN0YW1wIjoiMjAyMC0wNS0xOVQwNDo1MzowNS40NDNaIiwiaWF0IjoxNTg5ODY0MDE0fQ.pnQXvaa9vFlIM4AKnptrbzbfztEGPXWU-ISDH-K1Vsg",
                        "description": "Authentication header",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "oke",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DataProducts": {
            "type": "object",
            "properties": {
                "area_kota": {
                    "type": "string"
                },
                "area_provinsi": {
                    "type": "string"
                },
                "komoditas": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "price_usd": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                },
                "tgl_parsed": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "models.DataProductsWithGraph": {
            "type": "object",
            "properties": {
                "grafik_data_price": {
                    "type": "object",
                    "$ref": "#/definitions/models.GraphData"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DataProducts"
                    }
                }
            }
        },
        "models.GraphData": {
            "type": "object",
            "properties": {
                "maksimum": {
                    "type": "string"
                },
                "median": {
                    "type": "string"
                },
                "minimum": {
                    "type": "string"
                },
                "rata-rata": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
