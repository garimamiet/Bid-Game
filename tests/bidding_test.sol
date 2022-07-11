// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;
import "remix_tests.sol";
import "remix_accounts.sol";
import "./bidding.sol";

contract testBidding is bidding(TestsAccounts.getAccount(0)) {
    address acc0;
    address acc1;
    address acc2;

    function beforeAll() public {
        acc0 = TestsAccounts.getAccount(0);
        acc1 = TestsAccounts.getAccount(1);
        acc2 = TestsAccounts.getAccount(2);
    }

    function testStartBidding() public {
        startBidding();
        Assert.equal(readyToStart, false, "");
    }

    /// #value: 1000000000000000000
    /// #sender: account-1
    function testBid1() public payable {
        Assert.equal(msg.sender, acc1, "sender is not correct");
        Assert.equal(msg.value, 1000000000000000000, "value should be 1 ether");
        uint256 expiryBeforeBid = expiry;
        bid();
        uint256 expiryAfterBid = expiry;
        Assert.equal(
            expiryAfterBid - expiryBeforeBid,
            10 minutes,
            "expiry should shift by 10 minutes"
        );
        Assert.equal(
            rewardCollected,
            950000000000000000,
            "rewardCollected ahould be 0.95 ether"
        );
        Assert.equal(lastBid, 1000000000000000000, "lastBid should be 1 ether");
        Assert.equal(lastBidder, acc1, "lastBidder should be sender");
    }

    /// #value: 2000000000000000000
    /// #sender: account-2
    function testBid2() public payable {
        Assert.equal(msg.sender, acc2, "sender is not correct");
        Assert.equal(msg.value, 2 ether, "value is not correct");
        uint256 expiryBeforeBid = expiry;
        bid();
        uint256 expiryAfterBid = expiry;
        Assert.equal(
            expiryAfterBid - expiryBeforeBid,
            10 minutes,
            "expiry should shift by 10 minutes"
        );
        Assert.equal(
            rewardCollected,
            2850000000000000000,
            "rewardCollected should be 2.85 ether"
        );
        Assert.equal(lastBid, 2 ether, "lastBid should be 2 ether");
        Assert.equal(lastBidder, acc2, "lastBidder should be account-2");
    }

    /// #sender: account-0
    function testWinner() public {
        expiry -= 5000;
        address winner = checkWinner();
        Assert.equal(winner, acc2, "winner should be account-2");
    }

    /// #sender: account-2
    function testClaimReward() public {
        uint256 balanceBeforeReward = acc2.balance;
        claimReward();
        uint256 balanceAfterReward = acc2.balance;
        Assert.greaterThan(
            balanceAfterReward,
            balanceBeforeReward,
            "balance should increase after reward"
        );
        Assert.equal(rewardCollected, 0, "rewardCollected should be reset");
        Assert.equal(readyToStart, true, "readyToStart should gets reset");
        Assert.equal(lastBidder, address(0), "lastBidder should gets reset");
        Assert.equal(lastBid, 0, "lastBid should gets reset");
        Assert.equal(start, 0, "start should gets reset");
        Assert.equal(expiry, 0, "expiry should gets reset");
    }

    /// #sender: account-0
    function testStartBiddingAgain() public {
        startBidding();
        Assert.equal(readyToStart, false, "");
    }
}
