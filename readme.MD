# Golang S3 Wrapper Library
The Golang S3 Wrapper Library is a flexible and extensible wrapper around various S3-compatible storage services. It provides a unified interface for interacting with object storage services like AWS S3, MinIO, and any other S3-compatible storage systems. This library abstracts the underlying implementation details, enabling developers to switch between different storage providers with minimal code changes.

## features
- **Unified Interface**: Interact with multiple S3-compatible storage services through a consistent set of methods.
- **Create Presigned URLs**: Generate presigned URLs for securely sharing or accessing objects stored in S3.
- **Context Support**: Leverage Go's context package for handling timeouts, cancellation, and request-scoped data.
- **Extensible Design**: Easily extend the library to support additional S3-compatible storage providers by implementing the provided interface.

## use case
- **Multi-cloud Storage Solutions**: Seamlessly switch between different S3-compatible providers without rewriting storage logic.
- **Secure File Sharing**: Generate presigned URLs for temporary, secure access to files stored in S3.
- **Data Backup & Retrieval**: Upload and retrieve files from various S3-compatible storage systems.