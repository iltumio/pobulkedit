package main

//Packages to import
import (
    "fmt"
    "github.com/alexflint/go-arg"
    "path/filepath"
    "strings"
    "io/ioutil"
)

//Args for the input
type args struct {
  Path     string `arg:"positional"`
  Msgid    string `arg:"required" help:"MessageId"`
  Msgstr   string `arg:"required" help:"Msgstr - new translation"`
  R        bool   `arg:"-r" help:"recoursive"`
}

//Function that show program description when the -h argument is called
func (args) Description() string {
	return "This program will recoursively update .po translation files contained in a folder modifying the translation of a specific MessageId"
}


// Function that read files contained inside the given directory. Every file will be scanned
// to find the given msgid. If it will be found, the function will replace the msgstr with the
// given one
func updateTranslationInDir(dir string, msgid string, translation string){

    //search all *.po files inside the given directory
    files, err := filepath.Glob(dir + "*.po")
    if err != nil {
      fmt.Println("File Error")
    }

    //Loop into every file
    for _, file := range files {
      //Print current processed file name
      fmt.Println("Processing file: " + file)

      //Open the current file
      input, err := ioutil.ReadFile(file)
      //Check if there are errors opening the file
      if err != nil {
        fmt.Println(err)
        return
      }

      //Split the content in a array of lines
      lines := strings.Split(string(input),"\n")

      //Loop on each line
      for i, line := range lines {
        //Check if the msgid is present in the current processed line
        if(strings.Contains(line,msgid)){
          //Print out some useful information
          fmt.Printf("Found on line: %d - ", i)
          //Save the current line value for the logfile
          var tmp string = lines[i+1]
          //Replace the msgstr with the translation that is passed by argument
          lines[i+1] = "msgstr \"" + translation + "\""
          //Print out some other useful information
          fmt.Printf("Replaced ' %s ' with ' %s ' \n",tmp,lines[i+1])
          //Write the log file --coming soon--
        }
      }

      //Write the output in a new file
      output := strings.Join(lines,"\n")
      //Create the file with the extension .new in the same directory of the original one
      err = ioutil.WriteFile(file + ".new", []byte(output),0644)
      //Check for errors
      if err != nil{
        fmt.Println(err)
        return
      }



    }
}

func main() {
  //Args
  var args args
  //Args parsing
  arg.MustParse(&args)

  //If recoursive --r is selected
  if args.R == true {
    if args.Path != "" {
      updateTranslationInDir(args.Path,args.Msgid,args.Msgstr)
    }
  } else {
    fmt.Println(" program is not yet ready to do something more...")
  }
}
