# Educational Ransomware in Go

This is a simple ransomware written in Go for educational purposes. **IT SHOULD NOT BE USED FOR MALICIOUS PURPOSES!**

## Directories

- **crypt**: Contains functions for encrypting and decrypting files.
- **server**: Contains code for creating a TCP server that can be used by the ransomware to communicate with the attacker.
- **MapDirectories**: Provides functions for mapping directories on the operating system, allowing the ransomware to search and encrypt files more effectively.

## Disclaimer

This ransomware has been created for educational purposes and should only be used for learning and testing purposes in controlled environments. It should not be used to cause harm or for illegal activities. The author assumes no responsibility for the misuse or malicious use of this software.

## Compatibility

This ransomware is designed to work on Windows operating systems. However, it is important to note that its use in production environments or on systems that are not your own may be illegal and is strictly prohibited.

## Building and Running

To build the ransomware, run the following command:<br />
<br />
`make build`<br />
<br />
This command will compile the Go code and produce an executable file.

### Running the Ransomware

To run the ransomware, use the following command: <br />
<br />
`make build`
