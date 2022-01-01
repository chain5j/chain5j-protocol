mock-gen:
	mockgen -destination ./mock/mock_apis.go -package mock -source protocol/apis.go
	mockgen -destination ./mock/mock_apps.go -package mock -source protocol/apps.go
	mockgen -destination ./mock/mock_blockchain.go -package mock -source protocol/blockchain.go
	mockgen -destination ./mock/mock_broadcaster.go -package mock -source protocol/broadcaster.go
	mockgen -destination ./mock/mock_config.go -package mock -source protocol/config.go
	mockgen -destination ./mock/mock_consensus.go -package mock -source protocol/consensus.go
	mockgen -destination ./mock/mock_database.go -package mock -source protocol/database.go
	mockgen -destination ./mock/mock_node.go -package mock -source protocol/node.go
	mockgen -destination ./mock/mock_nodekey.go -package mock -source protocol/nodekey.go
	mockgen -destination ./mock/mock_p2p.go -package mock -source protocol/p2p.go
	mockgen -destination ./mock/mock_packer.go -package mock -source protocol/packer.go
	mockgen -destination ./mock/mock_permission.go -package mock -source protocol/permission.go
	mockgen -destination ./mock/mock_syncer.go -package mock -source protocol/syncer.go
	mockgen -destination ./mock/mock_txpool.go -package mock -source protocol/txpool.go
	mockgen -destination ./mock/mock_vm.go -package mock -source protocol/vm.go

mock-dep:
	go get -u github.com/golang/mock/gomock
	go get -u github.com/golang/mock/mockgen
