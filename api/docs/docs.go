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
        "/api/employees": {
            "get": {
                "description": "Get all employees",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Get all employees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Employee"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create a new employee",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Create a new employee",
                "parameters": [
                    {
                        "description": "Employee object",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Employee"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "409": {
                        "description": "Conflict"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/employees/{id}": {
            "get": {
                "description": "Get an employee by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Get an employee by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Update an existing employee by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Update an existing employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateEmployee object",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateEmployee"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "409": {
                        "description": "Conflict"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete an employee by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Delete an employee by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/games": {
            "get": {
                "description": "Get all games",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Get all games",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Game"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create a new game",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Create a new game",
                "parameters": [
                    {
                        "description": "Game object",
                        "name": "game",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Game"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Game"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "409": {
                        "description": "Conflict"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/games/{id}": {
            "get": {
                "description": "Get an game by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Get an game by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Game"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "Delete an game by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Delete an game by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "Employee": {
            "description": "Employee account information",
            "type": "object",
            "required": [
                "elo",
                "email",
                "name",
                "password"
            ],
            "properties": {
                "elo": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "games": {
                    "description": "Assuming Employee1ID is the foreign key",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Game"
                    }
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "Game": {
            "description": "Game information between an employee and an user",
            "type": "object",
            "required": [
                "employee1ID",
                "employee2ID",
                "winnerFloat",
                "winnerUuid"
            ],
            "properties": {
                "employee1": {
                    "$ref": "#/definitions/Employee"
                },
                "employee1ID": {
                    "type": "string"
                },
                "employee2": {
                    "$ref": "#/definitions/Employee"
                },
                "employee2ID": {
                    "type": "string"
                },
                "winnerFloat": {
                    "type": "integer"
                },
                "winnerUuid": {
                    "type": "string"
                }
            }
        },
        "LoginRequest": {
            "description": "Login request with email and password",
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "UpdateEmployee": {
            "description": "UpdateEmployee account information",
            "type": "object",
            "properties": {
                "elo": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "TIME-SYS-API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}