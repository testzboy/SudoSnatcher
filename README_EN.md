# SudoSnatcher

Linux Post-Exploitation Tool: SudoSnatcher

# Disclaimer

This tool is intended for security research and educational purposes only. The author assumes no legal liability or responsibility. Passwords are stored locally only and are not uploaded to the cloud.

# Use Case

After obtaining regular or privileged user login access, this tool can be used to capture `sudo` passwords. Users can set up a watering hole attack to steal legitimate user passwords.

# Usage

Download the specified version of the binary file from [Releases](https://github.com/testzboy/SudoSnatcher/releases). You can run it with the default configuration or specify parameters.

The `-i` parameter specifies the post-exploitation path, and the `-o` parameter specifies the password storage path.

Default paths are as follows:

```
$ ./linux_amd64_SudoSnatcher -h
Usage of ./linux_amd64_SudoSnatcher:
  -i string
    	Path to the script for the alias (default "/tmp/.cache")
  -o string
    	Output file path for saved passwords (default "/tmp/.pass")
```

After running, enter `quit` to automatically clean up traces and restore the device to its default state, retaining only the generated password file.

# Password Types

Passwords are categorized into three states:

```
test:111111:fail
test:000000:success
test:000000:valid
```

- **fail**: Incorrect password
- **success**: Correct password
- **valid**: Password pending verification in a sudo session environment
