package rpc

import "context"

type SweepAllRequest struct {
	// Destination public address.
	Address string `json:"address"`

	// Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`

	// (Optional) Sweep from this set of subaddresses in the account.
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`

	// (Optional) Use outputs in all subaddresses within an account.
	SubaddrIndicesAll bool `json:"subaddr_indices_all,omitempty"`

	// (Optional) Priority for sending the sweep transfer, partially determines fee.
	Priority Priority `json:"priority,omitempty"`

	// Specify the number of separate outputs of smaller denomination that will be created by sweep operation.
	Outputs uint64 `json:"outputs,omitempty"`

	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size,omitempty"`

	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// (Optional, defaults to a random ID) 16 characters hex encoded.
	PaymentId string `json:"payment_id,omitempty"`

	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`

	// (Optional) Include outputs below this amount.
	BelowAmount uint64 `json:"below_amount,omitempty"`

	// (Optional) If true, Do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`

	// (Optional) return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`

	// (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

type SweepAllResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`

	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`

	// The amount transferred for every transaction.
	AmountList []int `json:"amount_list"`

	// The amount of fees paid for every transaction.
	FeeList []int `json:"fee_list"`

	// Metric used for adjusting fee.
	WeightList []int `json:"weight_list"`

	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`

	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`

	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`

	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`

	SpentKeyImagesList []KeyImages `json:"spent_key_images_list"`
}

// Send all unlocked balance to an address.
func (c *Client) SweepAll(ctx context.Context, req *SweepAllRequest) (*SweepAllResponse, error) {
	resp := &SweepAllResponse{}
	err := c.Do(ctx, "sweep_all", &req, resp)
	return resp, err
}
