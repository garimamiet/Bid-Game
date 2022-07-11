// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract bidding {
    uint256 public constant DURATION = 60 minutes;
    uint256 public constant EXTENSION = 10 minutes;
    uint256 public expiry;
    uint256 public start;
    uint256 public constant COMMISSION = 500; // commission in basis points.
    uint256 public rewardCollected;
    uint256 private commissionCollected;
    uint256 public lastBid;
    bool public readyToStart;
    address public admin;
    address public lastBidder;

    event bidStarted(uint256 startTime);
    event newBid(address user, uint256 bidValue);
    event rewardClaimed(address winner, uint256 reward);
    event commissionWithdrawn();

    constructor(address _admin) {
        admin = _admin;
        lastBid = 0;
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
        commissionCollected += (msg.value * COMMISSION) / 10000;
        expiry += EXTENSION;
        lastBid = msg.value;
        lastBidder = msg.sender;
        emit newBid(msg.sender, msg.value);
    }

    function checkWinner() public view returns (address) {
        if (block.timestamp > expiry) {
            return lastBidder;
        } else {
            return address(0);
        }
    }

    function claimReward() public {
        require(block.timestamp > expiry, "bidding not over yet");
        require(msg.sender == lastBidder, "Not Winner");
        require(rewardCollected > 0, "no reward to collect");
        uint256 temp = rewardCollected;
        rewardCollected = 0;
        readyToStart = true;
        lastBidder = address(0);
        lastBid = 0;
        start = 0;
        expiry = 0;
        (bool success, ) = msg.sender.call{value: temp}("");
        require(success, "claim reward failed");
        emit rewardClaimed(msg.sender, temp);
    }

    function withdrawCommission() public onlyAdmin {
        require(block.timestamp > expiry, "bidding not over yet");
        require(commissionCollected > 0, "No commission");
        uint256 temp = commissionCollected;
        commissionCollected = 0;
        (bool success, ) = admin.call{value: temp}("");
        require(success, "commission transfer failed");
        emit commissionWithdrawn();
    }
}
