//SPDX-License-Identifier: MIT
//
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
//✅ 作业2：在测试网上发行一个图文并茂的 NFT
contract JxdNFT is ERC721 {

    //浏览器查看:https://sepolia.etherscan.io/tx/0x20da031277b78833201e5b73d90a81b03d2618fe080d31ca78aa3ed1037dfce4
    //合约地址：0x70eBA544Ac2a0B589178b64217399789509cBcE7
    //部署地址：0x8aaCA58C897F28d2b0DedeBe63cB6eef03EA3E9B

    uint256 private _nextTokenId ;

    mapping(uint256 tokenId => string tokenURI) private _tokenURIs;

    constructor() ERC721("JxdNFT","JxdNFT"){
        _nextTokenId=1;
        Ownable(msg.sender);//表示初始化持有人地址,默认合约发起人是持有者
    }
    //这个需要实现
    function mintNFT(string memory url) public {
        uint256 tokenId =_nextTokenId++;
        _safeMint(msg.sender, tokenId);
        _tokenURIs[tokenId]=url;
    }
    //重写了url 方法 需要拿到tokenURL
    function tokenURI(uint256 tokenId) public view override(ERC721) returns (string memory) {
        return _tokenURIs[tokenId];
    }
}