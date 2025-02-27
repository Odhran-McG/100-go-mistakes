package main

import (
	"fmt"
	"log"
	"net/http"
)

func clientShadowing() error {
	var tracing = true
	//	declare client
	var client *http.Client

	if tracing {
		//	variable shadowed, client assigned to tracing client
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		//	variable shadowed, client assigned to default client
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}
	// use client
	return fmt.Errorf("return")
}

func clientShadowing_improved() error {
	var tracing = true

	var client *http.Client
	// declare error var here, use direct assignment operator =
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	// mutualise error handling
	if err != nil {
		return err
	}
	log.Println(client)
	// use client
	return fmt.Errorf("return")
}

func createClientWithTracing() (*http.Client, error) {
	return http.DefaultClient, nil
}

func createDefaultClient() (*http.Client, error) {
	return http.DefaultClient, nil
}
