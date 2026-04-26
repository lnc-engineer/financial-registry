concept

os.ReadFile reads file content into memory

How it works

file path passed via CLI
program reads content as bytes
converted to string

TEST
valid file - content printed
missing file - error message

code patern

nano docs/session-3-file-reading.mddata, err := os.ReadFile(file)


System connection
this is the ingestion layer of my registry
accepts external data and process it


