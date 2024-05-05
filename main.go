package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/apple"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yml") // Set the config file path

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Set up the Goth providers
	goth.UseProviders(
		google.New(viper.GetString("google_client_id"), viper.GetString("google_client_secret"), viper.GetString("google_url_callback")),
		facebook.New(viper.GetString("facebook_key"), viper.GetString("facebook_secret"), viper.GetString("facebook_callback_url")),
		apple.New(viper.GetString("apple_key"), viper.GetString("apple_secret"), viper.GetString("apple_callback_url"), nil),
	)

	// Initialize Echo
	e := echo.New()

	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Login Page</h1><p><a href='/auth/google'>Login with Google</a></p><p><a href='/auth/facebook'>Login with Facebook</a></p><p><a href='/auth/apple'>Login with Apple</a></p>")
	})

	e.GET("/auth/:provider", func(c echo.Context) error {
		// Retrieve the provider from the URL parameter
		provider := c.Param("provider")
		if provider == "" {
			return c.String(http.StatusBadRequest, "Provider not specified")
		}

		// set provider so gothic can get the provider from request (need to do this as we don't use gorilla mux, to prevent error: "you must select a provider")
		q := c.Request().URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request().URL.RawQuery = q.Encode()

		req := c.Request()
		res := c.Response().Writer
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			return c.JSON(http.StatusOK, gothUser)
		}
		gothic.BeginAuthHandler(res, req)
		return nil
	})

	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		req := c.Request()
		res := c.Response().Writer
		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
