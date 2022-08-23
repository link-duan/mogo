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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v2/base/instances": {
            "get": {
                "description": "gets all instances, databases, and table nested data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Gets all instance database and table data filtered by permissions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.RespInstanceSimple"
                            }
                        }
                    }
                }
            }
        },
        "/api/v2/base/shorturls": {
            "post": {
                "description": "Create short links",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Create short links",
                "parameters": [
                    {
                        "description": "params",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ReqShortURLCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            }
        },
        "/api/v2/base/su/{s-code}": {
            "get": {
                "description": "Get short links",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Get short links",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "short code",
                        "name": "s-code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Moved Permanently",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/instances/{instance-id}/table-dependencies": {
            "get": {
                "description": "Result of table dependency resolution",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Result of table dependency resolution",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "instance id",
                        "name": "instance-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "databaseName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "tableName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/core.ResPage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/view.RespTableDependencies"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes-results/{result-id}": {
            "patch": {
                "description": "only support excelProcess update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Updates the action on the execution result",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "result id",
                        "name": "result-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "excelProcess",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes/{node-id}/crontab": {
            "post": {
                "description": "isRetry: 0 no 1 yes\nretryInterval: the unit is in seconds, 100 means 100s",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Creating a scheduled node scheduling task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ReqCreateCrontab"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            },
            "patch": {
                "description": "isRetry: 0 no 1 yes\nretryInterval: the unit is in seconds, 100 means 100s",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Updating a scheduled node scheduling task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ReqUpdateCrontab"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes/{node-id}/lock-acquire": {
            "post": {
                "description": "Force the file edit lock to be acquired",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Force the file edit lock to be acquired",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes/{node-id}/results": {
            "get": {
                "description": "Obtain the node execution result record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Obtain the node execution result record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "current",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "isExcludeCrontabResult",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.RespNodeResultList"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/workers": {
            "get": {
                "description": "The scheduled task list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "The scheduled task list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Current 总记录数",
                        "name": "current",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "nodeName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "PageSize 每页记录数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort 顺序",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "ClickHouse 10; MySQL 11; OfflineSync 20",
                        "name": "tertiary",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Total 总页数",
                        "name": "total",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/core.ResPage"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/view.RespWorkerList"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/workers/dashboard": {
            "get": {
                "description": "Kanban on the execution status of a scheduled task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Kanban on the execution status of a scheduled task",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "isInCharge",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "start",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.RespWorkerDashboard"
                        }
                    }
                }
            }
        },
        "/api/v2/storage": {
            "post": {
                "description": "Creating a log library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storage"
                ],
                "summary": "Creating a log library",
                "parameters": [
                    {
                        "type": "string",
                        "name": "brokers",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "consumers",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "databaseId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "days",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "kafkaSkipBrokenMessages",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "rawLogField",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Raw JSON data",
                        "name": "source",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "tableName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "timeField",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "topics",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "1 string 2 float",
                        "name": "typ",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            }
        },
        "/api/v2/storage/mapping-json": {
            "post": {
                "description": "Kafka JSON field mapping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storage"
                ],
                "summary": "Kafka JSON field mapping",
                "parameters": [
                    {
                        "type": "string",
                        "name": "data",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.MappingStruct"
                        }
                    }
                }
            }
        },
        "/api/v2/storage/{storage-id}": {
            "patch": {
                "description": "Storage update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storage"
                ],
                "summary": "Storage update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "table id",
                        "name": "storage-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "kafkaBrokers",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "min 1 max 8",
                        "name": "kafkaConsumerNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "kafkaSkipBrokenMessages",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "kafkaTopic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "mergeTreeTTL",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Res"
                        }
                    }
                }
            }
        },
        "/api/v2/storage/{storage-id}/analysis-fields": {
            "get": {
                "description": "Storage analysis field list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storage"
                ],
                "summary": "Storage analysis field list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "table id",
                        "name": "storage-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.RespStorageAnalysisFields"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.Pagination": {
            "type": "object",
            "properties": {
                "current": {
                    "description": "Current means current page number",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "PageSize means max item count of a page",
                    "type": "integer"
                },
                "sort": {
                    "description": "Sort means sort expression",
                    "type": "string"
                },
                "total": {
                    "description": "Total means total page count",
                    "type": "integer"
                }
            }
        },
        "core.Res": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code means response business code",
                    "type": "integer"
                },
                "data": {
                    "description": "Data means response data payload"
                },
                "msg": {
                    "description": "Msg means response extra message",
                    "type": "string"
                }
            }
        },
        "core.ResPage": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code means response business code",
                    "type": "integer"
                },
                "data": {
                    "description": "Data means response data payload"
                },
                "msg": {
                    "description": "Msg means response extra message",
                    "type": "string"
                },
                "pagination": {
                    "$ref": "#/definitions/core.Pagination"
                }
            }
        },
        "view.MappingStruct": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.MappingStructItem"
                    }
                }
            }
        },
        "view.MappingStructItem": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "view.ReqCreateCrontab": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.ReqCrontabArg"
                    }
                },
                "cron": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "dutyUid": {
                    "type": "integer"
                },
                "isRetry": {
                    "type": "integer"
                },
                "retryInterval": {
                    "type": "integer"
                },
                "retryTimes": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "view.ReqCrontabArg": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "val": {
                    "type": "string"
                }
            }
        },
        "view.ReqShortURLCreate": {
            "type": "object",
            "properties": {
                "originUrl": {
                    "type": "string"
                }
            }
        },
        "view.ReqUpdateCrontab": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.ReqCrontabArg"
                    }
                },
                "cron": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "dutyUid": {
                    "type": "integer"
                },
                "isRetry": {
                    "type": "integer"
                },
                "retryInterval": {
                    "type": "integer"
                },
                "retryTimes": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "view.RespDatabaseSimple": {
            "type": "object",
            "properties": {
                "cluster": {
                    "type": "string"
                },
                "databaseName": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "iid": {
                    "type": "integer"
                },
                "isCreateByCV": {
                    "type": "integer"
                },
                "tables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespTableSimple"
                    }
                }
            }
        },
        "view.RespInstanceSimple": {
            "type": "object",
            "properties": {
                "databases": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespDatabaseSimple"
                    }
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "instanceName": {
                    "type": "string"
                }
            }
        },
        "view.RespNodeResult": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "cost": {
                    "type": "integer"
                },
                "ctime": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "excelProcess": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "nodeId": {
                    "type": "integer"
                },
                "result": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "uid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "view.RespNodeResultList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespNodeResult"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "view.RespStorageAnalysisFields": {
            "type": "object",
            "properties": {
                "keys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.StorageAnalysisField"
                    }
                }
            }
        },
        "view.RespTableDependencies": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespTableDeps"
                    }
                },
                "utime": {
                    "type": "integer"
                }
            }
        },
        "view.RespTableDeps": {
            "type": "object",
            "properties": {
                "database": {
                    "type": "string"
                },
                "deps": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "engine": {
                    "type": "string"
                },
                "table": {
                    "type": "string"
                },
                "totalBytes": {
                    "type": "integer"
                },
                "totalRows": {
                    "type": "integer"
                }
            }
        },
        "view.RespTableSimple": {
            "type": "object",
            "properties": {
                "createType": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "did": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "tableName": {
                    "type": "string"
                }
            }
        },
        "view.RespUserSimpleInfo": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
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
        "view.RespWorkerDashboard": {
            "type": "object",
            "properties": {
                "flows": {
                    "description": "Execution trend chart",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.WorkerStatsRow"
                    }
                },
                "nodeFailed": {
                    "description": "node status",
                    "type": "integer"
                },
                "nodeSuccess": {
                    "description": "node status",
                    "type": "integer"
                },
                "nodeUnknown": {
                    "description": "node status",
                    "type": "integer"
                },
                "workerFailed": {
                    "description": "Execution status of each periodic schedule",
                    "type": "integer"
                },
                "workerSuccess": {
                    "description": "Execution status of each periodic schedule",
                    "type": "integer"
                },
                "workerUnknown": {
                    "description": "Execution status of each periodic schedule",
                    "type": "integer"
                }
            }
        },
        "view.RespWorkerList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespWorkerRow"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "view.RespWorkerRow": {
            "type": "object",
            "properties": {
                "chargePerson": {
                    "$ref": "#/definitions/view.RespUserSimpleInfo"
                },
                "cost": {
                    "type": "integer"
                },
                "crontab": {
                    "type": "string"
                },
                "endTime": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "iid": {
                    "type": "integer"
                },
                "nodeId": {
                    "type": "integer"
                },
                "nodeName": {
                    "type": "string"
                },
                "startTime": {
                    "type": "integer"
                },
                "status": {
                    "description": "unknown 0; success 1; failed 2",
                    "type": "integer"
                },
                "tertiary": {
                    "type": "integer"
                }
            }
        },
        "view.StorageAnalysisField": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "ctime": {
                    "type": "integer"
                },
                "field": {
                    "type": "string"
                },
                "hashTyp": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "rootName": {
                    "type": "string"
                },
                "tid": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                },
                "utime": {
                    "type": "integer"
                }
            }
        },
        "view.WorkerStatsRow": {
            "type": "object",
            "properties": {
                "failed": {
                    "type": "integer"
                },
                "success": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "integer"
                },
                "unknown": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.4.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ClickVisual API",
	Description:      "Defines interface prefixes in terms of module overrides：\n- base : the global basic readable information module\n- storage : the log module\n- alarm : the alarm module\n- pandas : the data analysis module\n- cmdb : the configuration module\n- sysop : the system management module",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
