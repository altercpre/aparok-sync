# Aparok Sync

`aparok-sync` is an open-source protocol designed for the post-cloud era, enabling high-speed, peer-to-peer synchronization of encrypted communication channels (spaces). It provides a communication layer for building private, decentralized applications offering unparalleled control, privacy, and performance.

## Core Principles

Each `aparok-sync` space (communication channel) is:

* **End-to-end encrypted**: Ensuring complete privacy of all messages and data.
* **User-owned**: Users maintain full control over their data and connections.
* **Permissionless**: Free from centralized lock-in, enabling true decentralization.

Thanks to its local-first design, `aparok-sync` achieves better-than-cloud performance when peers are physically close while maintaining global accessibility. At its core, data in `aparok-sync` is stored as encrypted Directed Acyclic Graphs (DAGs), representing various formats such as chats, pages, or databases. Its end-to-end encrypted structure ensures that no external entity can view a channel's content.

## Use Cases

`aparok-sync` enables developers to build private, decentralized alternatives to apps like Telegram, Discord, Notion, or even health-focused tools like Strava and Oura—without centralized infrastructure.

Paired with [`aparok-store`](https://github.com/altercpre/aparok-store), it provides a local-first foundation for intelligent apps, including those using LLMs. Developers work with a local object store, while sync, communication, and scaling are handled seamlessly by the protocol—offering privacy, offline support, and strong performance with minimal complexity.

## Key Features

* **Encrypted, user-owned channels**: Maintain control over communication and data.
* **Permissionless operation**: No strict reliance on centralized services.
* **Local-first sync**: Functions offline and over peer-to-peer connections.
* **Seamless provider switching**: Retain access to channels even when changing sync providers.
* **Speed and security**: Combines fast performance with robust security.
* **Scalable infrastructure**: Efficiently supports large-scale collaboration, including large groups.

## Protocol overview
Please read the [overview](https://sync.any.org) of protocol entities and design.

## Implementation

You can find the various parts of the protocol implemented in Go in the following repositories:
- [`aparok-sync-node`](https://github.com/altercpre/aparok-sync-node) — implementation of a sync node responsible for storing spaces and objects.
- [`aparok-sync-filenode`](https://github.com/altercpre/aparok-sync-filenode) — implementation of a file node responsible for storing files.
- [`aparok-sync-consensusnode`](https://github.com/altercpre/aparok-sync-consensusnode) — implementation of a consensus node responsible for ACL changes monitoring and validation.
- [`aparok-sync-coordinator`](https://github.com/altercpre/aparok-sync-coordinator) — implementation of a coordinator node responsible for network configuration management.

## Contribution
Thank you for your desire to develop Aparok together!

❤️ This project and everyone involved in it is governed by the [Code of Conduct](https://github.com/altercpre/.github/blob/main/docs/CODE_OF_CONDUCT.md).

🧑‍💻 Check out our [contributing guide](https://github.com/altercpre/.github/blob/main/docs/CONTRIBUTING.md) to learn about asking questions, creating issues, or submitting pull requests.

🫢 For security findings, please email [security@altercpre.w8n.pw](mailto:security@altercpre.w8n.pw) and refer to our [security guide](https://github.com/altercpre/.github/blob/main/docs/SECURITY.md) for more information.

🤝 Follow us on [Github](https://github.com/altercpre) and join the [Contributors Community](https://github.com/orgs/altercpre/discussions).

---
Made by Altercpre — privacy-first, security-focused 🔒

Licensed under [MIT License](./LICENSE.md).
