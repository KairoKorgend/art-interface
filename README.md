# Art-Interface

Make sure that you have the following installed on your system:

- Go (Golang)
- A web browser to view the interface

## Setup

Clone the repository to your local machine.

```bash
git clone https://gitea.kood.tech/kairokorgend/art.git
cd art\art-interface
```

Make sure you have the following directory structure:

```bash
art-interface/
│   main.go
│   art_decoder.go
│   README.md
├── templates/
│   └── index.html
└── assets/
    ├── css/
    │   └── style.css
    └── images/
        └── background-img.jpg
```

## Running the Server

Navigate to the root directory of the project and execute the following command in your terminal to compile and run the server:

```bash
go run .
```

This command tells Go to compile and run all the .go files in the current directory. The server will start and listen on port 8080. You should see output in your terminal indicating that the server is running:

```bash
Server starting on port 8080...
```

## Using the Application

Choose the Functionality from the dropdown (either Decode or Encode).
Enter the text you wish to encode or decode in the Text area.
Click the Generate button to submit the form and process your input.
The result will be displayed in the area below the button.
