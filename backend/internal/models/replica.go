package models

type Replica struct {
    Name      string `json:"name"`
    Current   bool   `json:"current"`
    NodeName  string `json:"nodeName"`
    Status    string `json:"status"`
    StartTime string `json:"startTime"`
}
