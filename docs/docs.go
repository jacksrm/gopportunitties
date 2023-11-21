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
        "/opening": {
            "post": {
                "description": "Creates a new opening",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "openings"
                ],
                "summary": "Create Opening",
                "parameters": [
                    {
                        "description": "Data needed to create an opening",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateOpening"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.RequestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.RequestResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateOpening": {
            "type": "object",
            "required": [
                "company",
                "link",
                "location",
                "remote",
                "role",
                "salary"
            ],
            "properties": {
                "company": {
                    "type": "string",
                    "minLength": 10
                },
                "link": {
                    "type": "string"
                },
                "location": {
                    "type": "string",
                    "minLength": 10
                },
                "remote": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string",
                    "minLength": 10
                },
                "salary": {
                    "type": "integer"
                }
            }
        },
        "dto.RequestResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
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
