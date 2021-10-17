// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract Riddler is ERC721URIStorage, Ownable {
    
    struct Riddle {
        string riddle;
        string answer;
    }

    mapping(uint => address) public riddlesSolvedBy;
    Riddle[] private riddles;

    constructor() 
    ERC721("Riddler", "RDL") {}
    
    
    function addRiddle(string memory riddle, string memory answer) 
    public 
    onlyOwner 
    returns (uint){
        Riddle memory r = Riddle(riddle, answer);
        riddles.push(r);
        
        return riddles.length;
    }
    
    function getRiddle(uint index) 
    public
    view
    returns (Riddle memory){
        return riddles[index];
    }
    
    function solve(uint idx, string memory answer, string memory tokenURI) 
    public
    returns (bool) {
        
        Riddle memory riddle = getRiddle(idx);
        
        if (compareStrings(riddle.answer, answer)) {
            
            _mint(msg.sender, idx);
            _setTokenURI(idx, tokenURI);
            
            riddlesSolvedBy[idx] = msg.sender;
            return true;
        }
        
        return false;
        
    }
    
    
    function compareStrings(string memory a, string memory b) private pure returns (bool) {
        return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))));
    }
    
}