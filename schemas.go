package main

var schemas = `
{
    "API": {
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
                            "Status": {
                                "description": "Status of the transport entity ",
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
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset. Argument is a JSON encoded string containing only an assetID.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an assetID for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssets": {
            "description": "Delete the state of all assets. No arguments are accepted. For each managed asset, the state and history are erased, and the asset is removed if necessary from recent states.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAllAssets function",
                    "enum": [
                        "deleteAllAssets"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAsset": {
            "description": "Delete one or more properties from an asset. Argument is a JSON encoded string containing an AssetID and an array of qualified property names. An example would be {'assetID':'A1',['event.common.EquipProvider', 'event.customer.LoadCarried']} and the result of that invoke would be the removal of the EquipProvider field and the LoadCarried field with a recalculation of the alert and compliance status.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested assetID with a list or qualified property names.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "qualPropsToDelete": {
                                "items": {
                                    "description": "The qualified name of a property. E.g. 'event.common.EquipProvider', 'event.custom.LoadCarried', etc.",
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "required": [
                            "assetID",
                            "qualPropsToDelete"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deletePropertiesFromAsset function",
                    "enum": [
                        "deletePropertiesFromAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. AssetID is required along with one or more writable properties. Establishes the next asset state. ",
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
                            "Status": {
                                "description": "Status of the transport entity ",
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
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setCreateOnUpdate": {
            "description": "Allow updateAsset to redirect to createAsset when assetID does not exist.",
            "properties": {
                "args": {
                    "description": "True for redirect allowed, false for error on asset does not exist.",
                    "items": {
                        "setCreateOnUpdate": {
                            "type": "boolean"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "setCreateOnUpdate function",
                    "enum": [
                        "setCreateOnUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level in the contract.",
            "properties": {
                "args": {
                    "description": "logging levels indicate what you see",
                    "items": {
                        "logLevel": {
                            "enum": [
                                "CRITICAL",
                                "ERROR",
                                "WARNING",
                                "NOTICE",
                                "INFO",
                                "DEBUG"
                            ],
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "setLoggingLevel function",
                    "enum": [
                        "setLoggingLevel"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "nickname": {
                                "default": "LHE",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. AssetID is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an assetID for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                      "assetID": {
                          "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                      "DistanceTravelled": {
                          "description": "Distance travelled by the asset.",
                          "type": "number"
                      },
                      "EquipProvider": {
                          "description": "transport entity currently in possession of asset",
                          "type": "string"
                      },
                      "Status": {
                          "description": "Status of the transport entity ",
                          "type": "string"
                      },
                      "LoadCarried": {
                          "description": "Load Carried by the asset in KGS.",
                          "type": "number"
                      },
                      "timestamp": {
                          "description": "RFC3339nanos formatted timestamp.",
                          "type": "string"
                      },
                      "threshold": {
                            "description": "Load threshold inclusive in KGS.",
                            "type": "number"
                      }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetSamples": {
            "description": "Returns a string generated from the schema containing sample Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSamples function",
                    "enum": [
                        "readAssetSamples"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected sample data",
                    "type": "string"
                }
            },
            "type": "object"
        },
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
        }
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
              "DistanceTravelled": {
                  "description": "Distance travelled by the asset.",
                  "type": "number"
              },
              "EquipProvider": {
                  "description": "transport entity currently in possession of asset",
                  "type": "string"
              },
              "Status": {
                  "description": "Status of the transport entity ",
                  "type": "string"
              },
              "LoadCarried": {
                  "description": "Load Carried by the asset in KGS.",
                  "type": "number"
              },
              "timestamp": {
                  "description": "RFC3339nanos formatted timestamp.",
                  "type": "string"
              },
              "threshold": {
                    "description": "Load threshold inclusive in KGS.",
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
                    "default": "LHE",
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
                "timestamp": {
                    "description": "RFC3339nanos formatted timestamp.",
                    "type": "string"
                },
                "threshold": {
                      "description": "Load threshold inclusive in KGS.",
                      "type": "number"
                }
            },
            "type": "object"
        }
    }
}`
