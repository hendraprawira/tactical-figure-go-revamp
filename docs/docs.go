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
        "/figure/multi": {
            "get": {
                "description": "Get All Tactical Multi Line Figure",
                "tags": [
                    "Tactical Figure"
                ],
                "summary": "Get All Tactical Multi Line Figure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/tacticalfigure.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/tacticalfigure.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/figure/point": {
            "get": {
                "description": "Get All Tactical Point Figure",
                "tags": [
                    "Tactical Figure"
                ],
                "summary": "Get All Tactical Point Figure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/tacticalfigure.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/tacticalfigure.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/figure/single": {
            "get": {
                "description": "Get All Tactical Single Line Figure",
                "tags": [
                    "Tactical Figure"
                ],
                "summary": "Get All Tactical Multi Single Line Figure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Point"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/tacticalfigure.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/tacticalfigure.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Point": {
            "type": "object",
            "properties": {
                "altitude": {
                    "type": "string"
                },
                "amplifications": {
                    "type": "string"
                },
                "color": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "opacity": {
                    "type": "integer"
                }
            }
        },
        "tacticalfigure.ErrorResponse": {
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
