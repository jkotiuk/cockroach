// Code generated by gen_batch.go; DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package roachpb

import (
	"bytes"
	"fmt"
	"strconv"
)

// GetInner returns the error contained in the union.
func (ru ErrorDetail) GetInner() error {
	switch t := ru.GetValue().(type) {
	case *ErrorDetail_NotLeaseHolder:
		return t.NotLeaseHolder
	case *ErrorDetail_RangeNotFound:
		return t.RangeNotFound
	case *ErrorDetail_RangeKeyMismatch:
		return t.RangeKeyMismatch
	case *ErrorDetail_ReadWithinUncertaintyInterval:
		return t.ReadWithinUncertaintyInterval
	case *ErrorDetail_TransactionAborted:
		return t.TransactionAborted
	case *ErrorDetail_TransactionPush:
		return t.TransactionPush
	case *ErrorDetail_TransactionRetry:
		return t.TransactionRetry
	case *ErrorDetail_TransactionReplay:
		return t.TransactionReplay
	case *ErrorDetail_TransactionStatus:
		return t.TransactionStatus
	case *ErrorDetail_WriteIntent:
		return t.WriteIntent
	case *ErrorDetail_WriteTooOld:
		return t.WriteTooOld
	case *ErrorDetail_OpRequiresTxn:
		return t.OpRequiresTxn
	case *ErrorDetail_ConditionFailed:
		return t.ConditionFailed
	case *ErrorDetail_LeaseRejected:
		return t.LeaseRejected
	case *ErrorDetail_NodeUnavailable:
		return t.NodeUnavailable
	case *ErrorDetail_Send:
		return t.Send
	case *ErrorDetail_RaftGroupDeleted:
		return t.RaftGroupDeleted
	case *ErrorDetail_ReplicaCorruption:
		return t.ReplicaCorruption
	case *ErrorDetail_ReplicaTooOld:
		return t.ReplicaTooOld
	case *ErrorDetail_AmbiguousResult:
		return t.AmbiguousResult
	case *ErrorDetail_StoreNotFound:
		return t.StoreNotFound
	case *ErrorDetail_HandledRetryableTxnError:
		return t.HandledRetryableTxnError
	case *ErrorDetail_IntegerOverflow:
		return t.IntegerOverflow
	case *ErrorDetail_UnsupportedRequest:
		return t.UnsupportedRequest
	case *ErrorDetail_MixedSuccess:
		return t.MixedSuccess
	case *ErrorDetail_TimestampBefore:
		return t.TimestampBefore
	case *ErrorDetail_TxnAlreadyEncounteredError:
		return t.TxnAlreadyEncounteredError
	case *ErrorDetail_IntentMissing:
		return t.IntentMissing
	case *ErrorDetail_MergeInProgress:
		return t.MergeInProgress
	default:
		return nil
	}
}

// GetInner returns the Request contained in the union.
func (ru RequestUnion) GetInner() Request {
	switch t := ru.GetValue().(type) {
	case *RequestUnion_Get:
		return t.Get
	case *RequestUnion_Put:
		return t.Put
	case *RequestUnion_ConditionalPut:
		return t.ConditionalPut
	case *RequestUnion_Increment:
		return t.Increment
	case *RequestUnion_Delete:
		return t.Delete
	case *RequestUnion_DeleteRange:
		return t.DeleteRange
	case *RequestUnion_ClearRange:
		return t.ClearRange
	case *RequestUnion_Scan:
		return t.Scan
	case *RequestUnion_BeginTransaction:
		return t.BeginTransaction
	case *RequestUnion_EndTransaction:
		return t.EndTransaction
	case *RequestUnion_AdminSplit:
		return t.AdminSplit
	case *RequestUnion_AdminMerge:
		return t.AdminMerge
	case *RequestUnion_AdminTransferLease:
		return t.AdminTransferLease
	case *RequestUnion_AdminChangeReplicas:
		return t.AdminChangeReplicas
	case *RequestUnion_HeartbeatTxn:
		return t.HeartbeatTxn
	case *RequestUnion_Gc:
		return t.Gc
	case *RequestUnion_PushTxn:
		return t.PushTxn
	case *RequestUnion_ResolveIntent:
		return t.ResolveIntent
	case *RequestUnion_ResolveIntentRange:
		return t.ResolveIntentRange
	case *RequestUnion_Merge:
		return t.Merge
	case *RequestUnion_TruncateLog:
		return t.TruncateLog
	case *RequestUnion_RequestLease:
		return t.RequestLease
	case *RequestUnion_ReverseScan:
		return t.ReverseScan
	case *RequestUnion_ComputeChecksum:
		return t.ComputeChecksum
	case *RequestUnion_CheckConsistency:
		return t.CheckConsistency
	case *RequestUnion_InitPut:
		return t.InitPut
	case *RequestUnion_TransferLease:
		return t.TransferLease
	case *RequestUnion_LeaseInfo:
		return t.LeaseInfo
	case *RequestUnion_WriteBatch:
		return t.WriteBatch
	case *RequestUnion_Export:
		return t.Export
	case *RequestUnion_Import:
		return t.Import
	case *RequestUnion_QueryTxn:
		return t.QueryTxn
	case *RequestUnion_QueryIntent:
		return t.QueryIntent
	case *RequestUnion_AdminScatter:
		return t.AdminScatter
	case *RequestUnion_AddSstable:
		return t.AddSstable
	case *RequestUnion_RecomputeStats:
		return t.RecomputeStats
	case *RequestUnion_Refresh:
		return t.Refresh
	case *RequestUnion_RefreshRange:
		return t.RefreshRange
	case *RequestUnion_GetSnapshotForMerge:
		return t.GetSnapshotForMerge
	case *RequestUnion_RangeStats:
		return t.RangeStats
	default:
		return nil
	}
}

// GetInner returns the Response contained in the union.
func (ru ResponseUnion) GetInner() Response {
	switch t := ru.GetValue().(type) {
	case *ResponseUnion_Get:
		return t.Get
	case *ResponseUnion_Put:
		return t.Put
	case *ResponseUnion_ConditionalPut:
		return t.ConditionalPut
	case *ResponseUnion_Increment:
		return t.Increment
	case *ResponseUnion_Delete:
		return t.Delete
	case *ResponseUnion_DeleteRange:
		return t.DeleteRange
	case *ResponseUnion_ClearRange:
		return t.ClearRange
	case *ResponseUnion_Scan:
		return t.Scan
	case *ResponseUnion_BeginTransaction:
		return t.BeginTransaction
	case *ResponseUnion_EndTransaction:
		return t.EndTransaction
	case *ResponseUnion_AdminSplit:
		return t.AdminSplit
	case *ResponseUnion_AdminMerge:
		return t.AdminMerge
	case *ResponseUnion_AdminTransferLease:
		return t.AdminTransferLease
	case *ResponseUnion_AdminChangeReplicas:
		return t.AdminChangeReplicas
	case *ResponseUnion_HeartbeatTxn:
		return t.HeartbeatTxn
	case *ResponseUnion_Gc:
		return t.Gc
	case *ResponseUnion_PushTxn:
		return t.PushTxn
	case *ResponseUnion_ResolveIntent:
		return t.ResolveIntent
	case *ResponseUnion_ResolveIntentRange:
		return t.ResolveIntentRange
	case *ResponseUnion_Merge:
		return t.Merge
	case *ResponseUnion_TruncateLog:
		return t.TruncateLog
	case *ResponseUnion_RequestLease:
		return t.RequestLease
	case *ResponseUnion_ReverseScan:
		return t.ReverseScan
	case *ResponseUnion_ComputeChecksum:
		return t.ComputeChecksum
	case *ResponseUnion_CheckConsistency:
		return t.CheckConsistency
	case *ResponseUnion_InitPut:
		return t.InitPut
	case *ResponseUnion_LeaseInfo:
		return t.LeaseInfo
	case *ResponseUnion_WriteBatch:
		return t.WriteBatch
	case *ResponseUnion_Export:
		return t.Export
	case *ResponseUnion_Import:
		return t.Import
	case *ResponseUnion_QueryTxn:
		return t.QueryTxn
	case *ResponseUnion_QueryIntent:
		return t.QueryIntent
	case *ResponseUnion_AdminScatter:
		return t.AdminScatter
	case *ResponseUnion_AddSstable:
		return t.AddSstable
	case *ResponseUnion_RecomputeStats:
		return t.RecomputeStats
	case *ResponseUnion_Refresh:
		return t.Refresh
	case *ResponseUnion_RefreshRange:
		return t.RefreshRange
	case *ResponseUnion_GetSnapshotForMerge:
		return t.GetSnapshotForMerge
	case *ResponseUnion_RangeStats:
		return t.RangeStats
	default:
		return nil
	}
}

// SetInner sets the error in the union.
func (ru *ErrorDetail) SetInner(r error) bool {
	var union isErrorDetail_Value
	switch t := r.(type) {
	case *NotLeaseHolderError:
		union = &ErrorDetail_NotLeaseHolder{t}
	case *RangeNotFoundError:
		union = &ErrorDetail_RangeNotFound{t}
	case *RangeKeyMismatchError:
		union = &ErrorDetail_RangeKeyMismatch{t}
	case *ReadWithinUncertaintyIntervalError:
		union = &ErrorDetail_ReadWithinUncertaintyInterval{t}
	case *TransactionAbortedError:
		union = &ErrorDetail_TransactionAborted{t}
	case *TransactionPushError:
		union = &ErrorDetail_TransactionPush{t}
	case *TransactionRetryError:
		union = &ErrorDetail_TransactionRetry{t}
	case *TransactionReplayError:
		union = &ErrorDetail_TransactionReplay{t}
	case *TransactionStatusError:
		union = &ErrorDetail_TransactionStatus{t}
	case *WriteIntentError:
		union = &ErrorDetail_WriteIntent{t}
	case *WriteTooOldError:
		union = &ErrorDetail_WriteTooOld{t}
	case *OpRequiresTxnError:
		union = &ErrorDetail_OpRequiresTxn{t}
	case *ConditionFailedError:
		union = &ErrorDetail_ConditionFailed{t}
	case *LeaseRejectedError:
		union = &ErrorDetail_LeaseRejected{t}
	case *NodeUnavailableError:
		union = &ErrorDetail_NodeUnavailable{t}
	case *SendError:
		union = &ErrorDetail_Send{t}
	case *RaftGroupDeletedError:
		union = &ErrorDetail_RaftGroupDeleted{t}
	case *ReplicaCorruptionError:
		union = &ErrorDetail_ReplicaCorruption{t}
	case *ReplicaTooOldError:
		union = &ErrorDetail_ReplicaTooOld{t}
	case *AmbiguousResultError:
		union = &ErrorDetail_AmbiguousResult{t}
	case *StoreNotFoundError:
		union = &ErrorDetail_StoreNotFound{t}
	case *HandledRetryableTxnError:
		union = &ErrorDetail_HandledRetryableTxnError{t}
	case *IntegerOverflowError:
		union = &ErrorDetail_IntegerOverflow{t}
	case *UnsupportedRequestError:
		union = &ErrorDetail_UnsupportedRequest{t}
	case *MixedSuccessError:
		union = &ErrorDetail_MixedSuccess{t}
	case *BatchTimestampBeforeGCError:
		union = &ErrorDetail_TimestampBefore{t}
	case *TxnAlreadyEncounteredErrorError:
		union = &ErrorDetail_TxnAlreadyEncounteredError{t}
	case *IntentMissingError:
		union = &ErrorDetail_IntentMissing{t}
	case *MergeInProgressError:
		union = &ErrorDetail_MergeInProgress{t}
	default:
		return false
	}
	ru.Value = union
	return true
}

// SetInner sets the Request in the union.
func (ru *RequestUnion) SetInner(r Request) bool {
	var union isRequestUnion_Value
	switch t := r.(type) {
	case *GetRequest:
		union = &RequestUnion_Get{t}
	case *PutRequest:
		union = &RequestUnion_Put{t}
	case *ConditionalPutRequest:
		union = &RequestUnion_ConditionalPut{t}
	case *IncrementRequest:
		union = &RequestUnion_Increment{t}
	case *DeleteRequest:
		union = &RequestUnion_Delete{t}
	case *DeleteRangeRequest:
		union = &RequestUnion_DeleteRange{t}
	case *ClearRangeRequest:
		union = &RequestUnion_ClearRange{t}
	case *ScanRequest:
		union = &RequestUnion_Scan{t}
	case *BeginTransactionRequest:
		union = &RequestUnion_BeginTransaction{t}
	case *EndTransactionRequest:
		union = &RequestUnion_EndTransaction{t}
	case *AdminSplitRequest:
		union = &RequestUnion_AdminSplit{t}
	case *AdminMergeRequest:
		union = &RequestUnion_AdminMerge{t}
	case *AdminTransferLeaseRequest:
		union = &RequestUnion_AdminTransferLease{t}
	case *AdminChangeReplicasRequest:
		union = &RequestUnion_AdminChangeReplicas{t}
	case *HeartbeatTxnRequest:
		union = &RequestUnion_HeartbeatTxn{t}
	case *GCRequest:
		union = &RequestUnion_Gc{t}
	case *PushTxnRequest:
		union = &RequestUnion_PushTxn{t}
	case *ResolveIntentRequest:
		union = &RequestUnion_ResolveIntent{t}
	case *ResolveIntentRangeRequest:
		union = &RequestUnion_ResolveIntentRange{t}
	case *MergeRequest:
		union = &RequestUnion_Merge{t}
	case *TruncateLogRequest:
		union = &RequestUnion_TruncateLog{t}
	case *RequestLeaseRequest:
		union = &RequestUnion_RequestLease{t}
	case *ReverseScanRequest:
		union = &RequestUnion_ReverseScan{t}
	case *ComputeChecksumRequest:
		union = &RequestUnion_ComputeChecksum{t}
	case *CheckConsistencyRequest:
		union = &RequestUnion_CheckConsistency{t}
	case *InitPutRequest:
		union = &RequestUnion_InitPut{t}
	case *TransferLeaseRequest:
		union = &RequestUnion_TransferLease{t}
	case *LeaseInfoRequest:
		union = &RequestUnion_LeaseInfo{t}
	case *WriteBatchRequest:
		union = &RequestUnion_WriteBatch{t}
	case *ExportRequest:
		union = &RequestUnion_Export{t}
	case *ImportRequest:
		union = &RequestUnion_Import{t}
	case *QueryTxnRequest:
		union = &RequestUnion_QueryTxn{t}
	case *QueryIntentRequest:
		union = &RequestUnion_QueryIntent{t}
	case *AdminScatterRequest:
		union = &RequestUnion_AdminScatter{t}
	case *AddSSTableRequest:
		union = &RequestUnion_AddSstable{t}
	case *RecomputeStatsRequest:
		union = &RequestUnion_RecomputeStats{t}
	case *RefreshRequest:
		union = &RequestUnion_Refresh{t}
	case *RefreshRangeRequest:
		union = &RequestUnion_RefreshRange{t}
	case *GetSnapshotForMergeRequest:
		union = &RequestUnion_GetSnapshotForMerge{t}
	case *RangeStatsRequest:
		union = &RequestUnion_RangeStats{t}
	default:
		return false
	}
	ru.Value = union
	return true
}

// SetInner sets the Response in the union.
func (ru *ResponseUnion) SetInner(r Response) bool {
	var union isResponseUnion_Value
	switch t := r.(type) {
	case *GetResponse:
		union = &ResponseUnion_Get{t}
	case *PutResponse:
		union = &ResponseUnion_Put{t}
	case *ConditionalPutResponse:
		union = &ResponseUnion_ConditionalPut{t}
	case *IncrementResponse:
		union = &ResponseUnion_Increment{t}
	case *DeleteResponse:
		union = &ResponseUnion_Delete{t}
	case *DeleteRangeResponse:
		union = &ResponseUnion_DeleteRange{t}
	case *ClearRangeResponse:
		union = &ResponseUnion_ClearRange{t}
	case *ScanResponse:
		union = &ResponseUnion_Scan{t}
	case *BeginTransactionResponse:
		union = &ResponseUnion_BeginTransaction{t}
	case *EndTransactionResponse:
		union = &ResponseUnion_EndTransaction{t}
	case *AdminSplitResponse:
		union = &ResponseUnion_AdminSplit{t}
	case *AdminMergeResponse:
		union = &ResponseUnion_AdminMerge{t}
	case *AdminTransferLeaseResponse:
		union = &ResponseUnion_AdminTransferLease{t}
	case *AdminChangeReplicasResponse:
		union = &ResponseUnion_AdminChangeReplicas{t}
	case *HeartbeatTxnResponse:
		union = &ResponseUnion_HeartbeatTxn{t}
	case *GCResponse:
		union = &ResponseUnion_Gc{t}
	case *PushTxnResponse:
		union = &ResponseUnion_PushTxn{t}
	case *ResolveIntentResponse:
		union = &ResponseUnion_ResolveIntent{t}
	case *ResolveIntentRangeResponse:
		union = &ResponseUnion_ResolveIntentRange{t}
	case *MergeResponse:
		union = &ResponseUnion_Merge{t}
	case *TruncateLogResponse:
		union = &ResponseUnion_TruncateLog{t}
	case *RequestLeaseResponse:
		union = &ResponseUnion_RequestLease{t}
	case *ReverseScanResponse:
		union = &ResponseUnion_ReverseScan{t}
	case *ComputeChecksumResponse:
		union = &ResponseUnion_ComputeChecksum{t}
	case *CheckConsistencyResponse:
		union = &ResponseUnion_CheckConsistency{t}
	case *InitPutResponse:
		union = &ResponseUnion_InitPut{t}
	case *LeaseInfoResponse:
		union = &ResponseUnion_LeaseInfo{t}
	case *WriteBatchResponse:
		union = &ResponseUnion_WriteBatch{t}
	case *ExportResponse:
		union = &ResponseUnion_Export{t}
	case *ImportResponse:
		union = &ResponseUnion_Import{t}
	case *QueryTxnResponse:
		union = &ResponseUnion_QueryTxn{t}
	case *QueryIntentResponse:
		union = &ResponseUnion_QueryIntent{t}
	case *AdminScatterResponse:
		union = &ResponseUnion_AdminScatter{t}
	case *AddSSTableResponse:
		union = &ResponseUnion_AddSstable{t}
	case *RecomputeStatsResponse:
		union = &ResponseUnion_RecomputeStats{t}
	case *RefreshResponse:
		union = &ResponseUnion_Refresh{t}
	case *RefreshRangeResponse:
		union = &ResponseUnion_RefreshRange{t}
	case *GetSnapshotForMergeResponse:
		union = &ResponseUnion_GetSnapshotForMerge{t}
	case *RangeStatsResponse:
		union = &ResponseUnion_RangeStats{t}
	default:
		return false
	}
	ru.Value = union
	return true
}

type reqCounts [40]int32

// getReqCounts returns the number of times each
// request type appears in the batch.
func (ba *BatchRequest) getReqCounts() reqCounts {
	var counts reqCounts
	for _, ru := range ba.Requests {
		switch ru.GetValue().(type) {
		case *RequestUnion_Get:
			counts[0]++
		case *RequestUnion_Put:
			counts[1]++
		case *RequestUnion_ConditionalPut:
			counts[2]++
		case *RequestUnion_Increment:
			counts[3]++
		case *RequestUnion_Delete:
			counts[4]++
		case *RequestUnion_DeleteRange:
			counts[5]++
		case *RequestUnion_ClearRange:
			counts[6]++
		case *RequestUnion_Scan:
			counts[7]++
		case *RequestUnion_BeginTransaction:
			counts[8]++
		case *RequestUnion_EndTransaction:
			counts[9]++
		case *RequestUnion_AdminSplit:
			counts[10]++
		case *RequestUnion_AdminMerge:
			counts[11]++
		case *RequestUnion_AdminTransferLease:
			counts[12]++
		case *RequestUnion_AdminChangeReplicas:
			counts[13]++
		case *RequestUnion_HeartbeatTxn:
			counts[14]++
		case *RequestUnion_Gc:
			counts[15]++
		case *RequestUnion_PushTxn:
			counts[16]++
		case *RequestUnion_ResolveIntent:
			counts[17]++
		case *RequestUnion_ResolveIntentRange:
			counts[18]++
		case *RequestUnion_Merge:
			counts[19]++
		case *RequestUnion_TruncateLog:
			counts[20]++
		case *RequestUnion_RequestLease:
			counts[21]++
		case *RequestUnion_ReverseScan:
			counts[22]++
		case *RequestUnion_ComputeChecksum:
			counts[23]++
		case *RequestUnion_CheckConsistency:
			counts[24]++
		case *RequestUnion_InitPut:
			counts[25]++
		case *RequestUnion_TransferLease:
			counts[26]++
		case *RequestUnion_LeaseInfo:
			counts[27]++
		case *RequestUnion_WriteBatch:
			counts[28]++
		case *RequestUnion_Export:
			counts[29]++
		case *RequestUnion_Import:
			counts[30]++
		case *RequestUnion_QueryTxn:
			counts[31]++
		case *RequestUnion_QueryIntent:
			counts[32]++
		case *RequestUnion_AdminScatter:
			counts[33]++
		case *RequestUnion_AddSstable:
			counts[34]++
		case *RequestUnion_RecomputeStats:
			counts[35]++
		case *RequestUnion_Refresh:
			counts[36]++
		case *RequestUnion_RefreshRange:
			counts[37]++
		case *RequestUnion_GetSnapshotForMerge:
			counts[38]++
		case *RequestUnion_RangeStats:
			counts[39]++
		default:
			panic(fmt.Sprintf("unsupported request: %+v", ru))
		}
	}
	return counts
}

var requestNames = []string{
	"Get",
	"Put",
	"CPut",
	"Inc",
	"Del",
	"DelRng",
	"ClearRng",
	"Scan",
	"BeginTxn",
	"EndTxn",
	"AdmSplit",
	"AdmMerge",
	"AdmTransferLease",
	"AdmChangeReplicas",
	"HeartbeatTxn",
	"Gc",
	"PushTxn",
	"ResolveIntent",
	"ResolveIntentRng",
	"Merge",
	"TruncLog",
	"RequestLease",
	"RevScan",
	"ComputeChksum",
	"ChkConsistency",
	"InitPut",
	"TransferLease",
	"LeaseInfo",
	"WriteBatch",
	"Export",
	"Import",
	"QueryTxn",
	"QueryIntent",
	"AdmScatter",
	"AddSstable",
	"RecomputeStats",
	"Refresh",
	"RefreshRng",
	"GetSnapshotForMerge",
	"RngStats",
}

// Summary prints a short summary of the requests in a batch.
func (ba *BatchRequest) Summary() string {
	if len(ba.Requests) == 0 {
		return "empty batch"
	}
	counts := ba.getReqCounts()
	var buf struct {
		bytes.Buffer
		tmp [10]byte
	}
	for i, v := range counts {
		if v != 0 {
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}
			buf.Write(strconv.AppendInt(buf.tmp[:0], int64(v), 10))
			buf.WriteString(" ")
			buf.WriteString(requestNames[i])
		}
	}
	return buf.String()
}

// The following types are used to group the allocations of Responses
// and their corresponding isResponseUnion_Value union wrappers together.
type getResponseAlloc struct {
	union ResponseUnion_Get
	resp  GetResponse
}
type putResponseAlloc struct {
	union ResponseUnion_Put
	resp  PutResponse
}
type conditionalPutResponseAlloc struct {
	union ResponseUnion_ConditionalPut
	resp  ConditionalPutResponse
}
type incrementResponseAlloc struct {
	union ResponseUnion_Increment
	resp  IncrementResponse
}
type deleteResponseAlloc struct {
	union ResponseUnion_Delete
	resp  DeleteResponse
}
type deleteRangeResponseAlloc struct {
	union ResponseUnion_DeleteRange
	resp  DeleteRangeResponse
}
type clearRangeResponseAlloc struct {
	union ResponseUnion_ClearRange
	resp  ClearRangeResponse
}
type scanResponseAlloc struct {
	union ResponseUnion_Scan
	resp  ScanResponse
}
type beginTransactionResponseAlloc struct {
	union ResponseUnion_BeginTransaction
	resp  BeginTransactionResponse
}
type endTransactionResponseAlloc struct {
	union ResponseUnion_EndTransaction
	resp  EndTransactionResponse
}
type adminSplitResponseAlloc struct {
	union ResponseUnion_AdminSplit
	resp  AdminSplitResponse
}
type adminMergeResponseAlloc struct {
	union ResponseUnion_AdminMerge
	resp  AdminMergeResponse
}
type adminTransferLeaseResponseAlloc struct {
	union ResponseUnion_AdminTransferLease
	resp  AdminTransferLeaseResponse
}
type adminChangeReplicasResponseAlloc struct {
	union ResponseUnion_AdminChangeReplicas
	resp  AdminChangeReplicasResponse
}
type heartbeatTxnResponseAlloc struct {
	union ResponseUnion_HeartbeatTxn
	resp  HeartbeatTxnResponse
}
type gCResponseAlloc struct {
	union ResponseUnion_Gc
	resp  GCResponse
}
type pushTxnResponseAlloc struct {
	union ResponseUnion_PushTxn
	resp  PushTxnResponse
}
type resolveIntentResponseAlloc struct {
	union ResponseUnion_ResolveIntent
	resp  ResolveIntentResponse
}
type resolveIntentRangeResponseAlloc struct {
	union ResponseUnion_ResolveIntentRange
	resp  ResolveIntentRangeResponse
}
type mergeResponseAlloc struct {
	union ResponseUnion_Merge
	resp  MergeResponse
}
type truncateLogResponseAlloc struct {
	union ResponseUnion_TruncateLog
	resp  TruncateLogResponse
}
type requestLeaseResponseAlloc struct {
	union ResponseUnion_RequestLease
	resp  RequestLeaseResponse
}
type reverseScanResponseAlloc struct {
	union ResponseUnion_ReverseScan
	resp  ReverseScanResponse
}
type computeChecksumResponseAlloc struct {
	union ResponseUnion_ComputeChecksum
	resp  ComputeChecksumResponse
}
type checkConsistencyResponseAlloc struct {
	union ResponseUnion_CheckConsistency
	resp  CheckConsistencyResponse
}
type initPutResponseAlloc struct {
	union ResponseUnion_InitPut
	resp  InitPutResponse
}
type leaseInfoResponseAlloc struct {
	union ResponseUnion_LeaseInfo
	resp  LeaseInfoResponse
}
type writeBatchResponseAlloc struct {
	union ResponseUnion_WriteBatch
	resp  WriteBatchResponse
}
type exportResponseAlloc struct {
	union ResponseUnion_Export
	resp  ExportResponse
}
type importResponseAlloc struct {
	union ResponseUnion_Import
	resp  ImportResponse
}
type queryTxnResponseAlloc struct {
	union ResponseUnion_QueryTxn
	resp  QueryTxnResponse
}
type queryIntentResponseAlloc struct {
	union ResponseUnion_QueryIntent
	resp  QueryIntentResponse
}
type adminScatterResponseAlloc struct {
	union ResponseUnion_AdminScatter
	resp  AdminScatterResponse
}
type addSSTableResponseAlloc struct {
	union ResponseUnion_AddSstable
	resp  AddSSTableResponse
}
type recomputeStatsResponseAlloc struct {
	union ResponseUnion_RecomputeStats
	resp  RecomputeStatsResponse
}
type refreshResponseAlloc struct {
	union ResponseUnion_Refresh
	resp  RefreshResponse
}
type refreshRangeResponseAlloc struct {
	union ResponseUnion_RefreshRange
	resp  RefreshRangeResponse
}
type getSnapshotForMergeResponseAlloc struct {
	union ResponseUnion_GetSnapshotForMerge
	resp  GetSnapshotForMergeResponse
}
type rangeStatsResponseAlloc struct {
	union ResponseUnion_RangeStats
	resp  RangeStatsResponse
}

// CreateReply creates replies for each of the contained requests, wrapped in a
// BatchResponse. The response objects are batch allocated to minimize
// allocation overhead.
func (ba *BatchRequest) CreateReply() *BatchResponse {
	br := &BatchResponse{}
	br.Responses = make([]ResponseUnion, len(ba.Requests))

	counts := ba.getReqCounts()

	var buf0 []getResponseAlloc
	var buf1 []putResponseAlloc
	var buf2 []conditionalPutResponseAlloc
	var buf3 []incrementResponseAlloc
	var buf4 []deleteResponseAlloc
	var buf5 []deleteRangeResponseAlloc
	var buf6 []clearRangeResponseAlloc
	var buf7 []scanResponseAlloc
	var buf8 []beginTransactionResponseAlloc
	var buf9 []endTransactionResponseAlloc
	var buf10 []adminSplitResponseAlloc
	var buf11 []adminMergeResponseAlloc
	var buf12 []adminTransferLeaseResponseAlloc
	var buf13 []adminChangeReplicasResponseAlloc
	var buf14 []heartbeatTxnResponseAlloc
	var buf15 []gCResponseAlloc
	var buf16 []pushTxnResponseAlloc
	var buf17 []resolveIntentResponseAlloc
	var buf18 []resolveIntentRangeResponseAlloc
	var buf19 []mergeResponseAlloc
	var buf20 []truncateLogResponseAlloc
	var buf21 []requestLeaseResponseAlloc
	var buf22 []reverseScanResponseAlloc
	var buf23 []computeChecksumResponseAlloc
	var buf24 []checkConsistencyResponseAlloc
	var buf25 []initPutResponseAlloc
	var buf26 []requestLeaseResponseAlloc
	var buf27 []leaseInfoResponseAlloc
	var buf28 []writeBatchResponseAlloc
	var buf29 []exportResponseAlloc
	var buf30 []importResponseAlloc
	var buf31 []queryTxnResponseAlloc
	var buf32 []queryIntentResponseAlloc
	var buf33 []adminScatterResponseAlloc
	var buf34 []addSSTableResponseAlloc
	var buf35 []recomputeStatsResponseAlloc
	var buf36 []refreshResponseAlloc
	var buf37 []refreshRangeResponseAlloc
	var buf38 []getSnapshotForMergeResponseAlloc
	var buf39 []rangeStatsResponseAlloc

	for i, r := range ba.Requests {
		switch r.GetValue().(type) {
		case *RequestUnion_Get:
			if buf0 == nil {
				buf0 = make([]getResponseAlloc, counts[0])
			}
			buf0[0].union.Get = &buf0[0].resp
			br.Responses[i].Value = &buf0[0].union
			buf0 = buf0[1:]
		case *RequestUnion_Put:
			if buf1 == nil {
				buf1 = make([]putResponseAlloc, counts[1])
			}
			buf1[0].union.Put = &buf1[0].resp
			br.Responses[i].Value = &buf1[0].union
			buf1 = buf1[1:]
		case *RequestUnion_ConditionalPut:
			if buf2 == nil {
				buf2 = make([]conditionalPutResponseAlloc, counts[2])
			}
			buf2[0].union.ConditionalPut = &buf2[0].resp
			br.Responses[i].Value = &buf2[0].union
			buf2 = buf2[1:]
		case *RequestUnion_Increment:
			if buf3 == nil {
				buf3 = make([]incrementResponseAlloc, counts[3])
			}
			buf3[0].union.Increment = &buf3[0].resp
			br.Responses[i].Value = &buf3[0].union
			buf3 = buf3[1:]
		case *RequestUnion_Delete:
			if buf4 == nil {
				buf4 = make([]deleteResponseAlloc, counts[4])
			}
			buf4[0].union.Delete = &buf4[0].resp
			br.Responses[i].Value = &buf4[0].union
			buf4 = buf4[1:]
		case *RequestUnion_DeleteRange:
			if buf5 == nil {
				buf5 = make([]deleteRangeResponseAlloc, counts[5])
			}
			buf5[0].union.DeleteRange = &buf5[0].resp
			br.Responses[i].Value = &buf5[0].union
			buf5 = buf5[1:]
		case *RequestUnion_ClearRange:
			if buf6 == nil {
				buf6 = make([]clearRangeResponseAlloc, counts[6])
			}
			buf6[0].union.ClearRange = &buf6[0].resp
			br.Responses[i].Value = &buf6[0].union
			buf6 = buf6[1:]
		case *RequestUnion_Scan:
			if buf7 == nil {
				buf7 = make([]scanResponseAlloc, counts[7])
			}
			buf7[0].union.Scan = &buf7[0].resp
			br.Responses[i].Value = &buf7[0].union
			buf7 = buf7[1:]
		case *RequestUnion_BeginTransaction:
			if buf8 == nil {
				buf8 = make([]beginTransactionResponseAlloc, counts[8])
			}
			buf8[0].union.BeginTransaction = &buf8[0].resp
			br.Responses[i].Value = &buf8[0].union
			buf8 = buf8[1:]
		case *RequestUnion_EndTransaction:
			if buf9 == nil {
				buf9 = make([]endTransactionResponseAlloc, counts[9])
			}
			buf9[0].union.EndTransaction = &buf9[0].resp
			br.Responses[i].Value = &buf9[0].union
			buf9 = buf9[1:]
		case *RequestUnion_AdminSplit:
			if buf10 == nil {
				buf10 = make([]adminSplitResponseAlloc, counts[10])
			}
			buf10[0].union.AdminSplit = &buf10[0].resp
			br.Responses[i].Value = &buf10[0].union
			buf10 = buf10[1:]
		case *RequestUnion_AdminMerge:
			if buf11 == nil {
				buf11 = make([]adminMergeResponseAlloc, counts[11])
			}
			buf11[0].union.AdminMerge = &buf11[0].resp
			br.Responses[i].Value = &buf11[0].union
			buf11 = buf11[1:]
		case *RequestUnion_AdminTransferLease:
			if buf12 == nil {
				buf12 = make([]adminTransferLeaseResponseAlloc, counts[12])
			}
			buf12[0].union.AdminTransferLease = &buf12[0].resp
			br.Responses[i].Value = &buf12[0].union
			buf12 = buf12[1:]
		case *RequestUnion_AdminChangeReplicas:
			if buf13 == nil {
				buf13 = make([]adminChangeReplicasResponseAlloc, counts[13])
			}
			buf13[0].union.AdminChangeReplicas = &buf13[0].resp
			br.Responses[i].Value = &buf13[0].union
			buf13 = buf13[1:]
		case *RequestUnion_HeartbeatTxn:
			if buf14 == nil {
				buf14 = make([]heartbeatTxnResponseAlloc, counts[14])
			}
			buf14[0].union.HeartbeatTxn = &buf14[0].resp
			br.Responses[i].Value = &buf14[0].union
			buf14 = buf14[1:]
		case *RequestUnion_Gc:
			if buf15 == nil {
				buf15 = make([]gCResponseAlloc, counts[15])
			}
			buf15[0].union.Gc = &buf15[0].resp
			br.Responses[i].Value = &buf15[0].union
			buf15 = buf15[1:]
		case *RequestUnion_PushTxn:
			if buf16 == nil {
				buf16 = make([]pushTxnResponseAlloc, counts[16])
			}
			buf16[0].union.PushTxn = &buf16[0].resp
			br.Responses[i].Value = &buf16[0].union
			buf16 = buf16[1:]
		case *RequestUnion_ResolveIntent:
			if buf17 == nil {
				buf17 = make([]resolveIntentResponseAlloc, counts[17])
			}
			buf17[0].union.ResolveIntent = &buf17[0].resp
			br.Responses[i].Value = &buf17[0].union
			buf17 = buf17[1:]
		case *RequestUnion_ResolveIntentRange:
			if buf18 == nil {
				buf18 = make([]resolveIntentRangeResponseAlloc, counts[18])
			}
			buf18[0].union.ResolveIntentRange = &buf18[0].resp
			br.Responses[i].Value = &buf18[0].union
			buf18 = buf18[1:]
		case *RequestUnion_Merge:
			if buf19 == nil {
				buf19 = make([]mergeResponseAlloc, counts[19])
			}
			buf19[0].union.Merge = &buf19[0].resp
			br.Responses[i].Value = &buf19[0].union
			buf19 = buf19[1:]
		case *RequestUnion_TruncateLog:
			if buf20 == nil {
				buf20 = make([]truncateLogResponseAlloc, counts[20])
			}
			buf20[0].union.TruncateLog = &buf20[0].resp
			br.Responses[i].Value = &buf20[0].union
			buf20 = buf20[1:]
		case *RequestUnion_RequestLease:
			if buf21 == nil {
				buf21 = make([]requestLeaseResponseAlloc, counts[21])
			}
			buf21[0].union.RequestLease = &buf21[0].resp
			br.Responses[i].Value = &buf21[0].union
			buf21 = buf21[1:]
		case *RequestUnion_ReverseScan:
			if buf22 == nil {
				buf22 = make([]reverseScanResponseAlloc, counts[22])
			}
			buf22[0].union.ReverseScan = &buf22[0].resp
			br.Responses[i].Value = &buf22[0].union
			buf22 = buf22[1:]
		case *RequestUnion_ComputeChecksum:
			if buf23 == nil {
				buf23 = make([]computeChecksumResponseAlloc, counts[23])
			}
			buf23[0].union.ComputeChecksum = &buf23[0].resp
			br.Responses[i].Value = &buf23[0].union
			buf23 = buf23[1:]
		case *RequestUnion_CheckConsistency:
			if buf24 == nil {
				buf24 = make([]checkConsistencyResponseAlloc, counts[24])
			}
			buf24[0].union.CheckConsistency = &buf24[0].resp
			br.Responses[i].Value = &buf24[0].union
			buf24 = buf24[1:]
		case *RequestUnion_InitPut:
			if buf25 == nil {
				buf25 = make([]initPutResponseAlloc, counts[25])
			}
			buf25[0].union.InitPut = &buf25[0].resp
			br.Responses[i].Value = &buf25[0].union
			buf25 = buf25[1:]
		case *RequestUnion_TransferLease:
			if buf26 == nil {
				buf26 = make([]requestLeaseResponseAlloc, counts[26])
			}
			buf26[0].union.RequestLease = &buf26[0].resp
			br.Responses[i].Value = &buf26[0].union
			buf26 = buf26[1:]
		case *RequestUnion_LeaseInfo:
			if buf27 == nil {
				buf27 = make([]leaseInfoResponseAlloc, counts[27])
			}
			buf27[0].union.LeaseInfo = &buf27[0].resp
			br.Responses[i].Value = &buf27[0].union
			buf27 = buf27[1:]
		case *RequestUnion_WriteBatch:
			if buf28 == nil {
				buf28 = make([]writeBatchResponseAlloc, counts[28])
			}
			buf28[0].union.WriteBatch = &buf28[0].resp
			br.Responses[i].Value = &buf28[0].union
			buf28 = buf28[1:]
		case *RequestUnion_Export:
			if buf29 == nil {
				buf29 = make([]exportResponseAlloc, counts[29])
			}
			buf29[0].union.Export = &buf29[0].resp
			br.Responses[i].Value = &buf29[0].union
			buf29 = buf29[1:]
		case *RequestUnion_Import:
			if buf30 == nil {
				buf30 = make([]importResponseAlloc, counts[30])
			}
			buf30[0].union.Import = &buf30[0].resp
			br.Responses[i].Value = &buf30[0].union
			buf30 = buf30[1:]
		case *RequestUnion_QueryTxn:
			if buf31 == nil {
				buf31 = make([]queryTxnResponseAlloc, counts[31])
			}
			buf31[0].union.QueryTxn = &buf31[0].resp
			br.Responses[i].Value = &buf31[0].union
			buf31 = buf31[1:]
		case *RequestUnion_QueryIntent:
			if buf32 == nil {
				buf32 = make([]queryIntentResponseAlloc, counts[32])
			}
			buf32[0].union.QueryIntent = &buf32[0].resp
			br.Responses[i].Value = &buf32[0].union
			buf32 = buf32[1:]
		case *RequestUnion_AdminScatter:
			if buf33 == nil {
				buf33 = make([]adminScatterResponseAlloc, counts[33])
			}
			buf33[0].union.AdminScatter = &buf33[0].resp
			br.Responses[i].Value = &buf33[0].union
			buf33 = buf33[1:]
		case *RequestUnion_AddSstable:
			if buf34 == nil {
				buf34 = make([]addSSTableResponseAlloc, counts[34])
			}
			buf34[0].union.AddSstable = &buf34[0].resp
			br.Responses[i].Value = &buf34[0].union
			buf34 = buf34[1:]
		case *RequestUnion_RecomputeStats:
			if buf35 == nil {
				buf35 = make([]recomputeStatsResponseAlloc, counts[35])
			}
			buf35[0].union.RecomputeStats = &buf35[0].resp
			br.Responses[i].Value = &buf35[0].union
			buf35 = buf35[1:]
		case *RequestUnion_Refresh:
			if buf36 == nil {
				buf36 = make([]refreshResponseAlloc, counts[36])
			}
			buf36[0].union.Refresh = &buf36[0].resp
			br.Responses[i].Value = &buf36[0].union
			buf36 = buf36[1:]
		case *RequestUnion_RefreshRange:
			if buf37 == nil {
				buf37 = make([]refreshRangeResponseAlloc, counts[37])
			}
			buf37[0].union.RefreshRange = &buf37[0].resp
			br.Responses[i].Value = &buf37[0].union
			buf37 = buf37[1:]
		case *RequestUnion_GetSnapshotForMerge:
			if buf38 == nil {
				buf38 = make([]getSnapshotForMergeResponseAlloc, counts[38])
			}
			buf38[0].union.GetSnapshotForMerge = &buf38[0].resp
			br.Responses[i].Value = &buf38[0].union
			buf38 = buf38[1:]
		case *RequestUnion_RangeStats:
			if buf39 == nil {
				buf39 = make([]rangeStatsResponseAlloc, counts[39])
			}
			buf39[0].union.RangeStats = &buf39[0].resp
			br.Responses[i].Value = &buf39[0].union
			buf39 = buf39[1:]
		default:
			panic(fmt.Sprintf("unsupported request: %+v", r))
		}
	}
	return br
}
