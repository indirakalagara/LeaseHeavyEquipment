package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "TRADELANE",
        "version": "The version number of the current contract"
    },
    "event": {
        "DistanceTravelled": 123.456,
        "EquipProvider": "transport entity currently in possession of asset",
        "LoadCarried": 123.456,
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "extension": {},
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "threshold": 123.456,
        "timestamp": "2017-04-11T22:21:36.536200553+05:30"
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "DistanceTravelled": 123.456,
        "EquipProvider": "transport entity currently in possession of asset",
        "LoadCarried": 123.456,
        "alerts": {
            "active": [
                "OVERLOAD"
            ],
            "cleared": [
                "OVERLOAD"
            ],
            "raised": [
                "OVERLOAD"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "extension": {},
        "inCompliance": true,
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "threshold": 123.456,
        "timestamp": "2017-04-11T22:21:36.537967293+05:30"
    }
}`
