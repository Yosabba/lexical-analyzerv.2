# Lexical Analyzer

This program reads an input file, tokenizes its content, and writes the tokens to an output file. It also prints the tokens to the console.

## Golang needs to be installed to run
[https://go.dev/doc/install](https://go.dev/doc/install)

## How It Works

1. **Reading the Input File:**
   - The program prompts the user to enter the name of the input file.
   - It reads the content of the specified file using the `readUserFile` function.
   - The `readUserFile` function opens the file, reads its content, and returns it as a string.

2. **Tokenizing the Content:**
   - The program tokenizes the content of the input file using the `tokenizeFile` function.
   - The `tokenizeFile` function processes the input string character by character, identifying tokens such as identifiers, integer literals, operators, and delimiters.
   - It returns a slice of `Token` structs, each containing a lexeme and its corresponding token type.

3. **Writing Tokens to the Output File:**
   - The program prompts the user to enter the name of the output file.
   - It writes the tokens to the specified output file using the `writeTokensToFile` function.
   - The `writeTokensToFile` function creates the file, writes each token to the file in a formatted manner, and closes the file.

4. **Printing Tokens to the Console:**
   - The program prints the tokens to the console in a tabular format for easy viewing.

## Usage

1. Create a sample input file (e.g., `input.txt`) with some content:
    ```
    int x = 10;
    x = x + 1;
    ```

2. Run the program using the `go run` command:
    ```sh
    go run main.go
    ```

3. When prompted, enter the name of the input file (e.g., `input.txt`) and the desired name for the output file (e.g., `output.txt`).

4. Check the contents of the output file to verify that the tokens have been written correctly.

## Example

```sh
$ go run main.go
What is the name of your input file?
Please add the extension (i.e: example.txt)
input.txt
What is the name of your output file?
Please add the extension (i.e: output.txt)
output.txt
Lexeme     Token     
--------------------
int        identifier
x          identifier
=          operator  
10         int_literal
;          delimiter 
x          identifier
=          operator  
x          identifier
+          operator  
1          int_literal
;          delimiter