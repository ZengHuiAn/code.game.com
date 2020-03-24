package proto


type ServerInfo struct {
	ID      int32  `json:"id"`
	Gateway string  `json:"gateway"`
	Name    string  `json:"name"`
	// Flag    int32   `json:"flag"`
}