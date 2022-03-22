# LetsGo
I will be uploading here all my projects written in Go (or golang) in an attempt to showcase my self-learning of the language and hopefully help other fellow students

## Projects available

### [Massive Cheat Sheet](./Massive_Cheatsheet/massiveCheatSheet.go)
#### What is it
It is a transcription of the 7 hours video from freeCodeCamp.org [Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU)
It contains all the snippets of code used to demonstrate what Go can do for you and some comments from myself to help you read the cheat sheet.

#### How to use
You will find 15 different topics spread over 19 functions in the main function (lines 19 - 39)

You can either: 
- Comment out the specific topic you want to test in the main function and then play around with the snippets in their corresponding lines, or
- Grab the snippets and paste them into your scripts. If you are going to be watching the video, it may come in handy to copy-paste only the text shown on the screen into your IDE so you can try the same things Mike does in the video.

### [Basic concepts](./Basic%20concepts)
#### What is it 
Quick overview of the basics in Go. Currently available:

- [Calgolator](./Basic%20concepts/Calgolator/main.go). A simple calculator to showcase functions.
- [Gobonacci](./Basic%20concepts/Gobonacci/gobonacci.go). Fibonacci in our fave lang.
- [GoFizzBuzz](./Basic%20concepts/GoFizzBuzz/main.go). FizzBuzz written in Go.
- [Maps](./Basic%20concepts/Maps/main.go). The basic concepts of Maps (or "dicts" in Python, or "hashes" in Ruby, or "objects" in JavaScript)
- [nmapCaller](./Basic%20concepts/nmapCaller/main.go). Showcasing how to call other programs from Go.
- [Odd%20or%20even](./Basic%20concepts/Odd_or_Even/odd_or_even.go). Playing around with basic ints.
- [ReadingInput](./Basic%20concepts/ReadingInput/main.go). Quick overview of how to read a stream of input.
- [ReadingWeb](./Basic%20concepts/ReadingWeb/main.go). Quick overview of how to read from the web.
- [Structs](./Basic%20concepts/Structs/main.go). The basic concepts of Structs (very close to what a Class is... But it's not a class. In Go it's like a Map in steroids)

#### How to use
The main purpose is to consult the snippets of code to copy/paste.

### [Go Port Scanner](./Go_Port_Scanner/goPortScanner.go)
#### What is it
It is a simple Port Scanner written in Go to practice pointers, dereferencing and concurrency.
It's in a primitive state so it only scans the localhost and doesn't take any input!

#### How to use
Simply do `go run goPortScanner.go` and it will scan your localhost
