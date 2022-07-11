// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract bidding {
    uint256 public constant DURATION = 60 minutes;
    uint256 public constant EXTENSION = 10 minutes;
    uint256 public expiry;
    uint256 public start;
    uint256 public constant COMMISSION = 500; // commission in basis points.
    uint256 public rewardCollected;
    uint256 public lastBid;
    bool public readyToStart;
    address payable public admin;
    address public lastBidder;

    event bidStarted(uint256 startTime);
    event newBid(address user, uint256 bidValue);
    event rewardDistributed(address winner, uint256 reward);
    event commissionWithdrawn();

    constructor() {
        admin = payable(msg.sender);

        readyToStart = true;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "only Admin can do");
        _;
    }

    function transferOwnership(address payable newAdmin) public onlyAdmin {
        admin = newAdmin;
    }

    function startBidding() public onlyAdmin {
        require(readyToStart, "not ready");
        lastBid = 0;
        start = block.timestamp;
        expiry = start + DURATION;
        readyToStart = false;
        emit bidStarted(start);
    }

    function bid() public payable {
        require(msg.value > lastBid, "less than last bid");
        require(block.timestamp < expiry, "bidding expired");

        rewardCollected += (msg.value * (10000 - COMMISSION)) / 10000;
        expiry += EXTENSION;
        lastBid = msg.value;
        lastBidder = msg.sender;
        emit newBid(msg.sender, msg.value);
    }

    function distributeReward() public onlyAdmin {
        require(block.timestamp > expiry, "bidding not over yet");
        (bool success, ) = payable(lastBidder).call{value: rewardCollected}("");
        require(success, "reward distribution failed");
        emit rewardDistributed(lastBidder, rewardCollected);
        rewardCollected = 0;
        readyToStart = true;
        lastBidder = address(0);
        start = 0;
        expiry = 0;
    }

    function withdrawCommission() public onlyAdmin {
        require(readyToStart, "bidding going on");
        (bool success, ) = admin.call{value: address(this).balance}("");
        require(success, "commission transfer failed");
        emit commissionWithdrawn();
    }
}
