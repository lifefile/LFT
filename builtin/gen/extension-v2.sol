pragma solidity 0.4.24;
import './extension.sol';

/// @title Extension extends EVM global functions.
contract ExtensionV2 is Extension {
    function txGasPayer() public view returns(address) {
        return ExtensionV2Native(this).native_txGasPayer();
    }
}

contract ExtensionV2Native is ExtensionNative {    
    function native_txGasPayer()public view returns(address);
}
