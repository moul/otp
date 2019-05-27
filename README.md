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
$ echo -n "hello world!" | hexdump -C
00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 21              |hello world!|
```

```console
$ echo -n "@@@@@@@@@@@@" | hexdump -C
00000000  40 40 40 40 40 40 40 40  40 40 40 40              |@@@@@@@@@@@@|
```

```console
$ echo -n "hello world!" | otp "@@@@@@@@@@@@" | hexdump -C
00000000  28 25 2c 2c 2f 60 37 2f  32 2c 24 61              |(%,,/`7/2,$a|
```

```console
$ echo -n "@@@@@@@@@@@@" | otp "hello world!" | hexdump -C
00000000  28 25 2c 2c 2f 60 37 2f  32 2c 24 61              |(%,,/`7/2,$a|
```

```console
$ echo -n "hello world!" | otp "@@@@@@@@@@@@" | otp "@@@@@@@@@@@@" | hexdump -C
00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 21              |hello world!|
```

```console
$ echo -n "hello world!" | otp "hello world!" | hexdump -C
00000000  00 00 00 00 00 00 00 00  00 00 00 00              |............|
```

```console
$ printf "\0\0\0\0\0\0\0\0\0\0\0\0" | otp "hello world!" | hexdump -C
00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 21              |hello world!|
```
