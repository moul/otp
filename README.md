# otp

[One-Time Pad](https://en.wikipedia.org/wiki/One-time_pad) utility

## Synopsis

    INPUT | otp <pad>

## Description

The `otp` utility applies the [OTP](https://en.wikipedia.org/wiki/One-time_pad) cipher against the text passed as input with the key passed as argument.

As for [ROT13](https://en.wikipedia.org/wiki/ROT13), `otp` is its own inverse; that is, to undo `otp`, the same algorithm is applied with the same key.

    $plaintext == $plaintext | otp $key | otp $key

## Usage

```console
$ cat file.plain | otp "a-random-string-with-same-the-length-of-the-input-file" > /path/to/file.encrypted
```

## Example

```console
$ echo "hello world!" | otp "abcdefghijklm"
<random-looking-binary-output>
$ echo "hello world!"echo  | otp "abcdefghijklm" | otp "abcdefghijklm"
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
