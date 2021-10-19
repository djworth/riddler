// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract Riddler is ERC721URIStorage, Ownable {

    using EnumerableMap for EnumerableMap.UintToAddressMap;
    using Counters for Counters.Counter;
    
    struct Riddle {
        uint id;
        string tokenURI;
        string riddle;
        string answer;
    }

    mapping(uint => Riddle) private riddles;
    EnumerableMap.UintToAddressMap private riddlesSolvedBy;
    Counters.Counter public riddleCounter;

    constructor() 
    ERC721("Riddler", "RDL") {}
        
    function addRiddle(uint id, string memory tokenURI, string memory riddle, string memory answer) 
    public 
    onlyOwner 
    returns (bool){
        Riddle memory r = Riddle(id, tokenURI, riddle, answer);
        riddles[id] = r;
        Counters.increment(riddleCounter);
        
        return true;
    }
    
    function getRiddle(uint index) 
    public
    view
    returns (string memory){
        return riddles[index].riddle;
    }

    function numOfRiddles()
    public
    view
    returns (uint256) {
        return Counters.current(riddleCounter);
    }

    function hasRiddleBeenSolved(uint id)
    public
    view
    returns (bool) {
        return EnumerableMap.contains(riddlesSolvedBy, id);
    }
    
    function solve(uint id, string memory answer) 
    payable 
    public
    returns (bool) {
        
        require(EnumerableMap.contains(riddlesSolvedBy, id) == false, "Riddle has already been solved");

        Riddle memory riddle = riddles[id];
        
        if (compareStrings(riddle.answer, answer)) {
            
            _mint(msg.sender, id);
            _setTokenURI(id, riddle.tokenURI);
            
            EnumerableMap.set(riddlesSolvedBy, id, msg.sender);
            return true;
        }
        
        return false;
    }
    
    function compareStrings(string memory a, string memory b) private pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }
}