
package main

import (
    "bytes"
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
)

import "github.com/howeyc/gopass"

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
        safeFileName := strings.Replace(fileNames[i], fmt.Sprintf("%c", os.PathSeparator), "_", -1)
        files[safeFileName] = FileData{
            FileName: safeFileName,
            FileContent: data}
    }

    completeJson := GistInfoContainer{
        Files: files,
        IsPublic: false}

    filesEncoded, err := json.Marshal(completeJson)

    if err != nil {
        log.Fatal(err)
    }

    gitUserName, gitToken := GetGitAuthData()
    client := &http.Client{}
    req, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(filesEncoded))
    req.SetBasicAuth(gitUserName, gitToken)
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    var dat map[string]interface{}

    content, _ := ioutil.ReadAll(resp.Body)
    if err := json.Unmarshal(content, &dat); err != nil {
        panic(err)
    }
    if resp.StatusCode >= 300 {
        log.Fatal(resp.Status)
    }
    allDone := fmt.Sprintf("All done! Find your uploaded files @ https://gist.github.com/%s/", gitUserName)
    fmt.Println(allDone)
}

func GetGitAuthData() (string, string) {
    userName := GetGitParam("user.name")
    token := GetGitParam("user.token")
    if len(token) == 0 {
        token = GetGitPasswordForUser(userName)
    }
    return userName, token
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
    fmt.Printf("Password for %s: ", userName)
    password = string(gopass.GetPasswdMasked())
    return
}

func GetGistName() (gistName string) {
    fmt.Printf("Gist username:")
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
