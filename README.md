# otp

[One-Time Pad](https://en.wikipedia.org/wiki/One-time_pad) utility

## Synopsis

```
SYNOPSIS
    INPUT | otp <pad>

DESCRIPTION
    The otp utility 
```

```console
$ cat file.plain | otp "a-random-string-with-same-the-length-of-the-input-file" > /path/to/file.encrypted
```

## Example

```console
$ echo "hello world!" | otp "abcdefghijklm"
<random-looking-binary-output>
$ echo "hello world!" | otp "abcdefghijklm" | otp "abcdefghijklm"
hello world!
```

```console
$ echo "hello world!" | otp "abcdefghijklm" | hexdump -C
00000000  09 07 0f 08 0a 46 10 07  1b 06 0f 4d 67           |.....F.....Mg|
0000000d
# test
$ echo "hello world!" | otp "abcdefghijklm" | otp "abcdefghijklm" | hexdump -C
00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 21 0a           |hello world!.|
0000000d
```
