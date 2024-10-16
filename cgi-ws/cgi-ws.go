// MIT License
//
// Copyright (c) 2024 Alessandro Salerno
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
  "fmt"
  "net/http"
  "net/http/cgi"
  "os"
  "os/exec"
  "bufio"
  "strings"
  "io"
)

func download(filepath string, url string) {
  fmt.Print("Downloading <")
  fmt.Print(url)
  fmt.Print("> to ")
  fmt.Print(filepath)
  fmt.Print(" ... ")

  out, err := os.Create(filepath)
  if err != nil  {
    fmt.Println("IO Error")
    return
  }
  defer out.Close()

  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("Request Error")
    return
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    fmt.Println("Request Status Error")
    return
  }

  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    fmt.Println("IO Error")
    return
  }
  fmt.Println("Done!")
}

func htmcCGI(w http.ResponseWriter, r *http.Request) {
  handler := cgi.Handler{Path: "./bin/htmc", Dir: "./", Args: []string{"-ll", "off", "-ns"}}
  handler.ServeHTTP(w, r)
}

func exists(path string) (bool) {
  _, err := os.Stat(path)
  return err == nil
}

func main() {
  if !exists("./bin") {
    err := os.MkdirAll("./bin", 0755)
    if err != nil {
      fmt.Println("Unable to create bin/ directory")
      return
    }
  }

  if !exists("./tmp") {
    err := os.MkdirAll("./tmp", 0755)
    if err != nil {
      fmt.Println("Unable to create tmp/ directory")
      return
    }
  }

  if !exists("./include/libhtmc") {
    err := os.MkdirAll("./include/libhtmc", 0755)
    if err != nil {
      fmt.Println("Unable to create include/libhtmc/ directory")
      return
    }
  }

  if (!exists("./bin/htmc") && !exists("./bin/htmc.exe")) || !exists("./bin/libhtmc.a") || !exists("./include/libhtmc/libhtmc.h") {
    fmt.Print("You're missing some important htmc files, proceed with the download? [Y/n]: ")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    
    if (len(text) != 0 && text != "Y") {
      fmt.Println("\t1. Go to <https://github.com/Alessandro-Salerno/htmc> and download the latest release or latest-linux-bundle branch files")
      fmt.Println("\t2. Place the htmc executable (htmc or htmc.exe) in the bin/ directory")
      fmt.Println("\t3. Place the libhtmc.a file in the bin/ directory")
      fmt.Println("\t4. Copy the include/ directory in the working directory")
      fmt.Println("\t5. Make sure that include/libhtmc/libhtmc.h exists and is the libhtmc header file")
      fmt.Println("\t6. Make sure that GCC and GNU ld are installed")
      return
    }

    download("./bin/htmc", "https://alessandro-salerno.github.io/htmc/bin/htmc")
    download("./bin/libhtmc.a", "https://alessandro-salerno.github.io/htmc/bin/libhtmc.a")
    download("./include/libhtmc/libhtmc.h", "https://alessandro-salerno.github.io/htmc/include/libhtmc/libhtmc.h")
    download("./index.htmc", "https://alessandro-salerno.github.io/htmc/examples/index.htmc")
    
    // Temporary hard code
    if exec.Command("chmod", "+x", "./bin/htmc").Run() != nil {
      fmt.Println("Error while setting permissions for ./bin/htmc")
      return
    }
  }

  fmt.Println("Listening on localhost:80")
  fmt.Println()

  http.HandleFunc("/", htmcCGI)
  http.ListenAndServe("localhost:80", nil)
}
