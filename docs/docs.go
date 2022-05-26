// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/repositories": {
            "get": {
                "description": "get repositories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "List Repositories",
                "operationId": "listRepositories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RepositoryCollectionResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create a repository",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "repositories"
                ],
                "summary": "Create Repository",
                "operationId": "createRepository",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CreateRepository"
                        }
                    },
                    {
                        "type": "string",
                        "description": "organization id",
                        "name": "org_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account number",
                        "name": "account_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    }
                }
            }
        },
        "/repositories/:uuid": {
            "delete": {
                "tags": [
                    "repositories"
                ],
                "summary": "Delete a repository",
                "operationId": "deleteRepository",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CreateRepository": {
            "type": "object",
            "properties": {
                "distribution_arch": {
                    "description": "Architecture to restrict client usage to",
                    "type": "string",
                    "example": "x86_64"
                },
                "distribution_version": {
                    "description": "Version to restrict client usage to",
                    "type": "string",
                    "example": "7"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "description": "URL of the remote yum repository",
                    "type": "string"
                }
            }
        },
        "api.Links": {
            "type": "object",
            "properties": {
                "first": {
                    "description": "Path to first page of results",
                    "type": "string"
                },
                "last": {
                    "description": "Path to last page of results",
                    "type": "string"
                },
                "next": {
                    "description": "Path to next page of results",
                    "type": "string"
                },
                "prev": {
                    "description": "Path to previous page of results",
                    "type": "string"
                }
            }
        },
        "api.Repository": {
            "type": "object",
            "properties": {
                "account_id": {
                    "description": "Account ID of the owner",
                    "type": "string",
                    "readOnly": true
                },
                "distribution_arch": {
                    "description": "Architecture to restrict client usage to",
                    "type": "string",
                    "example": "x86_64"
                },
                "distribution_version": {
                    "description": "Version to restrict client usage to",
                    "type": "string",
                    "example": "7"
                },
                "name": {
                    "type": "string"
                },
                "org_id": {
                    "description": "Organization ID of the owner",
                    "type": "string",
                    "readOnly": true
                },
                "url": {
                    "description": "URL of the remote yum repository",
                    "type": "string"
                },
                "uuid": {
                    "type": "string",
                    "readOnly": true
                }
            }
        },
        "api.RepositoryCollectionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Requested Data",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Repository"
                    }
                },
                "links": {
                    "description": "Links to other pages of results",
                    "$ref": "#/definitions/api.Links"
                },
                "meta": {
                    "description": "Metadata about the request",
                    "$ref": "#/definitions/api.ResponseMetadata"
                }
            }
        },
        "api.ResponseMetadata": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "Total count of results",
                    "type": "integer"
                },
                "limit": {
                    "description": "Limit of results used for the request",
                    "type": "integer"
                },
                "offset": {
                    "description": "Offset into results used for the request",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "RhIdentity": {
            "type": "apiKey",
            "name": "x-rh-identity",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0.0",
	Host:             "api.example.com",
	BasePath:         "/api/content_sources/v1.0/",
	Schemes:          []string{},
	Title:            "ContentSourcesBackend",
	Description:      "API of the Content Sources application on [console.redhat.com](https://console.redhat.com)\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}