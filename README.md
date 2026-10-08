Go must be installed on the device running the CLI

Once installed cd into the repo root directory so that you can run the program.

Run the main.go file directly by specifying its path:
```
go run main.go
```
Testing Info:

The program will ask for the user to input the needed mode (e) for encryption, (d) for decryption. Once selected the program will ask for a string, a quit command (q), or a change mode command (r). Upon entering a string the program will ask for the key, once a valid integer is passed into it the result of the encryption or decryption will be displayed