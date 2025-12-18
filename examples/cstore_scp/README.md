# DICOM C-STORE SCP Example

This example demonstrates how to receive DICOM files from DICOM clients using the C-STORE operation (Service Class Provider - SCP).

## Features

- Accept incoming DICOM connections
- Handle C-ECHO verification requests
- Receive and store DICOM files
- Support for multiple concurrent connections
- Organize received files by AE title or study date
- Track server statistics
- Graceful shutdown

## Usage

### Basic Usage - Start Server

```bash
go run main.go
```

The server will start on port 11112 and save received files to `./received_dicom/`

### Specify Custom Port and Storage Directory

```bash
go run main.go -port 11112 -storage /data/dicom
```

### Enable Verbose Logging

```bash
go run main.go -verbose
```

### Organize Files by Study Date

```bash
go run main.go -organize-by-date -storage /data/dicom
```

This will organize files as: `/data/dicom/YYYY/MM/DD/sopinstance.dcm`

## Command Line Options

| Option | Default | Description |
|--------|---------|-------------|
| `-port` | 11112 | Port to listen on |
| `-storage` | ./received_dicom | Directory to store received files |
| `-max-conn` | 10 | Maximum concurrent connections (0 = unlimited) |
| `-verbose` | false | Enable verbose logging |
| `-organize-by-ae` | false | Organize files by calling AE title |
| `-organize-by-date` | false | Organize files by study date (YYYY/MM/DD) |

## Examples

### 1. Production Server with Custom Settings

```bash
go run main.go \
  -port 104 \
  -storage /var/dicom/incoming \
  -max-conn 50 \
  -organize-by-date \
  -verbose
```

### 2. Development Server

```bash
go run main.go \
  -port 11112 \
  -storage ./test_storage \
  -verbose
```

### 3. Organized Storage

```bash
# Organize by study date
go run main.go -organize-by-date -storage /data/archive

# Files will be saved as:
# /data/archive/2024/01/15/1.2.840.113619.2.55.3.12345.dcm
```

## Output Example

```
=== DICOM C-STORE SCP Example ===
Port:          11112
Storage Dir:   ./received_dicom
Max Conn:      10
Verbose:       false

Starting DICOM SCP server on port 11112...
Press Ctrl+C to stop

Accepted 4/4 presentation contexts from STORE_SCU

C-STORE received:
  SOP Class:     1.2.840.10008.5.1.4.1.1.2
  SOP Instance:  1.2.840.113619.2.55.3.12345
  Patient:       DOE^JOHN
  Study Date:    20240115
  Modality:      CT
  Saved to:      ./received_dicom/1.2.840.113619.2.55.3.12345.dcm
  Status:        SUCCESS

^C
Received signal: interrupt
Shutting down gracefully...

=== Server Statistics ===
Uptime:           5m23s
Connections:      3
Files Received:   25
Files Failed:     0
Bytes Received:   52428800 (50.00 MB)
Avg File Size:    2048.00 KB

Server stopped.
```

## File Organization

### Default (Flat Structure)

```
./received_dicom/
├── 1.2.840.113619.2.55.3.12345.dcm
├── 1.2.840.113619.2.55.3.12346.dcm
└── 1.2.840.113619.2.55.3.12347.dcm
```

### Organized by Date

```
./received_dicom/
└── 2024/
    └── 01/
        └── 15/
            ├── 1.2.840.113619.2.55.3.12345.dcm
            └── 1.2.840.113619.2.55.3.12346.dcm
```

### Organized by AE Title

```
./received_dicom/
├── MODALITY_1/
│   └── 1.2.840.113619.2.55.3.12345.dcm
└── MODALITY_2/
    └── 1.2.840.113619.2.55.3.12346.dcm
```

## Association Negotiation

The server accepts all presentation contexts by default. In production, you may want to:

1. **Validate AE Titles**: Check if the calling AE is in your whitelist
2. **Restrict SOP Classes**: Only accept specific SOP Classes
3. **Control Transfer Syntaxes**: Accept only transfer syntaxes you support

Modify the `handleAssociationNegotiation` function to implement these checks.

## Statistics

The server tracks:

- Uptime
- Total connections served
- Files received (success)
- Files failed
- Total bytes received
- Average file size

Statistics are displayed when the server shuts down.

## Graceful Shutdown

The server handles `SIGINT` and `SIGTERM` signals gracefully:

1. Stop accepting new connections
2. Wait for active connections to complete (up to 10 seconds)
3. Display statistics
4. Exit

To stop the server, press `Ctrl+C` or send a SIGTERM signal:

```bash
kill <pid>
```

## Production Deployment

### Running on Standard DICOM Port (104)

Port 104 requires root/administrator privileges:

```bash
# Linux/macOS
sudo go run main.go -port 104

# Or build and run as service
go build -o dicom-scp main.go
sudo ./dicom-scp -port 104
```

### Running as Systemd Service

Create `/etc/systemd/system/dicom-scp.service`:

```ini
[Unit]
Description=DICOM C-STORE SCP Server
After=network.target

[Service]
Type=simple
User=dicom
ExecStart=/usr/local/bin/dicom-scp -port 104 -storage /var/dicom/incoming
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl enable dicom-scp
sudo systemctl start dicom-scp
```

## Security Considerations

1. **AE Title Validation**: Implement whitelist to accept only authorized clients
2. **Network Security**: Use firewall rules to restrict access
3. **File Validation**: Validate received files before processing
4. **Disk Space**: Monitor disk usage to prevent DoS
5. **Rate Limiting**: Consider limiting connections per client
6. **Audit Logging**: Log all received files for compliance

## Requirements

- Go 1.21 or later
- Write permissions to storage directory
- Network access on listening port

## Related Examples

- [cstore_scu](../cstore_scu) - DICOM C-STORE sender (client)
- [write_dicom](../write_dicom) - Create DICOM files

## Error Handling

The server handles common errors gracefully:

- **Port already in use**: Check if another server is running
- **Permission denied**: May need root for ports < 1024
- **Disk full**: Server will return error status to client
- **Invalid DICOM**: Returns error and continues serving

## Testing

Test the server with the companion SCU example:

```bash
# Terminal 1: Start SCP
go run main.go -verbose

# Terminal 2: Send file with SCU
cd ../cstore_scu
go run main.go -file test.dcm
```

## Notes

- Default port 11112 avoids requiring root privileges
- Files are named using SOP Instance UID to ensure uniqueness
- The server supports multiple concurrent connections
- All presentation contexts are accepted by default
- Transfer syntax is preserved from received files
