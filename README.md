## **AetherLink Protocol - Universal Decentralized Content Access**



### **Quick Links**
- 🌐 **Website**: [https://autonomys.site](https://autonomys.site)
- 🔌 **MetaMask Snap**: [ANS-Snap](https://github.com/AetherLinkLabs/ans-mmsnap)

---

### **Overview**  
The **AetherLink Protocol** (based on the `ai3://` protocol) is a decentralized content access standard designed to enable seamless content discovery and retrieval from decentralized networks. Initially, it focuses on integrating with the **Autonomys** network, leveraging a unified URL format that binds user identities to content identifiers (CIDs) through smart contracts. While the protocol currently prioritizes Autonomys compatibility, it is architected for future expansion to support multiple decentralized storage networks, positioning itself as a universal access standard.

The `ai3-gateway` acts as a temporary compatibility layer that converts `ai3://<name>/<path>` URLs into network-specific links. For example, a URL like `ai3://alice.research/file1` would be converted to `http://alice.research.gateway.com/file1`, ensuring compatibility with existing browser ecosystems. As the protocol matures, native parsers will gradually replace the gateway, enabling direct support for `ai3://` URLs in browsers.

---

### **Problems Solved**  
The AetherLink Protocol addresses several critical challenges in decentralized content access:

- **Technical Complexity**: Eliminates the need for blockchain tool expertise by providing an intuitive URL-based access model.
- **Network-Specific Silos**: Initially focuses on resolving Autonomys content discoverability, while laying the foundation for seamless multi-network interoperability.
- **Identity-Content Gap**: Resolves the challenge of mapping cryptographic blockchain addresses to human-readable names, facilitating user-friendly content navigation and file hierarchies.

---

### **Evolution Roadmap**  

- **Phase 1 - Autonomys Foundation**  
  - Exclusive support for Autonomys network storage  
  - Basic name/CID binding contracts  
  - Gateway-based HTTPS conversion for browser compatibility  

- **Phase 2 - Multi-Network Expansion**  
  - Modular adapters to support IPFS, Arweave, Filecoin  
  - Cross-network redundancy system for enhanced reliability  
  - Storage-agnostic CID resolution  

- **Phase 3 - Protocol Standardization**  
  - Native browser integration for `ai3://` URLs  
  - Decentralized gateway network for better protocol distribution  
  - ENS-like advanced contract system for more flexible asset management  

---

### **Core Features**  

#### **Current Implementation**  

- **Autonomys-First Storage**  
  - Dedicated integration with Autonomys network infrastructure for content storage and retrieval  
- **Identity Anchoring**  
  - Bind human-readable names (e.g., `alice.research`) to CIDs on Autonomys, improving usability  
- **Name-to-Addr Resolution**  
  - Maps human-readable names to their corresponding `Addr` , allowing for convenient identity-based interactions. (A MetaMask Snap submitting, soon you'll be able to send transactions or messages by name directly from MetaMask, check out [ANS-Snap](https://github.com/AetherLinkLabs/ans-mmsnap) for more details.)
- **Gateway Transition**  
  - Temporary HTTPS conversion layer ensures compatibility with existing browsers while the protocol is evolving  

#### **Future Vision**  

- **Universal Storage Abstraction**  
  - A unified interface for accessing content across networks such as Autonomys, IPFS, Arweave, and Storj  
- **Cross-Network Resolution**  
  - A single URL format resolving content from multiple decentralized storage backends  
- **Content Portability**  
  - Seamless migration of content between supported storage networks  

---

## **Getting Started**

### **Prerequisites**
- Go version > 1.22
- Git

### **Installation**
1. Clone the repository and navigate to the project directory:
  ```sh
  git clone https://github.com/AetherLinkLabs/aetherlink.git && cd aetherlink
  ```

2. Compile the project:
  ```sh
  make storage
  ```

3. Set up environment variables:
  ```sh
  cp .env.sample .env
  ```
  Replace the placeholders in the `.env` file with your actual keys:
  ```
  TAURUS_API_TOKEN=<YOUR_TAURUS_API_TOKEN>
  TAURUS_PRIVATE_KEY=<YOUR_TAURUS_PRIVATE_KEY>
  ```

### **Usage**

#### **Register a Name**
Bind a name with the user address:
```sh
./dist/storage contract register --name xuanwu
```

#### **Upload a File**
Upload a file to the decentralized storage:
```sh
./dist/storage upload --path ./resources/hello.html --filename hello.html
```
You will receive a CID, for example, `baxxxxx`.

#### **Bind File Name and CID**
Bind the file name with the received CID:
```sh
./dist/storage contract bind --cid baxxxxx --name hello.html
```
