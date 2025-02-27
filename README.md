# AetherLink Protocol - Universal Decentralized Content Access

---

## **Overview**  
The **AetherLink Protocol** (based on the `ai3://` protocol) is a decentralized content access standard designed to enable seamless access to content stored on any decentralized network (e.g. Autonomys, IPFS, Arweave, Filecoin) using a unified URL format. By binding user identities to content identifiers (CIDs) in smart contracts, AetherLink enables direct addressing of decentralized content without relying on centralized servers or complex tools.

`ai3-gateway` is a temporary compatibility layer that converts `ai3://` URLs to `https://` links, ensuring compatibility with current browser ecosystems. In the future, as the protocol standardizes, native parsers (such as browser extensions or protocol handlers) will gradually replace the core role of the gateway.

---

## **Problems Solved**  
Current decentralized content access faces the following issues:  
1. **High Technical Barriers**: Users must understand blockchain tools (e.g., wallets, CID generators) to access content.  
2. **Protocol Fragmentation**: Different storage networks (Autonomys, Filecoin, Arweave, etc.) use independent addressing methods, lacking a unified entry point.  
3. **Identity and Content Disconnection**: Blockchain addresses are hard to remember and cannot be directly mapped to user-friendly names and file hierarchies.

The AetherLink Protocol addresses these issues in the following ways:  
- **Unified URL Standard**: The `ai3://<name>/<path>` format is compatible with all decentralized storage networks.  
- **Lightweight On-Chain Registration**: Smart contracts only manage name resolution and metadata mapping (e.g., storage type + CID), while the content itself remains fully decentralized.  
- **Progressive Decentralization**: A temporary gateway provides transitional support, ultimately aiming for native browser integration of the `ai3://` protocol.

---

## **Goals**  
1. **Be the HTTP for Decentralized Content**: Define a universal protocol standard that covers multi-chain and multi-storage network content addressing.  
2. **Lower the User Barrier**: Provide a Web2-like URL experience, allowing non-technical users to easily access decentralized content.  
3. **Full Compatibility with Existing Ecosystem**: Support flexible resolution of content from IPFS CIDs to Arweave transaction IDs without requiring changes to existing storage infrastructure.

---

## **Core Features**  
- **Decentralized Identity Binding**  
  Users can bind a human-readable name (e.g., `alice`) to their address, supporting cross-chain identity aggregation (with future expansion).  
- **Multi-Network Storage Compatibility**  
  Files can be uploaded to any decentralized storage (e.g., Autonomys, IPFS, Arweave, Storj), with the contract only recording the storage type and CID/path.  
- **Native Protocol Parsing**  
  The `ai3://` URL directly maps to the target content, eliminating the need for an intermediary gateway (long-term goal).  
  Later, this can be extended to `web3://` to support multiple decentralized networks.

---

# aetherlink

1. Call `registerName` to bind the name with the user address
```
./dist/aether contract register --name xuanwu
```

2. Call `registerFileCID` to bind the file name and CID `baxxxxx`
```
./dist/aether contract bind --cid baxxxxx --name hello.html
```

3. Call `getFileCID` to query the file CID by specifying the file name
```
./dist/aether contract resolve --name hello.html --address 0xXXXXXXXX
```

4. Upload the file
```
./dist/aether upload --path ./resources/hello.html --filename hello.html
```