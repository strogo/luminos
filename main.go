// Copyright (c) 2012-2014 José Carlos Nieto, https://menteslibres.net/xiam
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"log"
	"os"

	"menteslibres.net/gosexy/cli"
)

// Handy path separator.
const pathSeparator = string(os.PathSeparator)

// Version holds the software version.
const Version = "0.9"

func main() {
	// Software properties.
	cli.Name = "Luminos Markdown Server"
	cli.Homepage = "https://menteslibres.net/luminos"
	cli.Author = "J. Carlos Nieto"
	cli.Version = Version
	cli.AuthorEmail = "jose.carlos@menteslibres.net"

	// Shows banner
	cli.Banner()

	// Dispatches the command.
	if err := cli.Dispatch(); err != nil {
		log.Fatal("Could not start Luminos: ", err)
	}

}
