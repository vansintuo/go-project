package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func createFile(filename string, content string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.WriteString(content)
    if err != nil {
        return err
    }
    fmt.Printf("File %s created successfully.\n", filename)
    return nil
}

func mergeFiles(outputFile string, inputFiles []string) error {
    outFile, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer outFile.Close()

    for _, inputFile := range inputFiles {
        content, err := ioutil.ReadFile(inputFile)
        if err != nil {
            return err
        }
        _, err = outFile.Write(content)
        if err != nil {
            return err
        }
        _, err = outFile.WriteString("\n")
        if err != nil {
            return err
        }
    }

    fmt.Printf("Files merged into %s successfully.\n", outputFile)
    return nil
}

func main() {
    err1 := createFile("file1.txt", "This is the content of file 1.")
    err2 := createFile("file2.txt", "This is the content of file 2.")
    err3 := createFile("file3.txt", "This is the content of file 3.")

    if err1 != nil || err2 != nil || err3 != nil {
        fmt.Println("Error creating files:", err1, err2, err3)
        return
    }

    inputFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
    err := mergeFiles("merged.txt", inputFiles)
    if err != nil {
        fmt.Println("Error merging files:", err)
        return
    }

    fmt.Println("Program completed successfully.")
}