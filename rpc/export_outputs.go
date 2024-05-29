package rpc

import "context"

type ExportOutputsRequest struct {
	// (Optional) If true, export all outputs. Otherwise, export outputs since the last export. (Defaults to false)
	All bool `json:"all,omitempty"`
}

type ExportOutputsResponse struct {
	// Wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// Export all outputs in hex format.
func (c *Client) ExportOutputs(ctx context.Context, req *ExportOutputsRequest) (*ExportOutputsResponse, error) {
	resp := &ExportOutputsResponse{}
	err := c.Do(ctx, "export_outputs", &req, resp)
	return resp, err
}
