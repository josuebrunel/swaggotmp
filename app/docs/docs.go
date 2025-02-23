// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "demo.com",
        "contact": {
            "name": "API Support",
            "url": "http://demo.com/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/organization": {
            "get": {
                "description": "List organizations",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "List organizations",
                "operationId": "orgs-get",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an organization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "Create an organization",
                "operationId": "org-create",
                "parameters": [
                    {
                        "description": "Organization data",
                        "name": "organization",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RequestOrgCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        },
        "/organization/{org}/tag": {
            "get": {
                "description": "List tag tags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "List tag tags",
                "operationId": "tags-get",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "Create an tag",
                "operationId": "tag-create",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "tag data",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.PayloadTag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        },
        "/organization/{org}/tag/{tag}": {
            "get": {
                "description": "Get an tag",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "Get an tag",
                "operationId": "tag-get",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "tag ID",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete organization tag",
                "tags": [
                    "tag"
                ],
                "summary": "Delete organization tag",
                "operationId": "tag-delete",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "tag ID",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an organization tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "Update an organization tag",
                "operationId": "tag-update",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "tag ID",
                        "name": "tag",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "tag data",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.PayloadTag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        },
        "/organization/{org}/user": {
            "get": {
                "description": "List user users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "List user users",
                "operationId": "users-get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create an user",
                "operationId": "user-create",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RequestUserCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        },
        "/organization/{org}/user/{uuid}": {
            "get": {
                "description": "Get an user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get an user",
                "operationId": "user-get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete organization user",
                "tags": [
                    "user"
                ],
                "summary": "Delete organization user",
                "operationId": "user-delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an organization user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update an organization user",
                "operationId": "user-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "organization ID",
                        "name": "org",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RequestUserUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        },
        "/organization/{uuid}": {
            "get": {
                "description": "Get an organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "Get an organization",
                "operationId": "org-get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an organization",
                "tags": [
                    "organization"
                ],
                "summary": "Delete an organization",
                "operationId": "org-delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an organization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "organization"
                ],
                "summary": "Update an organization",
                "operationId": "org-update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Organization data",
                        "name": "organization",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RequestOrgUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        },
        "/user/types": {
            "get": {
                "description": "List users' type",
                "tags": [
                    "user"
                ],
                "summary": "List users's type",
                "operationId": "user-types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/service.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.PayloadTag": {
            "type": "object",
            "required": [
                "name",
                "type"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "service.RequestOrgCreate": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "service.RequestOrgUpdate": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orgParam": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "service.RequestUserCreate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "birth_place": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "service.RequestUserUpdate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "birth_place": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "orgParam": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "userParam": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "service.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Ekolo Swagger UI",
	Description:      "Ekolo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
