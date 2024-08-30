package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "ewm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ewm"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo-value-"
)

const (
	// Proof session constants
	DefaultSessionDuration     = time.Duration((float64(1) / 12) * 3_600 * 1000_000_000) // 5 mins
	TimestampLayout            = "2006-01-02 15:04:05.999999999 +0000 UTC"
	DefaultProofSubmissionsReq = 2
	DefaultSessionQuorumReq    = 66
	DefaultSyncLag             = 1000
	DefaultBridgeTxDayFreq     = 2

	// Fifo proof session index
	NoFifoIndex = "-1"

	// Proof session event types
	ProofSessionEventCreated      = "proof-session-created"
	ProofSessionEventParticipated = "proof-session-participated"
	ProofSessionEventExpired      = "proof-session-expired"
	ProofSessionEventConcluded    = "proof-session-concluded"
	ProofSessionEventStored       = "proof-session-stored"

	// Proof session event vars
	ProofSessionEventCreator          = "creator"
	ProofSessionEventIndex            = "proof-session-index"
	ProofSessionEventFinalized        = "finalized"
	ProofSessionEventSubmissions      = "submissions"
	ProofSessionEventExpiredIndex     = "proof-session-index" //maybe redundant
	ProofSessionEventExpiredProofHash = "hash"
	ProofSessionEventExpiredCount     = "count"
	ProofSessionEventStoredCount      = "stored"

	// Chain role event types
	ChainRoleEventInit                  = "chain-role-initialized"
	ChainRoleEventAssign                = "chain-role-assigned"
	ChainRoleEventAssignOpVal           = "chain-role-assign-operator-validator"
	ChainRoleEventRemove                = "chain-role-removed"
	ChainRoleEventEnableProofValidator  = "chain-role-enable-proof-validator"
	ChainRoleEventDisableProofValidator = "chain-role-disable-proof-validator"

	// Chain role event vars
	ChainRoleEventProofType      = "proof-type"
	ChainRoleEventCreator        = "creator"
	ChainRoleEventIndex          = "chain-role-index"
	ChainRoleEventSequencer      = "sequencer"
	ChainRoleEventSequencerIndex = "sequencer-index"
	ChainRoleEventOperator       = "operator"
	ChainRoleEventOperatorIndex  = "operator-index"
	ChainRoleEventValidator      = "validator"
	ChainRoleEventValidatorIndex = "validator-index"

	// Chain config event types
	ChainConfigEventInit   = "chain-config-initialized"
	ChainConfigEventUpdate = "chain-config-updated"

	// Chain config event vars
	ChainConfigEventCreator    = "creator"
	ChainConfigEventIndex      = "chain-config-index"
	ChainConfigEventProofType  = "proof-type"
	ChainConfigEventTimestamp  = "timestamp"
	ChainConfigUpdateEventType = "config-update-type"
)
