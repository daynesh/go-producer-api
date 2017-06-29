package requests

import (
	"encoding/json"
)

/* Payload should contain the following structure
   {
	   body: {...},                    // raw payload data as JSON
	   header: {
		   timestamp: 1497406310,      // in Unix epoch time
		   topic: "sometopic"          // pub-sub topic
	   }
   }
*/

type header struct {
	Timestamp int64  `json:"timestamp" binding:"required"`
	Topic     string `json:"topic" binding:"required"` // Topic name
}

// PublisherRequest specifies parameters for /publish requests
type PublisherRequest struct {
	Header header           `json:"header" binding:"required"`
	Body   *json.RawMessage `json:"body" binding:"required"`
}
