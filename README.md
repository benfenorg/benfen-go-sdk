# go-benfen-sdk
## api
```
api getDynamicFieldObject | code client_call.go function GetDynamicFieldObject
api getOwnedObjects | code client_call.go function GetOwnedObjects
api queryEvents | code client_stake.go function queryEvents
api queryTransactionBlocks | code client_call.go function queryTransactionBlocks
api resolveNameServiceAddress | code client_stake.go function resolveNameServiceAddress
api resolveNameServiceNames | code client_stake.go function resolveNameServiceNames
api subscribeEvent | code client_stake.go function SubscribeEventString api无法跑通
api subscribeTransaction | code no api无法跑通
api getStakes | code client_stake.go function GetStakes
api getStakesByIds | code client_stake.go function GetStakesByIds
api getMoveFunctionArgTypes | code client_stake.go function getMoveFunctionArgTypes
api getNormalizedMoveFunction | code client_stake.go function getNormalizedMoveFunction
api getNormalizedMoveModule | code client_stake.go function getNormalizedMoveModule
api getNormalizedMoveModulesByPackage | code client_stake.go function getNormalizedMoveModulesByPackage
api getNormalizedMoveStruct | code client_stake.go function getNormalizedMoveStruct
api getEvents | code client_call.go function GetEvents
api getObject | code client_call.go function GetObject
api multiGetObjects | code client_call.go function MultiGetObjects
api tryGetPastObject | code client_call.go function TryGetPastObject
api tryMultiGetPastObjects | code client_stake.go function tryMultiGetPastObjects
api batchTransaction | code client_call.go function BatchTransaction
api mergeCoins | code client_call.go function MergeCoins
api moveCall | code client_call.go function MoveCall
api pay | code client_call.go function Pay
api payAllBfc | code client_call.go function PayAllBFC
api payBfc | code client_call.go function PayBFC
api publish | code client_call.go function Publish
api requestAddStake | code client_stake.go function RequestAddStake
api requestWithdrawStake | code client_stake.go function RequestWithdrawStake
api splitCoin | code client_call.go function SplitCoin
api splitCoinEqual | code client_call.go function SplitCoinEqual
api transferObject | code client_call.go function TransferObject
api transferBfc | code client_call.go function TransferBFC
api devInspectTransactionBlock | code client_call.go function DevInspectTransactionBlock
api dryRunTransactionBlock | code client_call.go function DryRunTransaction
api executeTransactionBlock | code client_call.go function ExecuteTransactionBlock
api getDynamicFields | code client_call.go function GetDynamicFields
api getAllBalances | code client_call.go function GetAllBalances
api getAllCoins | code client_call.go function GetAllCoins
api getBalance | code client_call.go function GetBalance
api getCoinMetadata | code client_call.go function GetCoinMetadata
api getCoins | code client_call.go function GetCoins
api getTotalSupply | code client_call.go function GetTotalSupply
api getCommitteeInfo | code client_stake.go function GetCommiteeInfo
api getLatestSuiSystemState | code client_stake.go function GetLatestSuiSystemState
api getReferenceGasPrice | code client_call.go function GetReferenceGasPrice
api getValidatorsApy | code client_stake.go function GetValidatorsApy
api getChainIdentifier | code client_call.go function GetChainIdentifier
api getCheckpoint | code client_stake.go function GetCheckPoint
api getCheckpoints | code client_stake.go function GetCheckPoints
api getLatestCheckpointSequenceNumber | code client_call.go function GetLatestCheckpointSequenceNumber
api getProtocolConfig | code client_stake.go function GetProtocolConfig
api getTotalTransactionBlocks | code client_call.go function GetTotalTransactionBlocks
api getTransactionBlock | code client_call.go function GetTransactionBlock
api multiGetTransactionBlocks | code client_stake.go function multiGetTransactionBlocks
```
## cli
### keytool
```
cli zk-login-sign-and-execute-tx | code keytool_cmd.go ZkLoginSignAndExecuteTx
cli update-alias | code keytool_cmd.go UpdateAlias
cli list | code keytool_cmd.go list
cli export | code keytool_cmd.go export
cli zk-login-insecure-sign-personal-message | code keytool_cmd.go ZkLoginInsecureSignPersonalMessage
cli zk-login-sig-verify | code keytool_cmd.go ZkLoginSigVerify
cli zk-login-enter-token | code keytool_cmd.go ZkLoginEnterToken
cli unpack | code keytool_cmd.go Unpack
cli sign-kms | code keytool_cmd.go SignKms
cli sign | code keytool_cmd.go Sign
cli multi-sig-combine-partial-sig-legacy | code keytool_cmd.go MultiSigCombinePartialSigLegacy
cli multi-sig-combine-partial-sig | code keytool_cmd.go MultiSigCombinePartialSig
cli multi-sig-address | code keytool_cmd.go MultiSigAddress
cli load-keypair | 
cli import | 
cli generate-with-name | 
cli generate | 
cli decode-multi-sig | code keytool_cmd.go DecodeMultiSig
cli decode-raw-transaction |
cli decode-or-verify-tx |
cli convert | code keytool_cmd.go Convert
```
### client
```
cli active-address | code client_cmd.go ActiveAddress
cli active-env | code client_cmd.go ActiveEnv
cli addresses | code client_cmd.go Address
cli gas | code client_cmd.go Gas
cli object | code client_cmd.go Object
cli balance | code client_cmd.go Balance
cli chain-identifier | code client_cmd.go ChainIdentifier
cli envs | code client_cmd.go Envs
cli new-address | code client_cmd.go NewAddress
cli objects | code client_cmd.go Objects
cli merge-coin | code client_cmd.go MergeCoin
cli dynamic-field | code client_cmd.go DynamicField
cli split-coin | code client_cmd.go SplitCoin
cli pay | code client_cmd.go Pay
cli pay-bfc | code client_cmd.go PayBfc
cli transfer | code client_cmd.go Transfer
cli transfer-bfc | code client_cmd.go TransferBfc
cli switch | code client_cmd.go Switch
cli profile-transaction | code client_cmd.go ProfileTransaction
cli tx-block | code client_cmd.go TxBlock
cli publish | code client_cmd.go
cli call | code client_cmd.go Call
cli verify-bytecode-meter | code client_cmd.go Call VerifyBytecodeMeter
cli verify-source | code client_cmd.go Call VerifySource
cli replay-transaction | code client_cmd.go Call ReplayTransaction
cli replay-batch | code client_cmd.go Call ReplayBatch
cli replay-checkpoint | code client_cmd.go Call ReplayCheckpoint
cli faucet | code client_cmd.go Call Faucet
cli execute-signed-tx | code client_cmd.go Call ExecuteSignedTx
cli execute-combined-signed-tx | code client_cmd.go Call ExecuteCombinedSignedTx
```
### validator
```
cli make-validator-info | code validator_cmd.go MakeValidatorInfo
cli become-candidate | code validator_cmd.go BecomeCandidate                  
cli join-committee | code validator_cmd.go JoinCommittee
cli leave-committee | code validator_cmd.go LeaveCommittee
cli display-metadata | code validator_cmd.go DisplayMetadata       
cli update-metadata | code validator_cmd.go UpdateMetadata             
cli update-gas-price | code validator_cmd.go UpdateGasPrice   
cli report-validator | code validator_cmd.go ReportValidator   
cli serialize-payload-pop | code validator_cmd.go SerializePayloadPop
cli display-gas-price-update-raw-txn | code validator_cmd.go DisplayGasPriceUpdateRawTx
cli init-admin-capability | code validator_cmd.go InitAdminCapability
cli add-admin-capability | code validator_cmd.go AddAdminCapability
cli remove-admin-capability | code validator_cmd.go RemoveAdminCapability
cli add-operation-capability | code validator_cmd.go AddOperationCapability
cli remove-operation-capability | code validator_cmd.go RemoveOperationCapability
cli set-oracle-price-address | code validator_cmd.go SetOraclePriceAddress
```
## api cli自动检查(用于网络测试)
```
api benfen_rpc_client_check_test.go TestFeatures
cli
```