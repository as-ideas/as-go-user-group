
1.  quick run through my talk at the move box:
you can see the slides and examples here (it has additional info at the End)
http://go-talks.appspot.com/github.com/advincze/go-talks/boxtalk/talk.slide

2. Install the Go tools
https://golang.org/dl/
https://golang.org/doc/install

3. Setup GOPATH
https://golang.org/doc/code.html#Workspaces
https://github.com/golang/go/wiki/GOPATH
basically  as a beginner, 
- put all your  GO projects and their dependencies in one folder e.g. $HOME/dev/go 
- export this as an environment variable (add the following line to your ~/.bashrc or similar)
    export GOPATH=$HOME/dev/go
- verify 
    $ go env 
should include  GOPATH=/Users/yourname/dev/go
- download a third party lib e.g.:
    $ go get github.com/PuerkitoBio/goquery 
downloads and compiles the go query lib.
(you can import it now with import " github.com/PuerkitoBio/goquery " in your sources)

4. setup sublime text 
- install sublime : http://www.sublimetext.com/3
- install package control: https://packagecontrol.io/installation
- install gosublime : CMD+SHIFT+P -> install package -> gosublime
- restart sublime text from console (where GOPATH is set) 
    $ /Applications/Sublime\ Text.app/Contents/SharedSupport/bin/subl
- CMD+SHIFT+P -> "gosublime" shold give you options (and keyboard hints)

5. Write a hello World program

    $ mkdir -p  $GOPATH/src/github.com/myname/gotutorial
    $ cd $GOPATH/src/github.com/myname/gotutorial
    $ touch main.go
- find an example on 
http://play.golang.org/


6. Next steps

-  We agreed to meet every two weeks for two hours
I'll setup a doodle for this
- take the go tour https://tour.golang.orghttps://tour.golang.org until next time , if you experience any problems, have questions , write to this mailing list or take a note for the next session
- Sebastian suggested to try one of the chanllenges here:
http://golang-challenge.com/ 

