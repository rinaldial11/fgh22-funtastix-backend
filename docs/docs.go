// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login authentication",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register account",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 5,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "get all movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "List Movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.MovieHome"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add movie",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Add movie",
                "parameters": [
                    {
                        "type": "string",
                        "name": "author",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "name": "cast_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "duration",
                        "in": "formData"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "name": "genre_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "release_date",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "synopsis",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "add image",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "add banner",
                        "name": "banner",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/models.Movie"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "description": "Get movie details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Movie Details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "movie id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/models.Movie"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Delete movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "movie id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/models.Movie"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/profiles": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get current logged in profile info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Profile Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/models.Profile"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Edit current logged in profile",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Edit profile",
                "parameters": [
                    {
                        "type": "string",
                        "name": "confirm_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "minLength": 3,
                        "type": "string",
                        "name": "first_name",
                        "in": "formData"
                    },
                    {
                        "minLength": 3,
                        "type": "string",
                        "name": "last_name",
                        "in": "formData"
                    },
                    {
                        "minLength": 6,
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "phone_number",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "point",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "profile user",
                        "name": "picture",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "results": {
                                            "$ref": "#/definitions/models.Profile"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Movie": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "example": "Anthony Russo, Joe Russo"
                },
                "banner": {
                    "type": "string",
                    "example": "example/avengers.jpg"
                },
                "casts": {},
                "duration": {
                    "type": "string",
                    "example": "03:02:00"
                },
                "genre": {},
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "image": {
                    "type": "string",
                    "example": "example/avengers.jpg"
                },
                "releaseDate": {
                    "type": "string",
                    "example": "2018-12-12"
                },
                "synopsis": {
                    "type": "string",
                    "example": "The Avengers assemble to reverse the damage caused by Thanos in Avengers: Infinity War."
                },
                "title": {
                    "type": "string",
                    "example": "Avengers : Endgame"
                },
                "uploadedBy": {}
            }
        },
        "models.MovieHome": {
            "type": "object",
            "properties": {
                "genre": {},
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "image": {
                    "type": "string",
                    "example": "example/avengers.jpg"
                },
                "title": {
                    "type": "string",
                    "example": "Avengers : Endgame"
                }
            }
        },
        "models.Profile": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string",
                    "example": "Budiono"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lastName": {
                    "type": "string",
                    "example": "Siregar"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "08516839587"
                },
                "picture": {
                    "type": "string",
                    "example": "03f91853-f686-4190-a854-06f32dc17da7.jpeg"
                },
                "point": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "pageInfo": {},
                "results": {},
                "success": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "172.16.211.131:8888",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Funtastix",
	Description:      "Funtastix backend-app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}