package listfilesarray

import (

	"fmt"
    "os"
	"strings"
	"time"
    "path/filepath"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-listfilesarray")

// MyActivity is a stub for your Activity implementation
type listfilesarray struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &listfilesarray{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *listfilesarray) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *listfilesarray) Eval(ctx activity.Context) (done bool, err error) {
	
	
		loc := ctx.GetInput("Path").(string)
		subs := ctx.GetInput("SubDirectories[Y/N]").(string)
	
		dt := time.Now()
	
	// the function that handles each file or dir
	err = filepath.Walk(loc, func(pathX string, infoX os.FileInfo, errX error) error {

		if errX != nil {
			fmt.Println("error at a path \n", errX, pathX)
			return errX
		}

		if infoX.IsDir() {
			fmt.Println("\n'", pathX, "'", " is a directory.\n")
		} else if subs == "Y" {
				ctx.SetOutput("FileName", infoX.Name())
				ctx.SetOutput("Directory", filepath.Dir(pathX))
				ctx.SetOutput("Extension", filepath.Ext(pathX))
				ctx.SetOutput("Size", infoX.Size())
				ctx.SetOutput("ModTime", infoX.ModTime())
				
				diff := dt.Sub(infoX.ModTime())
				mins := int(diff.Minutes())
					ctx.SetOutput("MinutesDiff", mins)
			} else if filepath.Dir(pathX) == strings.Replace(loc, "/", "\\", -1) {
					ctx.SetOutput("FileName", infoX.Name())
					ctx.SetOutput("Directory", filepath.Dir(pathX))
					ctx.SetOutput("Extension", filepath.Ext(pathX))
					ctx.SetOutput("Size", infoX.Size())
					ctx.SetOutput("ModTime", infoX.ModTime())
					
					diff := dt.Sub(infoX.ModTime())
					mins := int(diff.Minutes())
						ctx.SetOutput("MinutesDiff", mins)
				}
	return nil
   })

	if err != nil {
		fmt.Println("error walking the path : \n", loc, err)
	}

	activityLog.Debugf("Activity has listed out the files Successfully")
	fmt.Println("Activity has listed out the files Successfully")
	
	return true, nil
}

