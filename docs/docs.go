// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/notebook": {
            "post": {
                "security": [
                    {
                        "user_token": []
                    }
                ],
                "description": "create a notebook",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notebook"
                ],
                "summary": "create a notebook",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/router.NotebookCreateInput"
                            }
                        }
                    }
                }
            }
        },
        "/notebook/list": {
            "get": {
                "security": [
                    {
                        "user_token": []
                    }
                ],
                "description": "get notebook list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "notebook"
                ],
                "summary": "get notebook list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Notebook"
                            }
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "login user info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.UserLoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/router.AuthOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Note": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdTime": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "encrypted": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "notebook": {
                    "$ref": "#/definitions/model.Notebook"
                },
                "notebookId": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "userId": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                },
                "versionCode": {
                    "type": "integer"
                },
                "versionKey": {
                    "type": "string"
                }
            }
        },
        "model.Notebook": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "notebook": {
                    "$ref": "#/definitions/model.Notebook"
                },
                "notes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Note"
                    }
                },
                "password": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                },
                "userId": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "optId": {
                    "type": "integer"
                },
                "salt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "router.AuthOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "router.NotebookCreateInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "router.UserLoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "opt": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "user_token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
