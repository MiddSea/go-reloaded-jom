package main

func TestMain(m *testing.M) {
    // Run all tests
    exitCode := m.Run()
    // Exit with the status code from the tests
    os.Exit(exitCode)
}

func TestMainFunction(t *testing.T) {
    // Create temporary input and output files
    inputFile, err := os.CreateTemp("", "input*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary input file: %v", err)
    }
    defer os.Remove(inputFile.Name())

    outputFile, err := os.CreateTemp("", "output*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary output file: %v", err)
    }
    defer os.Remove(outputFile.Name())

    // Write sample content to the input file
    sampleContent := "This is a test content."
    err = os.WriteFile(inputFile.Name(), []byte(sampleContent), 0644)
    if err != nil {
        t.Fatalf("Failed to write to input file: %v", err)
    }

    // Set up command-line arguments
    os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

    // Capture stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Run main function
    main()

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout
    captured, _ := io.ReadAll(r)

    // Check if the output message is correct
    expectedOutput := fmt.Sprintf("Conversion successful. Result saved in: %s\n", outputFile.Name())
    if string(captured) != expectedOutput {
        t.Errorf("Expected output: %s, got: %s", expectedOutput, string(captured))
    }

    // Check the content of the output file
    outputContent, err := os.ReadFile(outputFile.Name())
    if err != nil {
        t.Fatalf("Failed to read output file: %v", err)
    }

    if string(outputContent) != sampleContent {
        t.Errorf("Expected output content: %s, got: %s", sampleContent, string(outputContent))
    }
}

func TestMainFunction(t *testing.T) {
    // Create temporary input and output files
    inputFile, err := os.CreateTemp("", "input*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary input file: %v", err)
    }
    defer os.Remove(inputFile.Name())

    outputFile, err := os.CreateTemp("", "output*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary output file: %v", err)
    }
    defer os.Remove(outputFile.Name())

    // Write sample content to the input file
    sampleContent := "This is a test content."
    err = os.WriteFile(inputFile.Name(), []byte(sampleContent), 0644)
    if err != nil {
        t.Fatalf("Failed to write to input file: %v", err)
    }

    // Set up command-line arguments
    os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

    // Capture stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Run main function
    main()

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout
    captured, _ := io.ReadAll(r)

    // Check if the output message is correct
    expectedOutput := fmt.Sprintf("Conversion successful. Result saved in: %s\n", outputFile.Name())
    if string(captured) != expectedOutput {
        t.Errorf("Expected output: %s, got: %s", expectedOutput, string(captured))
    }

    // Check the content of the output file
    outputContent, err := os.ReadFile(outputFile.Name())
    if err != nil {
        t.Fatalf("Failed to read output file: %v", err)
    }

    if string(outputContent) != sampleContent {
        t.Errorf("Expected output content: %s, got: %s", sampleContent, string(outputContent))
    }
}
func TestMainWithEmptyInputFile(t *testing.T) {
    // Create a temporary empty input file
    tempInFile, err := os.CreateTemp("", "empty_input_*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary input file: %v", err)
    }
    defer os.Remove(tempInFile.Name())
    tempInFile.Close()

    // Create a temporary output file
    tempOutFile, err := os.CreateTemp("", "output_*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary output file: %v", err)
    }
    defer os.Remove(tempOutFile.Name())
    tempOutFile.Close()

    // Simulate command-line arguments
    os.Args = []string{"cmd", tempInFile.Name(), tempOutFile.Name()}

    // Capture stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Run main function
    main()

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout
    captured, _ := io.ReadAll(r)

    // Check if the output message is correct
    expectedOutput := fmt.Sprintf("Conversion successful. Result saved in: %s\n", tempOutFile.Name())
    if string(captured) != expectedOutput {
        t.Errorf("Expected output: %s, got: %s", expectedOutput, string(captured))
    }

    // Check if the output file is empty
    content, err := os.ReadFile(tempOutFile.Name())
    if err != nil {
        t.Fatalf("Failed to read output file: %v", err)
    }
    if len(content) != 0 {
        t.Errorf("Expected empty output file, got content: %s", string(content))
    }
}

func TestMainWithLargeFile(t *testing.T) {
	// Create a large temporary input file
	tempDir := t.TempDir()
	largeInputFile := filepath.Join(tempDir, "large_input.txt")
	largeOutputFile := filepath.Join(tempDir, "large_output.txt")

	// Generate a large input file (100MB)
	err := generateLargeFile(largeInputFile, 100*1024*1024)
	if err != nil {
		t.Fatalf("Failed to generate large input file: %v", err)
	}

	// Set up command-line arguments
	os.Args = []string{"cmd", largeInputFile, largeOutputFile}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main function
	main()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout
	out, _ := ioutil.ReadAll(r)

	// Check if the output file was created
	if _, err := os.Stat(largeOutputFile); os.IsNotExist(err) {
		t.Errorf("Output file was not created")
	}

	// Check if the success message was printed
	expectedMsg := fmt.Sprintf("Conversion successful. Result saved in: %s", largeOutputFile)
	if !strings.Contains(string(out), expectedMsg) {
		t.Errorf("Expected success message not found in output")
	}

	// Check if the output file size is reasonable (not empty, but smaller than input)
	infoIn, err := os.Stat(largeInputFile)
	if err != nil {
		t.Fatalf("Failed to get input file info: %v", err)
	}
	infoOut, err := os.Stat(largeOutputFile)
	if err != nil {
		t.Fatalf("Failed to get output file info: %v", err)
	}
	if infoOut.Size() == 0 {
		t.Errorf("Output file is empty")
	}
	if infoOut.Size() >= infoIn.Size() {
		t.Errorf("Output file is larger than or equal to input file")
	}
}

func generateLargeFile(filename string, size int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	data := []byte("This is a sample text for the large file. ")
	for i := 0; i < size; i += len(data) {
		if i+len(data) > size {
			data = data[:size-i]
		}
		_, err := f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestMainWithSpecialCharacters(t *testing.T) {
	// Create temporary input and output files
	inputFile, err := os.CreateTemp("", "input*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	outputFile, err := os.CreateTemp("", "output*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary output file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Write sample text with special characters to input file
	sampleText := "Hello, world! This is a test with special characters: @#$%^&*()_+"
	err = os.WriteFile(inputFile.Name(), []byte(sampleText), 0644)
	if err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}

	// Set up command-line arguments
	os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

	// Call main function
	main()

	// Read the output file
	output, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Check if the output matches the expected result
	expected := "Hello, world! This is a test with special characters: @#$%^&*()_+"
	if string(output) != expected {
		t.Errorf("Output does not match expected. Got: %s, Want: %s", string(output), expected)
	}
}

func TestMainOverwritesExistingOutputFile(t *testing.T) {
	// Create temporary input and output files
	inputFile, err := os.CreateTemp("", "input*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	outputFile, err := os.CreateTemp("", "output*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary output file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Write some initial content to the output file
	initialContent := "This is the initial content"
	err = os.WriteFile(outputFile.Name(), []byte(initialContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial content to output file: %v", err)
	}

	// Write some sample content to the input file
	sampleContent := "This is a sample content"
	err = os.WriteFile(inputFile.Name(), []byte(sampleContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write sample content to input file: %v", err)
	}

	// Set up command-line arguments
	os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

	// Call the main function
	main()

	// Read the content of the output file
	resultContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Check if the output file content has changed
	if string(resultContent) == initialContent {
		t.Errorf("Output file was not overwritten")
	}

	// Check if the output file content matches the processed input
	expectedContent := "This is a sample content"
	if string(resultContent) != expectedContent {
		t.Errorf("Output file content does not match expected. Got %q, want %q", string(resultContent), expectedContent)
	}
}

func TestMainCreatesOutputFile(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a sample input file
	inputFile := filepath.Join(tempDir, "input.txt")
	err = os.WriteFile(inputFile, []byte("Sample text"), 0644)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}

	// Set up the output file path (which doesn't exist yet)
	outputFile := filepath.Join(tempDir, "output.txt")

	// Set up command-line arguments
	os.Args = []string{"cmd", inputFile, outputFile}

	// Call the main function
	main()

	// Check if the output file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Output file was not created")
	}
}

func TestMain(t *testing.T) {
    // Create temporary input and output files
    inputFile, err := os.CreateTemp("", "input*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary input file: %v", err)
    }
    defer os.Remove(inputFile.Name())

    outputFile, err := os.CreateTemp("", "output*.txt")
    if err != nil {
        t.Fatalf("Failed to create temporary output file: %v", err)
    }
    defer os.Remove(outputFile.Name())

    // Write sample content to the input file
    sampleContent := "This is a test content."
    err = os.WriteFile(inputFile.Name(), []byte(sampleContent), 0644)
    if err != nil {
        t.Fatalf("Failed to write to input file: %v", err)
    }

    // Set up command-line arguments
    os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

    // Capture stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Run main function
    main()

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout
    captured, _ := io.ReadAll(r)

    // Check if the output message is correct
    expectedOutput := fmt.Sprintf("Conversion successful. Result saved in: %s\n", outputFile.Name())
    if string(captured) != expectedOutput {
        t.Errorf("Expected output: %s, got: %s", expectedOutput, string(captured))
    }

    // Read the content of the output file
    resultContent, err := os.ReadFile(outputFile.Name())
    if err != nil {
        t.Fatalf("Failed to read output file: %v", err)
    }

    // Check if the output file content matches the input
    if string(resultContent) != sampleContent {
        t.Errorf("Output file content does not match input. Got %q, want %q", string(resultContent), sampleContent)
    }
}

func TestMainInvalidInputFile(t *testing.T) {
    // Save the original os.Args
    originalArgs := os.Args
    defer func() { os.Args = originalArgs }()

    // Set up test arguments with an invalid input file
    os.Args = []string{"cmd", "nonexistent_file.txt", "output.txt"}

    // Capture stdout to check the output
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Run the main function
    main()

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout

    // Read the captured output
    var buf bytes.Buffer
    io.Copy(&buf, r)
    output := buf.String()

    // Check if the expected error message is in the output
    expectedError := "Failed to read input file"
    if !strings.Contains(output, expectedError) {
        t.Errorf("Expected error message '%s' not found in output: %s", expectedError, output)
    }
}

func TestMainWriteResultError(t *testing.T) {
	// Create a temporary input file
	inputFile, err := os.CreateTemp("", "input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	// Write some content to the input file
	_, err = inputFile.WriteString("Sample text")
	if err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}
	inputFile.Close()

	// Create a read-only output file to simulate write error
	outputFile, err := os.CreateTemp("", "output_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary output file: %v", err)
	}
	defer os.Remove(outputFile.Name())
	outputFile.Close()

	// Make the output file read-only
	err = os.Chmod(outputFile.Name(), 0444)
	if err != nil {
		t.Fatalf("Failed to change output file permissions: %v", err)
	}

	// Redirect stdout to capture the output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set up arguments for main
	os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

	// Run main and expect it to panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected main to panic, but it didn't")
		}
	}()


func TestMainWithQuotedArguments(t *testing.T) {
    // Create temporary input and output files with spaces in their names
    tempDir := t.TempDir()
    inputFile := filepath.Join(tempDir, "input file.txt")
    outputFile := filepath.Join(tempDir, "output file.txt")

    // Create and write to the input file
    err := os.WriteFile(inputFile, []byte("Sample text"), 0644)
    if err != nil {
        t.Fatalf("Failed to create input file: %v", err)
    }

    // Set up command-line arguments with quotes
    os.Args = []string{"cmd", inputFile, outputFile}

    // Capture stdout
    oldStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Run main function
    main()

    // Restore stdout
    w.Close()
    os.Stdout = oldStdout
    captured, _ := io.ReadAll(r)

    // Check if the output message is correct
    expectedOutput := fmt.Sprintf("Conversion successful. Result saved in: %s\n", outputFile)
    if string(captured) != expectedOutput {
        t.Errorf("Expected output: %s, got: %s", expectedOutput, string(captured))
    }

    // Check if the output file was created
    if _, err := os.Stat(outputFile); os.IsNotExist(err) {
        t.Errorf("Output file was not created")
    }

    // Check the content of the output file
    content, err := os.ReadFile(outputFile)
    if err != nil {
        t.Fatalf("Failed to read output file: %v", err)
    }
    if string(content) != "Sample text" {
        t.Errorf("Expected output file content: 'Sample text', got: %s", string(content))
    }
}
	main()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout


func TestMainWithUnicodeCharacters(t *testing.T) {
	// Create temporary input and output files
	inputFile, err := os.CreateTemp("", "input*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	outputFile, err := os.CreateTemp("", "output*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary output file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Write Unicode characters to the input file
	inputContent := "こんにちは世界 (Hello World in Japanese)"
	err = os.WriteFile(inputFile.Name(), []byte(inputContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write to input file: %v", err)
	}

	// Set up command-line arguments
	os.Args = []string{"cmd", inputFile.Name(), outputFile.Name()}

	// Call the main function
	main()

	// Read the output file
	outputContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Check if the output matches the input (as no processing should change Unicode characters)
	if string(outputContent) != inputContent {
		t.Errorf("Unicode characters not processed correctly. Expected: %s, Got: %s", inputContent, string(outputContent))
	}
}
	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check if the error message is as expected
	expectedError := "failed to writeResult to file:"
	if !strings.Contains(output, expectedError) {
		t.Errorf("Expected error message containing %q, but got: %s", expectedError, output)
	}
}
