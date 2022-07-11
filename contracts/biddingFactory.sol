// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./bidding.sol";

contract biddingFactory {
    address public admin;
    address[] activeGames;

    constructor() {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "only Admin");
        _;
    }

    function transferOwnership(address newAdmin) public onlyAdmin {
        admin = newAdmin;
    }

    function getActiveGames(uint256 _id) public view returns (address) {
        if (_id <= activeGames.length - 1) {
            return activeGames[_id];
        } else {
            return address(0);
        }
    }

    function numberOfActiveGames() public view returns (uint256) {
        return activeGames.length;
    }

    function create() public onlyAdmin {
        bidding myBidding = new bidding(admin);
        activeGames.push(address(myBidding));
    }
}
