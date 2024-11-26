## 1.1.0 - 02-09-2024
### Updated
- migrated to lambda go
- replaced lambda layer with utel traces
- PR: https://github.com/Bancar/uala-adquirencia-orders/pull/85

## 1.0.2- 09-04-2024
### Updated
- add collab var to create orders to collabs if needed
- PR: https://github.com/Bancar/uala-adquirencia-orders/pull/80

## 1.0.1- 21-09-2023
### Updated
- add origin var to create orders with a different origin
- PR: https://github.com/Bancar/uala-adquirencia-orders/pull/78

## 1.0.0- 21-09-2023
### Added
- First release
- Receives a new pos orden by SQS
- Persists pos order in dynamodb
- PR: https://github.com/Bancar/uala-adquirencia-orders/pull/68
