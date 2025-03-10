// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FileRegistry {
    // 记录每个用户绑定的名称 (私有)
    mapping(address => string) private userNames;
    // 记录每个用户文件名称与CID的绑定关系 (私有)
    mapping(string => mapping(address => string)) private userFileCIDs;
    // 记录已注册的名称 (私有)
   mapping(string => address) private nameToAddress;

    // 事件：绑定名称
    event NameRegistered(address indexed user, string name);
    // 事件：文件CID绑定
    event FileCIDRegistered(address indexed user, string fileName, string cid);

    // 用户绑定名称
    function registerName(string memory name) public {
        // 检查名称是否已经被绑定
        require(bytes(name).length > 0, "Name cannot be empty");
        require(bytes(userNames[msg.sender]).length == 0, "User has already registered a name");
        require(nameToAddress[name] == address(0), "Name already registered");

        // 绑定名称
        userNames[msg.sender] = name;
        nameToAddress[name] = msg.sender;
        emit NameRegistered(msg.sender, name);
    }

    // 用户绑定文件的CID
    function registerFileCID(string memory fileName, string memory cid) public {
        // 确保用户已经绑定名称
        require(bytes(userNames[msg.sender]).length > 0, "User must register a name first");

        // 检查文件CID是否已绑定
        require(bytes(userFileCIDs[fileName][msg.sender]).length == 0, "File CID already registered");
        // 绑定文件CID
        userFileCIDs[fileName][msg.sender] = cid;
        emit FileCIDRegistered(msg.sender, fileName, cid);
    }

    // 获取用户的名称
    function getUserName(address user) public view returns (string memory) {
        return userNames[user];
    }

    // 获取用户文件的CID
    function getFileCID(string memory fileName, address user) public view returns (string memory) {
        string memory cid = userFileCIDs[fileName][user];
        require(bytes(cid).length > 0, "File not found");
        return cid;
    }

    // 检查名称是否已注册
    function getUserAddress(string memory name) public view returns (address) {
        return nameToAddress[name];
    }
}
