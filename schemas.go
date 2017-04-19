package main

var schemas = `
{
    "API": {
    "readAssetSchemas": {
        "description": "Returns a string generated from the schema containing APIs and Objects as specified in generate.json in the scripts folder.",
        "properties": {
            "args": {
                "description": "accepts no arguments",
                "items": {},
                "maxItems": 0,
                "minItems": 0,
                "type": "array"
            },
            "function": {
                "description": "readAssetSchemas function",
                "enum": [
                    "readAssetSchemas"
                ],
                "type": "string"
            },
            "method": "query",
            "result": {
                "description": "JSON encoded object containing selected schemas",
                "type": "string"
            }
        },
        "type": "object"
    },
    "createAsset": {
        "description": "Create an asset. One argument, a JSON encoded event. AssetID is required with zero or more writable properties. Establishes an initial asset state.",
        "properties": {
            "args": {
                "description": "args are JSON encoded strings",
                "items": {
                    "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                    "properties": {
                        "DistanceTravelled": {
                            "description": "Distance travelled by the asset.",
                            "type": "number"
                        },
                        "EquipProvider": {
                            "description": "transport entity currently in possession of asset",
                            "type": "string"
                        },
                        "LoadCarried": {
                            "description": "Load Carried by the asset in KGS.",
                            "type": "number"
                        },
                        "assetID": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "extension": {
                            "description": "Application-managed state. Opaque to contract.",
                            "properties": {},
                            "type": "object"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        },
                        "threshold": {
                            "description": "Load threshold inclusive in KGS.",
                            "type": "number"
                        },
                        "timestamp": {
                            "description": "RFC3339nanos formatted timestamp.",
                            "type": "string"
                        }
                    },
                    "required": [
                        "assetID"
                    ],
                    "type": "object"
                },
                "maxItems": 1,
                "minItems": 1,
                "type": "array"
            },
            "function": {
                "description": "createAsset function",
                "enum": [
                    "createAsset"
                ],
                "type": "string"
            }
        },
        "type": "object"
    },
    "objectModelSchemas": {
        "assetIDKey": {
            "description": "An object containing only an assetID for use as an argument to read or delete.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                    "description": "Temperature of the asset in CELSIUS.",
                    "type": "number"
                }
            },
            "required": [
                "assetID"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "SIMPLE",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of fields that constitute the complete asset state.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                    "description": "Temperature of the asset in CELSIUS.",
                    "type": "number"
                }
            },
            "type": "object"
        }
    }
}`
