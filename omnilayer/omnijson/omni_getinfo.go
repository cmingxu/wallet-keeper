package omnijson

/*
Result
{
  "omnicoreversion_int" : xxxxxxx,      // (number) client version as integer
  "omnicoreversion" : "x.x.x.x-xxx",    // (string) client version
  "mastercoreversion" : "x.x.x.x-xxx",  // (string) client version (DEPRECIATED)
  "bitcoincoreversion" : "x.x.x",       // (string) Bitcoin Core version
  "commitinfo" : "xxxxxxx",             // (string) build commit identifier
  "block" : nnnnnn,                     // (number) index of the last processed block
  "blocktime" : nnnnnnnnnn,             // (number) timestamp of the last processed block
  "blocktransactions" : nnnn,           // (number) Omni transactions found in the last processed block
  "totaltransactions" : nnnnnnnn,       // (number) Omni transactions processed in total
  "alerts" : [                          // (array of JSON objects) active protocol alert (if any)
    {
      "alerttype" : n                       // (number) alert type as integer
      "alerttype" : "xxx"                   // "alertexpiringbyblock", "alertexpiringbyblocktime", "alertexpiringbyclientversion" or "error"
      "alertexpiry" : "nnnnnnnnnn"          // (string) expiration criteria (can refer to block height, timestamp or client verion)
      "alertmessage" : "xxx"                // (string) information about the alert
    },
    ...
  ]
}
*/

type OmniGetInfoResult struct {
	VersionInt         int32  `json:"omnicoreversion_int"`
	Version            string `json:"omnicoreversion"`
	BitcoinCoreVersion string `json:"bitcoincoreversion"`
	CommitInfo         string `json:"commitinfo"`
	Block              int32  `json:"block"`
	BlockTimestamp     int32  `json:"blocktime"`
	BlockTransaction   int32  `json:"blocktransactions"`
	TotalTransaction   int32  `json:"totaltransactions"`
}

type OmniGetInfoCommand struct{}

func (OmniGetInfoCommand) Method() string {
	return "omni_getinfo"
}

func (OmniGetInfoCommand) ID() string {
	return "1"
}

func (OmniGetInfoCommand) Params() []interface{} {
	return nil
}
