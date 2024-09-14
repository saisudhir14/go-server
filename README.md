# 🌐 Go Local Server

Welcome to the **Go Local Server** project! This is a simple HTTP server built using Go that demonstrates handling HTTP requests, serving static files, and processing form data.

## 📁 Project Structure

The project includes the following files:

- **`main.go`**: The main Go source file for the server.
- **`form.html`**: An HTML form for sending POST requests.
- **`index.html`**: A static HTML page served at the root endpoint.

## 🗂️ Files

### `main.go`

This is the core file of the server. It sets up routes and handlers for different endpoints:

- **`/`** - Serves static files from the `./static` directory.
- **`/form`** - Handles POST requests with form data. Displays the submitted `name` and `password`.
- **`/hello`** - Responds with a "hello!" message if accessed via a GET request.

## 🚀 Setup and Running the Server

1. **Install Go**: Ensure Go is installed on your system. If not, download and install it from the official [Go website](https://go.dev/doc/install).
2. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
3. Run the Server:
<code> go run main.go </code>

4. Access the Server:
    * Home Page: Open http://localhost:8080 to view index.html.
    * Form Page: Open http://localhost:8080/form to access and use the form.
    * Hello Page: Open http://localhost:8080/hello to receive a "hello!" response.
    * The ./static directory contains your static files, including index.html and form.html.
    * The server listens on port 8080 by default. Feel free to change the port number in main.go if needed.
