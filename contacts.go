//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goplacetel.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goplacetel

import (
	"encoding/json"
	"fmt"
)

// ContactsReturn is to decode the json return
type ContactsReturn struct {
	Id             int         `json:"id"`
	UserId         int         `json:"user_id"`
	Speeddail      interface{} `json:"speeddail"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	EmailWork      string      `json:"email_work"`
	Company        string      `json:"company"`
	Address        string      `json:"address"`
	AddressWork    string      `json:"address_work"`
	PhoneWork      string      `json:"phone_work"`
	MobileWork     string      `json:"mobile_work"`
	Phone          string      `json:"phone"`
	Mobile         string      `json:"mobile"`
	Fax            string      `json:"fax"`
	FaxWork        string      `json:"fax_work"`
	FacebookUrl    string      `json:"facebook_url"`
	LinkedinUrl    string      `json:"linkedin_url"`
	XingUrl        string      `json:"xing_url"`
	TwitterAccount string      `json:"twitter_account"`
	Blocked        bool        `json:"blocked"`
	UpdatedAt      string      `json:"updated_at"`
	CreatedAt      string      `json:"created_at"`
}

// Contacts is to get all contacts from the api
func Contacts(token string) (*[]ContactsReturn, error) {

	// To decode the json data
	var contacts []ContactsReturn

	// To call the pages
	page := 1

	// Loop over all sites
	for {

		// Set config for new request
		r := Request{fmt.Sprintf("/contacts?page=%d&per_page=25", page), "GET", token, nil}

		// Send new request
		response, err := r.Send()
		if err != nil {
			return nil, err
		}

		// Decode data
		var decode []ContactsReturn

		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return nil, err
		}

		// Close response body
		response.Body.Close()

		// Add contacts
		for _, value := range decode {
			contacts = append(contacts, value)
		}

		// Check length & break the loop
		if len(decode) < 25 {
			break
		}

	}

	// Return data
	return &contacts, nil

}
