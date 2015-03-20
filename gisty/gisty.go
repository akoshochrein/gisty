
package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os/exec"
    "path"
    "strings"
)

type FileData struct {
    FileName string `json:"filename"`
    FileContent string `json:"content"`
}

type GistInfoContainer struct {
    // Description string `json:"description"`
    Files map[string]FileData `json:"files"`
    IsPublic bool `json:"public"`
}

func main() {
    files := map[string]FileData{}
    fileNames := GetFileNamesFromParams()
    for i := range fileNames {
        data := ReadFileData(fileNames[i])

        truncFileName := path.Base(fileNames[i])
        files[truncFileName] = FileData{
            FileName: truncFileName,
            FileContent: data}
    }

    completeJson := GistInfoContainer{
        Files: files,
        IsPublic: false}

    filesEncoded, err := json.Marshal(completeJson)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(filesEncoded))

    gitUserName, gitPassword := GetGitAuthData()
    url := fmt.Sprintf("https://%s:%s@api.github.com/gists", gitUserName, gitPassword)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(filesEncoded))

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    fmt.Printf(resp.Status)
}

func GetGitAuthData() (string, string) {
    userName := GetGitParam("user.name")
    password := GetGitParam("user.password")
    if len(password) == 0 {
        password = GetGitPasswordForUser(userName)
    }
    return userName, password
}

func GetGitParam(paramName string) (string) {
    paramValue, err := exec.Command("git", "config", "--get-all", paramName).Output()
    paramValueString := ""
    if err == nil {
        paramValueString = strings.Replace(string(paramValue), "\n", "", -1)
    }
    return paramValueString
}

func GetGitPasswordForUser(userName string) (password string) {
    // TODO turn echo off when asking for password
    fmt.Printf("Password for %s: ", userName)
    _, err := fmt.Scanf("%s", &password)
    if err != nil {
        log.Fatal(err)
    }
    return
}

func GetGistName() (gistName string) {
    fmt.Printf("Gist name:")
    _, err := fmt.Scanf("%s", &gistName)
    if err != nil {
        log.Fatal(err)
    }
    return
}

func GetFileNamesFromParams() (fileNameList []string) {
    flag.Parse()
    return flag.Args()
}

func ReadFileData(filePath string) (string) {
    dat, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
    return string(dat)
}
