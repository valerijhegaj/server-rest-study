package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"server-rest-study/test/api"
	"server-rest-study/test/format"
)

func Test_server(t *testing.T) {
	os.Chdir("..")
	defer os.Chdir("cmd")
	log.SetOutput(ioutil.Discard)
	go main()

	valerijhegaj := &apiParser.User{
		Username: "valerijhegaj", Password: "123", PORT: 4444,
	}

	//----------------------test1---------------------------------------
	// create user, log in, create file, get, update, delete
	{
		code, err := valerijhegaj.Register()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = valerijhegaj.LogIn(60)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		fileData := "some data"
		path := "valerijhegaj/test/test.txt"
		code, err = valerijhegaj.CreateFile(
			path, fileData,
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		fileData = "new data"
		code, err = valerijhegaj.UpdateFile(
			path, fileData,
		)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
	}

	//----------------------test2---------------------------------------
	// tries to create user with same nick
	{
		hacker := &apiParser.User{
			Username: valerijhegaj.Username, Password: "wrong", PORT: 4444,
		}
		code, err := hacker.Register()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = hacker.LogIn(60)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
	}

	aboba := &apiParser.User{
		Username: "aboba", Password: "abob", PORT: 4444,
	}

	//----------------------test3---------------------------------------
	// create another user, and create, get, update, delete foreign file
	{
		code, err := aboba.Register()
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = aboba.LogIn(60)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err = valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, _, err = aboba.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err := valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
	}

	//----------------------test4---------------------------------------
	// give r rights for files of one user to another
	// another try to create, get, update, delete foreign file
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"

		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "r")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
	}

	//----------------------test5---------------------------------------
	// give w rights for files of one user to another
	// another try to create, get, update, delete foreign file
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "w")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		if string(data) == fileData {
			t.Error(format.ErrorString("nil", string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData+"evil" {
			t.Error(format.ErrorString(fileData+"evil", string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil2")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData+"evil2" {
			t.Error(format.ErrorString(fileData+"evil2", string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
	}

	//----------------------test6---------------------------------------
	// give rw rights for files of one user to another
	// another try to create, get, update, delete foreign file
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "rw")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData+"evil" {
			t.Error(format.ErrorString(fileData+"evil", string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil2")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData+"evil2" {
			t.Error(format.ErrorString(fileData+"evil2", string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
	}

	//----------------------test7---------------------------------------
	// give rw and take away all rights for files of one user to another
	// another try to create, get, update, delete foreign file
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "rw")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			t.Error(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err = aboba.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		if string(data) == fileData {
			t.Error(format.ErrorString("nil", string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil2")
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			t.Error(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			t.Error(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			t.Error(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			t.Error(format.ErrorInt(http.StatusOK, code))
		}
	}
}
