# Go OAuth 2.0 Authentication Example

This is a simple Go web application demonstrating OAuth 2.0 authentication using the [Echo](https://echo.labstack.com/) web framework, the [Goth](https://github.com/markbates/goth) authentication library, and [Viper](https://github.com/spf13/viper) for configuration management.

## Getting Started

### Clone the Repository

To clone the repository, run:

```bash
git clone https://github.com/syahriarreza/go-simple-oauth2.git
cd go-simple-oauth2
```

### Install Dependencies

Install the necessary dependencies with:

```bash
go mod tidy
```

### Configuration

Create a `config.yml` file with the following format:

```yaml
google_client_id: "your-google-client-id"
google_client_secret: "your-google-client-secret"
google_url_callback: "http://localhost:3000/auth/google/callback"
facebook_key: "your-facebook-key"
facebook_secret: "your-facebook-secret"
facebook_callback_url: "http://localhost:3000/auth/facebook/callback"
apple_key: "your-apple-key"
apple_secret: "your-apple-secret"
apple_callback_url: "http://localhost:3000/auth/apple/callback"
```

Replace the placeholder values with your actual credentials.

### Running the Application

To run the application, use:

```bash
go run main.go
```

Now, navigate to `http://localhost:3000` in your browser to see the login page.

## How It Works

1. **Configuration**: Viper reads configuration settings from `config.yml`.
2. **Providers**: Goth is configured with OAuth 2.0 providers.
3. **Routes**: Echo defines routes for the login page, authentication, and callback.
4. **Server**: The server is started on port 3000.

## License

This project is licensed under the MIT License.