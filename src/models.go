package main

import (
    "encoding/json"
)

type Header struct {
    Timestamp int64     `json:"timestamp" binding:"required"`
    Type      string    `json:"type" binding:"required"`        // Topic name
    DeviceId  string    `json:"deviceId" binding:"required"`
    UserId    string    `json:"userId" binding:"required"`
}

/* Payload should contain the following structure
    {
        body: {...},                    // raw payload data as JSON
        header: {
            timestamp: 1497406310,      // in Unix epoch time
            type: "sometopic",          // pub-sub topic
            deviceId: "asdf-zxcv-qwer"  // unique identifier for device
            userId: "0123456789"        // if signed in (optional)
        }
    }
*/
type PublisherPayload struct {
    Header  Header           `json:"header" binding:"required"`
    Body    *json.RawMessage `json:"body" binding:"required"`
}