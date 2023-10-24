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
        "/create-user": {
            "post": {
                "description": "Create a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "create user request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/list-item": {
            "post": {
                "description": "List items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "List items",
                "parameters": [
                    {
                        "description": "list item request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.listItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Item"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/list-users": {
            "get": {
                "description": "List users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "description": "list users request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.listUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "login user request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.loginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "api.createUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.listItemRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "offset": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "api.listUsersRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "offset": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "api.loginUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.Item": {
            "type": "object",
            "properties": {
                "deleted": {
                    "type": "boolean"
                },
                "description": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "holder": {
                    "type": "integer"
                },
                "identifier_code": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer"
                },
                "modification_time": {
                    "type": "string"
                },
                "modifier": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "db.User": {
            "type": "object",
            "properties": {
                "activated": {
                    "type": "boolean"
                },
                "description": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "password": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "sql.NullString": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}