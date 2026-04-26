concept

os.ReadFile reads file content into memory

It returns:
-the file data (as bytes)
-an error if something goes wrong


How it works

file path is passed via CLI arguments

The program takes the file name from os.Args
It reads the file using os.ReadFile
It receives the content as bytes
It converts bytes to string for printing

TEST
valid file - content printed to terminal
program reads file successfully

missing file - error message is displayed "no such file or directory"
program does not crash


code patern

data, err := os.ReadFile(file)
if err != nil {
    fmt.Println("Error reading file:", err)
    continue
}

fmt.Println(string(data))


System connection
this is the ingestion layer of the registry
The system can now:
accepts external file input
read raw data into memory
handle failures safely


This is the first step in building a real data ingestion pipeline

