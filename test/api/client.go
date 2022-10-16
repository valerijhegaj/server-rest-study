package apiParser

import (
	"fmt"
	"log"
	"net/http"

	"server-rest-study/test/format"
)

func main() {
	valerijhegaj := &User{
		Username: "valerijhegaj", Password: "123", PORT: 4444,
	}
	{ // test 1
		code, err := valerijhegaj.Register()
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = valerijhegaj.LogIn(60)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		fileData := "some data"
		path := "valerijhegaj/test/test.txt"
		code, err = valerijhegaj.CreateFile(
			path, fileData,
		)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		fileData = "new data"
		code, err = valerijhegaj.UpdateFile(
			path, fileData,
		)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
	}

	aboba := &User{Username: "aboba", Password: "abob", PORT: 4444}

	{ // test 2
		code, err := aboba.Register()
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, err = aboba.LogIn(60)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err = valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, _, err = aboba.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err := valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
	}

	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "r")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
	}

	fmt.Println("test 4")
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "w")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}
		if string(data) == fileData {
			log.Fatal(format.ErrorString("nil", string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData+"evil" {
			log.Fatal(format.ErrorString(fileData+"evil", string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil2")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData+"evil2" {
			log.Fatal(format.ErrorString(fileData+"evil2", string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusNotFound {
			log.Fatal(format.ErrorInt(http.StatusNotFound, code))
		}
	}

	fmt.Println("test 5")
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "rw")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData+"evil" {
			log.Fatal(format.ErrorString(fileData+"evil", string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil2")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData+"evil2" {
			log.Fatal(format.ErrorString(fileData+"evil2", string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusNotFound {
			log.Fatal(format.ErrorInt(http.StatusNotFound, code))
		}
	}

	fmt.Println("test 6")
	{
		fileData := "data"
		path := "valerijhegaj/test.txt"
		code, err := valerijhegaj.CreateFile(path, fileData)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "r")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err := aboba.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}
		if string(data) != fileData {
			log.Fatal(format.ErrorString(fileData, string(data)))
		}

		code, err = valerijhegaj.GiveAccess(path, aboba.Username, "")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}

		code, data, err = aboba.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusForbidden {
			log.Fatal(format.ErrorInt(http.StatusForbidden, code))
		}
		if string(data) == fileData {
			log.Fatal(format.ErrorString("nil", string(data)))
		}

		code, err = aboba.CreateFile(path, fileData+"evil")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData+"evil" {
			log.Fatal(format.ErrorString(fileData+"evil", string(data)))
		}

		code, err = aboba.UpdateFile(path, fileData+"evil2")
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusCreated {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		code, data, err = valerijhegaj.GetFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusCreated, code))
		}
		if string(data) != fileData+"evil2" {
			log.Fatal(format.ErrorString(fileData+"evil2", string(data)))
		}

		code, err = aboba.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusOK {
			log.Fatal(format.ErrorInt(http.StatusOK, code))
		}

		code, err = valerijhegaj.DeleteFile(path)
		if err != nil {
			log.Fatal(format.ErrorString("without errors", err.Error()))
		}
		if code != http.StatusNotFound {
			log.Fatal(format.ErrorInt(http.StatusNotFound, code))
		}
	}
}
