// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "zeroidentidad",
            "url": "https://zeroidentidad.github.io/chat"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/cart": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Agregar producto al carrito",
                "parameters": [
                    {
                        "description": "Cart data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.Cart"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Cart"
                        }
                    }
                }
            }
        },
        "/api/cart/{userid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Listar productos en carrito",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id usuario",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.Cart"
                            }
                        }
                    }
                }
            }
        },
        "/api/cart/{userid}/item/{productid}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Eliminar producto en carrito",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id usuario",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Id producto",
                        "name": "productid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.message"
                        }
                    }
                }
            }
        },
        "/api/product": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Upsert: Agregar/modificar producto si recibe id en body",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Product"
                        }
                    }
                }
            }
        },
        "/api/product/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Obtener un producto",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id producto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Product"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Eliminar producto",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id Producto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.message"
                        }
                    }
                }
            }
        },
        "/api/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Listar productos con parametros opcionales de: orden por precio y buscar por nombre",
                "parameters": [
                    {
                        "enum": [
                            "ASC",
                            "DESC"
                        ],
                        "type": "string",
                        "description": "order options",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "string",
                        "description": "string name",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.Product"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.Cart": {
            "type": "object",
            "properties": {
                "productid": {
                    "type": "string"
                },
                "userid": {
                    "type": "string"
                }
            }
        },
        "http.Product": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "imageurl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "http.message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Test ecommerce RestAPI",
	Description:      "Este es un servidor rest api de prueba.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
