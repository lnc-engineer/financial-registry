concept

CLI programs take input from terminal using os.Args.


os.Args is how a Go program receives input from the terminal.
Index 0 = program name
Index 1+ = user input


how it works


Command line input becomes a slice of strings.


When I run:
go run ./cmd/ingestion-processor data.txt

The program receives:
["program", "data.txt"]

len(os.Args) tells me how many inputs there are.



TEST

Ran without input → program printed "No input file provided"

Ran with one file → printed file name

Ran with multiple files → loop processed both

No args → error message
One arg → processed file
Multiple args → loop processed all files



code pattern

if len(os.Args) < 2 { return }

for i := 1; i < len(os.Args); i++ {
    fmt.Println(os.Args[i])
}



System connection

This is the entry point of my ingestion system.

It allows external data (files) to be passed into the system.



