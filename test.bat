go test ./internal/constructor_logic/logic/constructor_test.go -v --race > testresults/logical_block.testres &
go test ./... --race > testresults/worldwide.testres &
wait