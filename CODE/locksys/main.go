package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// SmartLockClient holds the base URL for the smart lock API
type SmartLockClient struct {
    BaseURL string
}

// NewSmartLockClient creates a new SmartLockClient
func NewSmartLockClient(baseURL string) *SmartLockClient {
    return &SmartLockClient{BaseURL: baseURL}
}

// Lock sends a request to lock the door
func (c *SmartLockClient) Lock() error {
    return c.sendPostRequest("/lock", nil)
}

// Unlock sends a request to unlock the door
func (c *SmartLockClient) Unlock() error {
    return c.sendPostRequest("/unlock", nil)
}

// Status retrieves the current lock status
func (c *SmartLockClient) Status() (string, error) {
    resp, err := http.Get(c.BaseURL + "/status")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

// sendPostRequest is a helper function to send POST requests
func (c *SmartLockClient) sendPostRequest(path string, data interface{}) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }
    _, err = http.Post(c.BaseURL+path, "application/json", bytes.NewBuffer(jsonData))
    return err
}

func main() {
    client := NewSmartLockClient("http://localhost:8080")

    // Example operations
    err := client.Lock()
    if err != nil {
        fmt.Printf("Error locking: %v\n", err)
        return
    }
    fmt.Println("Locked")

    status, err := client.Status()
    if err != nil {
        fmt.Printf("Error getting status: %v\n", err)
        return
    }
    fmt.Printf("Status: %s\n", status)

    err = client.Unlock()
    if err != nil {
        fmt.Printf("Error unlocking: %v\n", err)
        return
    }
    fmt.Println("Unlocked")
}
