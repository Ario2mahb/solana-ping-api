package main

import "strings"

type ErrorIdentifier struct {
	Text  PingResultError
	Key   []string
	Short string
}

// transaction responses errors
var (
	BlockhashNotFoundText                  = `rpc response error: {"code":-32002,"message":"Transaction simulation failed: Blockhash not found","data":{"accounts":null,"err":"BlockhashNotFound","logs":[],"unitsConsumed":0}}`
	TransactionHasAlreadyBeenProcessedText = `rpc response error: {"code":-32002,"message":"Transaction simulation failed: This transaction has already been processed","data":{"accounts":null,"err":"AlreadyProcessed","logs":[],"unitsConsumed":0}}`
	RPCServerDeadlineExceededText          = `rpc: call error, err: failed to do request, err: Post "https://api.internal.mainnet-beta.solana.com": context deadline exceeded`
	ServiceUnavilable503Text               = `rpc: call error, err: get status code: 503, body: <html><body><h1>503 Service Unavailable</h1>
	No server is available to handle this request.
	</body></html>`
	TooManyRequest429Text = `rpc: call error, err: get status code: 429, body: <html><head><title>429 Too Many Requests</title></head>
	<body><center><h1>429 Too Many Requests</h1></center><hr><center>nginx/1.21.5</center></body></html>`
	NumSlotsBehindText    = `{count:5 : rpc response error: {"code":-32005,"message":"Node is behind by 153 slots","data":{"numSlotsBehind":153}}`
	RPCEOFText            = `rpc: call error, err: failed to do request, err: Post "https://api.internal.mainnet-beta.solana.com": EOF, body: `
	GatewayTimeout504Text = `rpc: call error, err: get status code: 504, body: <html><body><h1>504 Gateway Time-out</h1>
	The server didn't respond in time.
	</body></html>
	`
)

var (
	BlockhashNotFound = ErrorIdentifier{
		Text:  PingResultError(BlockhashNotFoundText),
		Key:   []string{"BlockhashNotFound"},
		Short: "BlockhashNotFound"}
	TransactionHasAlreadyBeenProcessed = ErrorIdentifier{
		Text:  PingResultError(TransactionHasAlreadyBeenProcessedText),
		Key:   []string{"AlreadyProcessed"},
		Short: "transaction has already been processed"}
	RPCServerDeadlineExceeded = ErrorIdentifier{
		Text:  PingResultError(RPCServerDeadlineExceededText),
		Key:   []string{"context deadline exceeded"},
		Short: "post to rpc server response timeout"}
	ServiceUnavilable503 = ErrorIdentifier{
		Text:  PingResultError(ServiceUnavilable503Text),
		Key:   []string{"code: 503"},
		Short: "503-service-unavailable"}
	TooManyRequest429 = ErrorIdentifier{
		Text:  PingResultError(TooManyRequest429Text),
		Key:   []string{"code: 429"},
		Short: "429-too-many-requests"}
	NumSlotsBehind = ErrorIdentifier{
		Text:  PingResultError(NumSlotsBehindText),
		Key:   []string{"numSlotsBehind"},
		Short: "numSlotsBehind"}
	RPCEOF = ErrorIdentifier{
		Text:  PingResultError(RPCEOFText),
		Key:   []string{"EOF"},
		Short: "rpc error EOF"}
	GatewayTimeout504 = ErrorIdentifier{
		Text:  PingResultError(GatewayTimeout504Text),
		Key:   []string{"code: 504"},
		Short: "504-gateway-timeout"}
)

func (e ErrorIdentifier) IsIdentical(p PingResultError) bool {
	for _, k := range e.Key {
		if strings.Contains(string(p), k) {
			return true
		}
	}
	return false
}