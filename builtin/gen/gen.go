package gen

//go:generate rm -rf ./compiled/
//go:generate solc --optimize-runs 200 --overwrite --bin-runtime --abi -o ./compiled authority.sol energy.sol executor.sol extension.sol extension-v2.sol measure.sol params.sol prototype.sol
//go:generate go-bindata -nometadata -ignore=_ -pkg gen -o bindata.go compiled/
