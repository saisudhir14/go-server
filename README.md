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

5. Containerize the application using docker:

    Commands:

    *  Build docker image
    ```docker build -t go-server-app .```
    * Run docker container
    ``` docker run -p 8080:8080 go-server-app ```
    * test the application!
    ```docker run -p 8080:8080 go-server```
    * Access the Application on in your browser: 
    ``` http://localhost:8080 ```

6. Export
 - docker login
 - docker tag go-server-app saisudhir14/go-server-app:v1  // Use your username instead
 - docker push saisudhir14/go-server-app:v1

But, let me say if you are trying to use either  Amazon Elastic Container Registry (ECR), Google Container Registry (GCR), or Azure Container Registry (ACR), the process is similar, but you need to configure authentication with the specific registry and push the image there...

- aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.us-east-1.amazonaws.com

you can use this command for AWS

But,  If you don't want to use a registry and just want to send the Docker image to another machine (for example, using a file transfer method or you want to transfer the Docker Image Directly (Manual Shipping)), you can:
- docker save -o go-server-app.tar go-server-app
save the image to a tar file

This creates a tarball go-server-app.tar containing the Docker image.

So, to load the Image on the Target Machine
On the target machine, you can load the image back into Docker:

docker load -i go-server-app.tar

to Run the Container on the Target Machine After loading the image, you can run the container as usual:
- docker run -p 8080:8080 go-server-app

Run Docker Images on Local Machines (e.g., Clients or Test Environments)
If you want others to run your containerized application locally:

You can send them the image via a Docker registry (public or private) or by sending them the image tarball.
Once they have access to the image, they can pull it with:
- docker pull saisudhir14/go-server-app:v1
Then, you can run the container on their local machine with:
docker run -p 8080:8080 saisudhir14/go-server-app:v1


Shipping Methods:

- Docker Registry (Docker Hub, ECR, GCR, etc.): Push the image to a registry and pull it from the target machine. 
- Manual Transfer: Save the image as a .tar file, transfer it, and load it on the target system. 
- CI/CD Pipeline: Use tools like Jenkins, GitHub Actions, or GitLab CI to automate the building, testing, and deployment process. 
- Cloud Deployment: Deploy your containerized application on cloud services (AWS ECS, GKE, Kubernetes, etc.). 
- Docker Compose: Use for multi-container applications and ship both the configuration and images.
